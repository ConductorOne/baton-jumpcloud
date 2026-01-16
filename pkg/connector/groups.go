package connector

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/conductorone/baton-jumpcloud/pkg/client"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/structpb"
)

type groupResourceType struct {
	resourceType *v2.ResourceType
	client       *client.Client
}

func (o *groupResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *groupResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Resource, *sdkResources.SyncOpResults, error) {
	skip, b, err := unmarshalSkipToken(&opts.PageToken)
	if err != nil {
		return nil, nil, err
	}

	options := &client.Options{}
	groups, nextPageToken, err := o.client.ListGroups(ctx, options.WithSkip(skip))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list groups: %w", err)
	}

	var rv []*v2.Resource
	for i := range groups {
		ur, err := groupResource(&groups[i])
		if err != nil {
			return nil, nil, err
		}
		rv = append(rv, ur)
	}

	pageToken, err := marshalSkipToken(nextPageToken, b)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal skip token during group list pagination: %w", err)
	}
	return rv, &sdkResources.SyncOpResults{NextPageToken: pageToken}, nil
}

func groupResource(group *jcapi2.UserGroup) (*v2.Resource, error) {
	trait, err := groupTrait(group)
	if err != nil {
		return nil, fmt.Errorf("failed to construct group trait during group resource creation: %w", err)
	}

	var annos annotations.Annotations
	annos.Update(trait)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeGroup.Id, group.GetId()),
		DisplayName: group.GetName(),
		Annotations: annos,
		Description: group.GetDescription(),
	}, nil
}

func groupTrait(group *jcapi2.UserGroup) (*v2.GroupTrait, error) {
	profile, err := structpb.NewStruct(map[string]interface{}{
		"type":  group.GetType(),
		"email": group.GetEmail(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to construct profile struct during group trait creation: %w", err)
	}

	ret := &v2.GroupTrait{
		Profile: profile,
	}

	return ret, nil
}

func (o *groupResourceType) Entitlements(
	ctx context.Context,
	resource *v2.Resource,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Entitlement, *sdkResources.SyncOpResults, error) {
	var rv []*v2.Entitlement

	rv = append(rv, groupEntitlement(resource))

	return rv, nil, nil
}

func groupEntitlement(resource *v2.Resource) *v2.Entitlement {
	return &v2.Entitlement{
		Id:          fmtResource(resource.Id, resource.Id.GetResource()),
		Resource:    resource,
		DisplayName: fmt.Sprintf("%s Group Member", resource.DisplayName),
		Description: fmt.Sprintf("Member of %s group", resource.DisplayName),
		GrantableTo: []*v2.ResourceType{resourceTypeUser},
		Purpose:     v2.Entitlement_PURPOSE_VALUE_ASSIGNMENT,
		Slug:        "member",
	}
}

func (o *groupResourceType) Grants(
	ctx context.Context,
	resource *v2.Resource,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Grant, *sdkResources.SyncOpResults, error) {
	skip, b, err := unmarshalSkipToken(&opts.PageToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal skip token during group grants pagination: %w", err)
	}

	options := &client.Options{}
	members, nextPageToken, err := o.client.ListGroupMembers(ctx, resource.Id.Resource, options.WithSkip(skip))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list group members: %w", err)
	}

	var rv []*v2.Grant
	for i := range members {
		member := &members[i]
		switch member.To.GetType() {
		case "user":
			rv = append(rv, groupGrant(resource, resourceTypeUser.Id, member))
		case "group":
			rv = append(rv, groupGrant(resource, resourceTypeGroup.Id, member))
		default:
			continue
		}
	}
	pt, err := marshalSkipToken(nextPageToken, b)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal skip token after listing group grants: %w", err)
	}
	return rv, &sdkResources.SyncOpResults{NextPageToken: pt}, nil
}

func (o *groupResourceType) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	if principal.Id.ResourceType != resourceTypeUser.Id {
		l.Warn(
			"baton-jumpcloud: only users can be granted group membership",
			zap.String("principal_type", principal.Id.ResourceType),
			zap.String("principal_id", principal.Id.Resource),
		)
		return nil, fmt.Errorf("only users can be granted group membership, got principal type %s during grant operation", principal.Id.ResourceType)
	}

	err := o.client.AddGroupMember(ctx, entitlement.Resource.Id.Resource, principal.Id.Resource)
	if err != nil {
		// If the user is already a member of the group, we get a 409 and we want to return success
		// with the GrantAlreadyExists annotation.
		if strings.Contains(err.Error(), codes.AlreadyExists.String()) {
			var annos annotations.Annotations
			annos.Update(&v2.GrantAlreadyExists{})
			return annos, nil
		}
		l.Error("jumpcloud-connector: failed to add user to group", zap.Error(err), zap.String("group", entitlement.Resource.Id.Resource), zap.String("user", principal.Id.Resource))
		return nil, fmt.Errorf("failed to add user to group: %w", err)
	}

	return nil, nil
}

func (o *groupResourceType) Revoke(ctx context.Context, grant *v2.Grant) (annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	entitlement := grant.Entitlement
	principal := grant.Principal

	if principal.Id.ResourceType != resourceTypeUser.Id {
		l.Warn(
			"baton-jumpcloud: only users can have group membership revoked",
			zap.String("principal_type", principal.Id.ResourceType),
			zap.String("principal_id", principal.Id.Resource),
		)
		return nil, errors.New("only users can have group membership revoked")
	}

	err := o.client.RemoveGroupMember(ctx, entitlement.Resource.Id.Resource, principal.Id.Resource)
	if err != nil {
		// If the user is already not a member of the group, we get a 404 and we want to return success
		// with the GrantAlreadyRevoked annotation.
		if strings.Contains(err.Error(), codes.NotFound.String()) {
			var annos annotations.Annotations
			annos.Update(&v2.GrantAlreadyRevoked{})
			return annos, nil
		}
		l.Error("jumpcloud-connector: failed to remove user from group", zap.Error(err), zap.String("group", entitlement.Resource.Id.Resource), zap.String("user", principal.Id.Resource))
		return nil, fmt.Errorf("failed to remove user from group: %w", err)
	}

	return nil, nil
}

func groupGrant(resource *v2.Resource, resoureTypeId string, member *jcapi2.GraphConnection) *v2.Grant {
	groupID := resource.Id.GetResource()
	ur := &v2.Resource{Id: &v2.ResourceId{ResourceType: resoureTypeId, Resource: member.To.Id}}

	var annos annotations.Annotations

	return &v2.Grant{
		Id: fmtResourceGrant(resource.Id, ur.Id, groupID),
		Entitlement: &v2.Entitlement{
			Id:       fmtResource(resource.Id, groupID),
			Resource: resource,
		},
		Annotations: annos,
		Principal:   ur,
	}
}

func newGroupBuilder(c *client.Client) *groupResourceType {
	return &groupResourceType{
		resourceType: resourceTypeGroup,
		client:       c,
	}
}
