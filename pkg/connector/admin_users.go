package connector

import (
	"context"
	"fmt"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	"google.golang.org/protobuf/types/known/structpb"
)

type adminUserResourceType struct {
	resourceType *v2.ResourceType
	ext          *ExtensionClient
}

func (o *adminUserResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func newAdminUserBuilder(ext *ExtensionClient) *adminUserResourceType {
	return &adminUserResourceType{
		resourceType: resourceTypeAdminUser,
		ext:          ext,
	}
}

func (o *adminUserResourceType) Entitlements(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *adminUserResourceType) Grants(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *adminUserResourceType) List(ctx context.Context, parentResourceID *v2.ResourceId, pt *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	skip, b, err := unmarshalSkipToken(pt)
	if err != nil {
		return nil, "", nil, err
	}

	users, resp, err := o.ext.UserList().Skip(skip).Execute(ctx)
	if err != nil {
		return nil, "", nil, err
	}
	defer resp.Body.Close()

	var rv []*v2.Resource
	for i := range users {
		ur, err := adminUserResource(ctx, &users[i])
		if err != nil {
			return nil, "", nil, err
		}
		rv = append(rv, ur)
	}

	pageToken, err := marshalSkipToken(len(users), skip, b)
	if err != nil {
		return nil, "", nil, err
	}

	return rv, pageToken, nil, nil
}

func adminUserResource(ctx context.Context, user *jcapi1.Userreturn) (*v2.Resource, error) {
	trait, err := adminUserTrait(ctx, user)
	if err != nil {
		return nil, err
	}
	var annos annotations.Annotations
	annos.Update(trait)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeAdminUser.Id, user.GetId()),
		DisplayName: adminUserDisplayName(user),
		Annotations: annos,
	}, nil
}

func adminUserDisplayName(user *jcapi1.Userreturn) string {
	dn := fmt.Sprintf("%s %s (Administrator)", user.GetFirstname(), user.GetLastname())
	if dn != " " {
		return dn
	}
	return user.GetEmail() + " (Administrator)"
}

func adminUserTrait(ctx context.Context, user *jcapi1.Userreturn) (*v2.UserTrait, error) {
	profile, err := structpb.NewStruct(map[string]interface{}{
		"id": user.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("baton-jumpcloud: failed to construct user profile for user trait: %w", err)
	}

	if user.HasOrganization() {
		profile.Fields["organization"] = structpb.NewStringValue(user.GetOrganization())
	}

	ret := &v2.UserTrait{
		Profile: profile,
		Status:  &v2.UserTrait_Status{},
	}
	ret.Status.Status = v2.UserTrait_Status_STATUS_ENABLED

	if user.GetSuspended() {
		ret.Status.Status = v2.UserTrait_Status_STATUS_DISABLED
		ret.Status.Details = "suspended"
	}

	email := user.GetEmail()
	if email != "" {
		ret.Emails = append(ret.Emails, &v2.UserTrait_Email{
			Address:   email,
			IsPrimary: true,
		})
	}

	return ret, nil
}
