package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	sdkEntitlements "github.com/conductorone/baton-sdk/pkg/types/entitlement"
	sdkGrants "github.com/conductorone/baton-sdk/pkg/types/grant"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
)

const adminAppID = "jumpcloud-administration-app"

type administrationAppBuilder struct {
	resourceType *v2.ResourceType
	client1      jc1Func
	ext          *ExtensionClient
}

func (o *administrationAppBuilder) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func newAdministrationAppBuilder(jc1 jc1Func, ext *ExtensionClient) *administrationAppBuilder {
	return &administrationAppBuilder{
		resourceType: resourceTypeAdministrationApp,
		client1:      jc1,
		ext:          ext,
	}
}

func (o *administrationAppBuilder) Entitlements(_ context.Context, resource *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	en := sdkEntitlements.NewAssignmentEntitlement(resource, "administrator", sdkEntitlements.WithGrantableTo(resourceTypeUser))
	rv = append(rv, en)
	return rv, "", nil, nil
}

func (o *administrationAppBuilder) lookForMatchingUser(ctx context.Context, email string) (*jcapi1.Systemuserreturn, error) {
	ctx, client := o.client1(ctx)
	list, resp, err := client.SystemusersApi.SystemusersList(ctx).Search(fmt.Sprintf("email:$eq:%s", email)).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if len(list.Results) == 0 {
		return nil, nil
	}

	if len(list.Results) > 1 {
		return nil, fmt.Errorf("found multiple users with email %s", email)
	}

	return &list.Results[0], nil
}

func (o *administrationAppBuilder) Grants(ctx context.Context, resource *v2.Resource, pt *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	skip, b, err := unmarshalSkipToken(pt)
	if err != nil {
		return nil, "", nil, err
	}

	users, resp, err := o.ext.UserList().Skip(skip).Execute(ctx)
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	ctx, client := o.client1(ctx)

	var rv []*v2.Grant
	for _, u := range users {
		userEmail := u.GetEmail()
		list, resp, err := client.SystemusersApi.SystemusersList(ctx).Filter(fmt.Sprintf("email:$eq:%s", userEmail)).Execute()
		if err != nil {
			return nil, "", nil, err
		}
		defer resp.Body.Close()

		if len(list.Results) != 1 {
			return nil, "", nil, fmt.Errorf("found multiple users with email %s", userEmail)
		}

		user := list.Results[0]

		rv = append(rv, sdkGrants.NewGrant(
			resource,
			"administrator",
			&v2.ResourceId{
				ResourceType: resourceTypeUser.Id,
				Resource:     user.GetId(),
			},
		))
	}

	pageToken, err := marshalSkipToken(len(users), skip, b)
	if err != nil {
		return nil, "", nil, err
	}

	return rv, pageToken, nil, nil
}

func (o *administrationAppBuilder) List(ctx context.Context, parentResourceID *v2.ResourceId, pt *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	var rv []*v2.Resource

	adminApp, err := sdkResources.NewAppResource("JumpCloud Administration", resourceTypeAdministrationApp, adminAppID, nil)
	if err != nil {
		return nil, "", nil, err
	}
	rv = append(rv, adminApp)

	return rv, "", nil, nil
}
