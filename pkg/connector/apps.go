package connector

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi2"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

const (
	apiUserType      = "user"
	apiUserGroupType = "user_group"
)

type appResourceType struct {
	resourceType *v2.ResourceType
	client1      jc1Func
	client2      jc2Func
}

func (o *appResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *appResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	token *pagination.Token,
) ([]*v2.Resource, string, annotations.Annotations, error) {
	ctx, client := o.client1(ctx)

	skip, b, err := unmarshalSkipToken(token)
	if err != nil {
		return nil, "", nil, err
	}

	apps, resp, err := client.ApplicationsApi.ApplicationsList(ctx).Skip(skip).Execute()
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	var rv []*v2.Resource
	for i := range apps.Results {
		ur, err := appResource(ctx, &apps.Results[i])
		if err != nil {
			return nil, "", nil, err
		}
		rv = append(rv, ur)
	}

	pageToken, err := marshalSkipToken(len(apps.Results), skip, b)
	if err != nil {
		return nil, "", nil, err
	}
	return rv, pageToken, nil, nil
}

func appResource(ctx context.Context, app *jcapi1.Application) (*v2.Resource, error) {
	trait, err := appTrait(ctx, app)
	if err != nil {
		return nil, err
	}

	var annos annotations.Annotations
	annos.Update(trait)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeApp.Id, app.GetId()),
		DisplayName: app.GetDisplayLabel(),
		Annotations: annos,
		Description: app.GetDescription(),
	}, nil
}

func appTrait(ctx context.Context, app *jcapi1.Application) (*v2.AppTrait, error) {
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
	token *pagination.Token,
) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	rv = append(rv, appEntitlement(ctx, resource))

	return rv, "", nil, nil
}

func appEntitlement(ctx context.Context, resource *v2.Resource) *v2.Entitlement {
	return &v2.Entitlement{
		Id:          fmtResource(resource.Id, resource.Id.GetResource()),
		Resource:    resource,
		DisplayName: fmt.Sprintf("%s app access", resource.DisplayName),
		Description: fmt.Sprintf("Assigned to %s app", resource.DisplayName),
		GrantableTo: []*v2.ResourceType{resourceTypeUser},
		Purpose:     v2.Entitlement_PURPOSE_VALUE_ASSIGNMENT,
		Slug:        resource.DisplayName,
	}
}

type graphRequest interface {
	Execute() ([]jcapi2.GraphConnection, *http.Response, error)
}

func (o *appResourceType) Grants(
	ctx context.Context,
	resource *v2.Resource,
	token *pagination.Token,
) ([]*v2.Grant, string, annotations.Annotations, error) {
	ctx, client := o.client2(ctx)

	b := pagination.Bag{}
	if token.Token == "" {
		b.Push(pagination.PageState{
			ResourceTypeID: apiUserType,
		})
		b.Push(pagination.PageState{
			ResourceTypeID: apiUserGroupType,
		})
	} else {
		err := b.Unmarshal(token.Token)
		if err != nil {
			return nil, "", nil, err
		}
	}

	current := b.Current()
	if current == nil {
		return nil, "", nil, nil
	}

	var skip int32
	if current.Token != "" {
		skip64, err := strconv.ParseInt(current.Token, 10, 32)
		if err != nil {
			return nil, "", nil, err
		}
		skip = int32(skip64)
	}
	var npt string
	var rv []*v2.Grant

	var req graphRequest

	if current.ResourceID == "" {
		// No resource ID set, so we are listing associations for the resource type
		req = client.ApplicationsApi.GraphApplicationAssociationsList(ctx, resource.Id.Resource).Skip(skip).Targets([]string{current.ResourceTypeID})
	} else if current.ResourceTypeID == apiUserGroupType {
		// We have a resourceID set, and our resource type is user group, so we are listing user group members
		req = client.UserGroupMembersMembershipApi.GraphUserGroupMembersList(ctx, current.ResourceID).Skip(skip)
	}

	if req == nil {
		return nil, "", nil, errors.New("unexpected state while listing application grants")
	}

	assignments, resp, err := req.Execute()
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	if len(assignments) != 0 {
		nextSkip := int64(len(assignments)) + int64(skip)
		npt = strconv.FormatInt(nextSkip, 10)
	}
	// pops if nextToken is empty, going to the next phase
	err = b.Next(npt)
	if err != nil {
		return nil, "", nil, err
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
		return nil, "", nil, err
	}

	return rv, pt, nil, nil
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

func newAppBuilder(jc1 jc1Func, jc2 jc2Func) *appResourceType {
	return &appResourceType{
		resourceType: resourceTypeApp,
		client1:      jc1,
		client2:      jc2,
	}
}
