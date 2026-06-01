package connector

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conductorone/baton-jumpcloud/pkg/client"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi1"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
)

const (
	apiUserType      = "user"
	apiUserGroupType = "user_group"
	adminAppID       = "2UjI0eIRo77RGFwi2GKpAa0Til0"
)

type appResourceType struct {
	resourceType *v2.ResourceType
	client       *client.Client
	usersCache   *usersCache
}

func (o *appResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *appResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Resource, *sdkResources.SyncOpResults, error) {
	var rv []*v2.Resource

	// If this is the first call to List, we need to create the JumpCloud Administration app
	if opts.PageToken.Token == "" {
		adminApp, err := sdkResources.NewAppResource("JumpCloud Administration", resourceTypeApp, adminAppID, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create JumpCloud Administration app resource: %w", err)
		}
		rv = append(rv, adminApp)
	}

	skip, b, err := unmarshalSkipToken(&opts.PageToken)
	if err != nil {
		return nil, nil, err
	}

	options := &client.Options{}
	apps, nextPageToken, err := o.client.ListApplications(ctx, options.WithSkip(skip))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list applications: %w", err)
	}

	for i := range apps {
		ur, err := appResource(&apps[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to construct app resource during application list: %w", err)
		}
		rv = append(rv, ur)
	}

	pageToken, err := marshalSkipToken(nextPageToken, b)
	if err != nil {
		return nil, nil, err
	}
	return rv, &sdkResources.SyncOpResults{NextPageToken: pageToken}, nil
}

func appResource(app *jcapi1.Application) (*v2.Resource, error) {
	trait, err := appTrait(app)
	if err != nil {
		return nil, fmt.Errorf("failed to construct app trait during app resource creation: %w", err)
	}

	var annos annotations.Annotations
	annos.Update(trait)

	// JumpCloud applications are SSO application connectors that hold their own
	// SSO configuration and credentials, so they are non-human identities of the
	// app-registration type (NHI Phase-1, K3).
	nhi := &v2.NonHumanIdentityTrait{}
	nhi.SetNhiType(v2.NonHumanIdentityTrait_NHI_TYPE_APP_REGISTRATION)
	nhi.SetNhiDetail("jumpcloud.application")
	annos.Update(nhi)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeApp.Id, app.GetId()),
		DisplayName: app.GetDisplayLabel(),
		Annotations: annos,
		Description: app.GetDescription(),
	}, nil
}

func appTrait(app *jcapi1.Application) (*v2.AppTrait, error) {
	ret := &v2.AppTrait{}

	if app.HasLogo() {
		ret.Logo = &v2.AssetRef{Id: app.Logo.GetUrl()}
	}

	if app.HasLearnMore() {
		ret.HelpUrl = app.GetLearnMore()
	}

	return ret, nil
}

func (o *appResourceType) Entitlements(
	ctx context.Context,
	resource *v2.Resource,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Entitlement, *sdkResources.SyncOpResults, error) {
	var rv []*v2.Entitlement

	rv = append(rv, appEntitlement(resource))

	return rv, nil, nil
}

func appEntitlement(resource *v2.Resource) *v2.Entitlement {
	return &v2.Entitlement{
		Id:          fmtResource(resource.Id, resource.Id.GetResource()),
		Resource:    resource,
		DisplayName: fmt.Sprintf("%s app access", resource.DisplayName),
		Description: fmt.Sprintf("Assigned to %s app", resource.DisplayName),
		GrantableTo: []*v2.ResourceType{resourceTypeUser},
		Purpose:     v2.Entitlement_PURPOSE_VALUE_ASSIGNMENT,
		Slug:        "access",
	}
}

type appAdminPrincipal interface {
	GetId() string
}

func (o *appResourceType) adminGrants(ctx context.Context, resource *v2.Resource, opts sdkResources.SyncOpAttrs) ([]*v2.Grant, *sdkResources.SyncOpResults, error) {
	skip, b, err := unmarshalSkipToken(&opts.PageToken)
	if err != nil {
		return nil, nil, err
	}

	appID := resource.Id.GetResource()

	options := &client.Options{}
	users, nextPageToken, err := o.client.ListAdminUsers(ctx, options.WithSkip(skip))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list admin users: %w", err)
	}

	var rv []*v2.Grant

	adminEmails := make([]string, len(users))
	for i := range users {
		adminEmails[i] = users[i].GetEmail()
	}
	systemUsersMap, err := o.usersCache.GetSystemUsersByEmailList(ctx, opts.Session, adminEmails)
	if err != nil {
		return nil, nil, err
	}

	for i := range users {
		adminUser := &users[i]
		var adminPrincipal appAdminPrincipal = adminUser

		// If the user is a system user, we need to fetch the user by email to get the ID
		if systemUser, ok := systemUsersMap[adminUser.GetEmail()]; ok {
			if systemUser != nil {
				adminPrincipal = systemUser
			}
		}

		ur := &v2.Resource{
			Id: &v2.ResourceId{
				ResourceType: resourceTypeUser.Id,
				Resource:     adminPrincipal.GetId(),
			},
		}

		rv = append(rv, &v2.Grant{
			Id: fmtResourceGrant(resource.Id, ur.Id, appID),
			Entitlement: &v2.Entitlement{
				Id:       fmtResource(resource.Id, appID),
				Resource: resource,
			},
			Principal: ur,
		})
	}

	pageToken, err := marshalSkipToken(nextPageToken, b)
	if err != nil {
		return nil, nil, err
	}

	return rv, &sdkResources.SyncOpResults{NextPageToken: pageToken}, nil
}

func (o *appResourceType) Grants(
	ctx context.Context,
	resource *v2.Resource,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Grant, *sdkResources.SyncOpResults, error) {
	if resource.Id.Resource == adminAppID {
		return o.adminGrants(ctx, resource, opts)
	}

	b := pagination.Bag{}
	if opts.PageToken.Token == "" {
		b.Push(pagination.PageState{
			ResourceTypeID: apiUserType,
		})
		b.Push(pagination.PageState{
			ResourceTypeID: apiUserGroupType,
		})
	} else {
		err := b.Unmarshal(opts.PageToken.Token)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal pagination bag during app grants: %w", err)
		}
	}

	current := b.Current()
	if current == nil {
		return nil, nil, nil
	}

	var skip int32
	if current.Token != "" {
		skip64, err := strconv.ParseInt(current.Token, 10, 32)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse skip value from pagination token during app grants: %w", err)
		}
		skip = int32(skip64)
	}
	var rv []*v2.Grant

	var assignments []jcapi2.GraphConnection
	nextPageToken := ""

	if current.ResourceID == "" {
		// No resource ID set, so we are listing associations for the resource type
		options := &client.Options{}
		associations, npt, err := o.client.ListApplicationAssociations(ctx, resource.Id.Resource, options.WithSkip(skip).WithTargets([]string{current.ResourceTypeID}))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list application associations: %w", err)
		}
		assignments = associations
		nextPageToken = npt
	} else if current.ResourceTypeID == apiUserGroupType {
		// We have a resourceID set, and our resource type is user group, so we are listing user group members
		options := &client.Options{}
		members, npt, err := o.client.ListGroupMembers(ctx, current.ResourceID, options.WithSkip(skip))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list group members: %w", err)
		}
		assignments = members
		nextPageToken = npt
	}

	if assignments == nil {
		return nil, nil, fmt.Errorf("unexpected pagination state while listing application grants: resourceID=%s, resourceTypeID=%s", current.ResourceID, current.ResourceTypeID)
	}

	// pops if nextToken is empty, going to the next phase
	err := b.Next(nextPageToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to advance to next page during app grants pagination: %w", err)
	}

	for i := range assignments {
		member := &assignments[i]
		switch member.To.GetType() {
		case apiUserType:
			rv = append(rv, appGrant(resource, resourceTypeUser.Id, member))
		case apiUserGroupType:
			rv = append(rv, appGrant(resource, resourceTypeGroup.Id, member))
			b.Push(pagination.PageState{
				ResourceTypeID: apiUserGroupType,
				ResourceID:     member.To.GetId(),
			})
		default:
			continue
		}
	}

	pt, err := b.Marshal()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal page token after listing app grants: %w", err)
	}

	return rv, &sdkResources.SyncOpResults{NextPageToken: pt}, nil
}

func appGrant(resource *v2.Resource, resoureTypeId string, member *jcapi2.GraphConnection) *v2.Grant {
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

func newAppBuilder(c *client.Client) *appResourceType {
	return &appResourceType{
		resourceType: resourceTypeApp,
		client:       c,
		usersCache:   newUsersCache(c),
	}
}
