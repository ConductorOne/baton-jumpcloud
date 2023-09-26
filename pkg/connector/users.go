package connector

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

type userResourceType struct {
	resourceType *v2.ResourceType
	client1      jc1Func
	client2      jc2Func
	managers     map[string]*jcapi1.Systemuserreturn
	ext          *ExtensionClient
}

func (o *userResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func newUserBuilder(jc1 jc1Func, jc2 jc2Func, ext *ExtensionClient) *userResourceType {
	return &userResourceType{
		resourceType: resourceTypeUser,
		client1:      jc1,
		client2:      jc2,
		managers:     make(map[string]*jcapi1.Systemuserreturn),
		ext:          ext,
	}
}

func (o *userResourceType) Entitlements(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *userResourceType) Grants(_ context.Context, _ *v2.Resource, _ *pagination.Token) ([]*v2.Grant, string, annotations.Annotations, error) {
	return nil, "", nil, nil
}

func (o *userResourceType) List(ctx context.Context, parentResourceID *v2.ResourceId, pt *pagination.Token) ([]*v2.Resource, string, annotations.Annotations, error) {
	l := ctxzap.Extract(ctx)

	ctx, client := o.client1(ctx)

	skip, b, err := unmarshalSkipToken(pt)
	if err != nil {
		return nil, "", nil, err
	}

	if b.Current() == nil {
		// Push onto stack in reverse
		b.Push(pagination.PageState{
			ResourceTypeID: "list-admin-users",
		})
		b.Push(pagination.PageState{
			ResourceTypeID: "list-users",
		})
	}
	var rv []*v2.Resource
	var pageToken string
	switch b.Current().ResourceTypeID {
	case "list-users":
		list, resp, err := client.SystemusersApi.SystemusersList(ctx).Skip(skip).Execute()
		if err != nil {
			return nil, "", nil, err
		}
		defer resp.Body.Close()

		for i := range list.Results {
			ur, err := o.userResource(ctx, &list.Results[i])
			if err != nil {
				return nil, "", nil, err
			}
			rv = append(rv, ur)
		}
		pageToken, err = marshalSkipToken(len(list.Results), skip, b)
		if err != nil {
			return nil, "", nil, err
		}
	case "list-admin-users":
		adminUsers, resp, err := o.ext.UserList().Skip(skip).Execute(ctx)
		if err != nil {
			return nil, "", nil, err
		}
		defer resp.Body.Close()

		for i := range adminUsers {
			adminEmail := adminUsers[i].GetEmail()
			adminUser, err := o.adminUserResource(ctx, &adminUsers[i])
			if err != nil {
				return nil, "", nil, err
			}

			// Check if the admin user is also a system user, if so we'll use that user instead
			systemUser, err := fetchUserByEmail(ctx, client, adminEmail)
			if err != nil && !errors.Is(err, errUserNotFoundForEmail) {
				return nil, "", nil, err
			}

			if systemUser != nil {
				continue
			}

			l.Debug("admin user not found as system user, creating", zap.String("email", adminEmail))
			rv = append(rv, adminUser)
		}
		pageToken, err = marshalSkipToken(len(adminUsers), skip, b)
		if err != nil {
			return nil, "", nil, err
		}
	default:
		return nil, "", nil, fmt.Errorf("baton-jumpcloud: unknown page state: %s", b.Current().ResourceTypeID)
	}

	return rv, pageToken, nil, nil
}

func (o *userResourceType) adminUserResource(ctx context.Context, user *jcapi1.Userreturn) (*v2.Resource, error) {
	profile := map[string]interface{}{
		"id": user.GetId(),
	}

	if user.HasOrganization() {
		profile["organization"] = user.GetOrganization()
	}

	userTraitOps := []sdkResources.UserTraitOption{
		sdkResources.WithUserProfile(profile),
	}

	status := v2.UserTrait_Status_STATUS_ENABLED
	if user.GetSuspended() {
		status = v2.UserTrait_Status_STATUS_DISABLED
	}
	userTraitOps = append(userTraitOps, sdkResources.WithStatus(status))

	email := user.GetEmail()
	if email != "" {
		userTraitOps = append(userTraitOps, sdkResources.WithEmail(email, true))
	}

	r, err := sdkResources.NewUserResource(
		fmt.Sprintf("%s %s", user.GetFirstname(), user.GetLastname()),
		o.resourceType,
		user.GetId(),
		userTraitOps,
	)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (o *userResourceType) userResource(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.Resource, error) {
	trait, err := o.userTrait(ctx, user)
	if err != nil {
		return nil, err
	}
	var annos annotations.Annotations
	annos.Update(trait)

	return &v2.Resource{
		Id:          fmtResourceId(resourceTypeUser.Id, user.GetId()),
		DisplayName: o.userDisplayName(user),
		Annotations: annos,
	}, nil
}

func (o *userResourceType) userDisplayName(user *jcapi1.Systemuserreturn) string {
	if dn := user.GetDisplayname(); dn != "" {
		return dn
	}
	return fmt.Sprintf("%s %s", user.GetFirstname(), user.GetLastname())
}

func (o *userResourceType) fetchManager(ctx context.Context, managerID string) (*jcapi1.Systemuserreturn, error) {
	ctx, client := o.client1(ctx)
	l := ctxzap.Extract(ctx)

	manager, ok := o.managers[managerID]
	if ok {
		return manager, nil
	}

	m, resp, err := client.SystemusersApi.SystemusersGet(ctx, managerID).Execute()
	if err != nil {
		l.Error(
			"baton-jumpcloud: failed to fetch manager details",
			zap.Error(err),
			zap.String("manager_id", managerID),
		)
		return nil, err
	}
	defer resp.Body.Close()

	o.managers[managerID] = m
	return m, nil
}

func (o *userResourceType) userTrait(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.UserTrait, error) {
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
		managerID := user.GetManager()
		profile.Fields["manager_id"] = structpb.NewStringValue(managerID)
		manager, err := o.fetchManager(ctx, managerID)
		if err == nil && manager != nil {
			profile.Fields["manager"] = structpb.NewStringValue(manager.GetEmail())
		}
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
