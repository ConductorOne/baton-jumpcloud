package connector

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

type groupResourceType struct {
	resourceType *v2.ResourceType
	client1      jc1Func
	client2      jc2Func
}

func (o *groupResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *groupResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	token *pagination.Token,
) ([]*v2.Resource, string, annotations.Annotations, error) {
	ctx, client := o.client2(ctx)

	skip, b, err := unmarshalSkipToken(token)
	if err != nil {
		return nil, "", nil, err
	}

	groups, resp, err := client.UserGroupsApi.GroupsUserList(ctx).Skip(skip).Execute()
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	var rv []*v2.Resource
	for i := range groups {
		ur, err := groupResource(ctx, &groups[i])
		if err != nil {
			return nil, "", nil, err
		}
		rv = append(rv, ur)
	}

	pageToken, err := marshalSkipToken(len(groups), skip, b)
	if err != nil {
		return nil, "", nil, err
	}
	return rv, pageToken, nil, nil
}

func groupResource(ctx context.Context, group *jcapi2.UserGroup) (*v2.Resource, error) {
	trait, err := groupTrait(ctx, group)
	if err != nil {
		return nil, err
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

func groupTrait(ctx context.Context, group *jcapi2.UserGroup) (*v2.GroupTrait, error) {
	profile, err := structpb.NewStruct(map[string]interface{}{
		"type":  group.GetType(),
		"email": group.GetEmail(),
	})
	if err != nil {
		return nil, fmt.Errorf("baton-jumpcloud: failed to construct role profile for role trait: %w", err)
	}

	ret := &v2.GroupTrait{
		Profile: profile,
	}

	return ret, nil
}

func (o *groupResourceType) Entitlements(
	ctx context.Context,
	resource *v2.Resource,
	token *pagination.Token,
) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	rv = append(rv, groupEntitlement(ctx, resource))

	return rv, "", nil, nil
}

func groupEntitlement(ctx context.Context, resource *v2.Resource) *v2.Entitlement {
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
	token *pagination.Token,
) ([]*v2.Grant, string, annotations.Annotations, error) {
	ctx, client := o.client2(ctx)

	skip, b, err := unmarshalSkipToken(token)
	if err != nil {
		return nil, "", nil, err
	}

	members, resp, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersList(ctx, resource.Id.Resource).Skip(skip).Execute()
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

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
	pt, err := marshalSkipToken(len(members), skip, b)
	if err != nil {
		return nil, "", nil, err
	}
	return rv, pt, nil, nil
}

func (o *groupResourceType) Grant(ctx context.Context, principal *v2.Resource, entitlement *v2.Entitlement) (annotations.Annotations, error) {
	ctx, client := o.client2(ctx)
	l := ctxzap.Extract(ctx)

	if principal.Id.ResourceType != resourceTypeUser.Id {
		l.Warn(
			"baton-jumpcloud: only users can be granted group membership",
			zap.String("principal_type", principal.Id.ResourceType),
			zap.String("principal_id", principal.Id.Resource),
		)
		return nil, errors.New("baton-jumpcloud: only users can be granted group membership")
	}

	resp, err := client.UserGroupsApi.GraphUserGroupMembersPost(ctx, entitlement.Resource.Id.Resource).Body(jcapi2.GraphOperationUserGroupMember{
		Id:   principal.Id.Resource,
		Op:   "add",
		Type: "user",
	}).Execute()
	if err != nil {
		// If the user is already a member of the group, we get a 409 and we want to return success.
		if resp == nil || resp.StatusCode != http.StatusConflict {
			l.Error("failed to add user to group", zap.Error(err), zap.String("group", entitlement.Resource.Id.Resource), zap.String("user", principal.Id.Resource))
			return nil, err
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusConflict {
		l.Error(
			"failed to add user to group",
			zap.Error(err),
			zap.String("group", entitlement.Resource.Id.Resource),
			zap.String("user", principal.Id.Resource),
			zap.Int("status", resp.StatusCode),
		)
		return nil, errors.New("baton-jumpcloud: failed adding user to group")
	}

	return nil, nil
}

func (o *groupResourceType) Revoke(ctx context.Context, grant *v2.Grant) (annotations.Annotations, error) {
	ctx, client := o.client2(ctx)
	l := ctxzap.Extract(ctx)

	entitlement := grant.Entitlement
	principal := grant.Principal

	if principal.Id.ResourceType != resourceTypeUser.Id {
		l.Warn(
			"baton-jumpcloud: only users can have group membership revoked",
			zap.String("principal_type", principal.Id.ResourceType),
			zap.String("principal_id", principal.Id.Resource),
		)
		return nil, errors.New("baton-jumpcloud: only users can have group membership revoked")
	}

	resp, err := client.UserGroupsApi.GraphUserGroupMembersPost(ctx, entitlement.Resource.Id.Resource).Body(jcapi2.GraphOperationUserGroupMember{
		Id:   principal.Id.Resource,
		Op:   "remove",
		Type: "user",
	}).Execute()
	if err != nil {
		// If the user is already not a member of the group, we get a 404 and we want to return success.
		if resp == nil || resp.StatusCode != http.StatusNotFound {
			l.Error("failed to remove user from group", zap.Error(err), zap.String("group", entitlement.Resource.Id.Resource), zap.String("user", principal.Id.Resource))
			return nil, err
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusNotFound {
		l.Error(
			"failed to remove user from group",
			zap.Error(err),
			zap.String("group", entitlement.Resource.Id.Resource),
			zap.String("user", principal.Id.Resource),
			zap.Int("status", resp.StatusCode),
		)
		return nil, errors.New("baton-jumpcloud: failed removing user from group")
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

func newGroupBuilder(jc1 jc1Func, jc2 jc2Func) *groupResourceType {
	return &groupResourceType{
		resourceType: resourceTypeGroup,
		client1:      jc1,
		client2:      jc2,
	}
}
