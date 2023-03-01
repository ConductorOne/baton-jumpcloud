package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	"google.golang.org/protobuf/types/known/structpb"
)

type userResourceType struct {
	resourceType *v2.ResourceType
	client1      jc1Func
	client2      jc2Func
}

func (o *userResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func newUserBuilder(jc1 jc1Func, jc2 jc2Func) *userResourceType {
	return &userResourceType{
		resourceType: resourceTypeUser,
		client1:      jc1,
		client2:      jc2,
	}
}
func (o *userResourceType) Entitlements(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *userResourceType) Grants(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *userResourceType) List(ctx context.Context, parentResourceID *v2.ResourceId, pt *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	ctx, client := o.client1(ctx)

	// todo manage skip
	list, _, err := client.SystemusersApi.SystemusersList(ctx).Execute()
	if err != nil {
		return nil, "", nil, err
	}

	var rv []*v2.Resource
	for _, user := range list.Results {
		ur, err := userResource(ctx, &user)
		if err != nil {
			return nil, "", nil, err
		}
		rv = append(rv, ur)
	}

	return rv, "", nil, nil
}

func userResource(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.Resource, error) {
	trait, err := userTrait(ctx, user)
	if err != nil {
		return nil, err
	}
	var annos annotations.Annotations
	annos.Update(trait)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeUser.Id, user.GetId()),
		DisplayName: userDisplayName(user),
		Annotations: annos,
	}, nil
}

func userDisplayName(user *jcapi1.Systemuserreturn) string {
	if dn := user.GetDisplayname(); dn != "" {
		return dn
	}
	return fmt.Sprintf("%s %s", user.GetFirstname(), user.GetLastname())
}

func userTrait(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.UserTrait, error) {
	profile, err := structpb.NewStruct(map[string]interface{}{
		"id": user.GetId(),
	})

	if err != nil {
		return nil, fmt.Errorf("baton-jumpcloud: failed to construct user profile for user trait: %w", err)
	}

	if user.HasJobTitle() {
		profile.Fields["job_title"] = structpb.NewStringValue(user.GetJobTitle())
	}

	if user.HasOrganization() {
		profile.Fields["organization"] = structpb.NewStringValue(user.GetOrganization())
	}

	if user.HasDepartment() {
		profile.Fields["department"] = structpb.NewStringValue(user.GetDepartment())
	}

	if user.HasUsername() {
		profile.Fields["username"] = structpb.NewStringValue(user.GetUsername())
	}

	if user.HasManager() {
		profile.Fields["manager_id"] = structpb.NewStringValue(user.GetManager())
	}

	if user.HasEmployeeIdentifier() {
		profile.Fields["employee_id"] = structpb.NewStringValue(user.GetEmployeeIdentifier())
	}

	if user.HasEmployeeType() {
		profile.Fields["employee_type"] = structpb.NewStringValue(user.GetEmployeeType())
	}

	ret := &v2.UserTrait{
		Profile: profile,
		Status:  &v2.UserTrait_Status{},
	}

	switch st := user.GetState(); st {
	case "", "ACTIVATED":
		ret.Status.Status = v2.UserTrait_Status_STATUS_ENABLED
	case "STAGED":
		ret.Status.Status = v2.UserTrait_Status_STATUS_DISABLED
		ret.Status.Details = strings.ToLower(st)
	case "SUSPENDED":
		ret.Status.Status = v2.UserTrait_Status_STATUS_DISABLED
		ret.Status.Details = strings.ToLower(st)
	}

	if user.GetAccountLocked() {
		ret.Status.Status = v2.UserTrait_Status_STATUS_DISABLED
		ret.Status.Details = "locked"
	}

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

	if user.HasAlternateEmail() {
		ret.Emails = append(ret.Emails, &v2.UserTrait_Email{
			Address: user.GetAlternateEmail(),
		})
	}

	return ret, nil
}
