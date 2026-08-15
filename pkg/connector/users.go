package connector

import (
	"context"
	"fmt"
	"strings"

	"github.com/conductorone/baton-jumpcloud/pkg/client"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

type userResourceType struct {
	resourceType *v2.ResourceType
	client       *client.Client
	usersCache   *usersCache
	managers     map[string]*jcapi1.Systemuserreturn
}

func (o *userResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func newUserBuilder(client *client.Client, syncRoles bool) *userResourceType {
	resourceType := resourceTypeUser
	if !syncRoles {
		// The user builder has no entitlements or grants of its own -- its only
		// Grants() output is the cross-type role grant gated below. When roles
		// aren't being synced, skip entitlement/grant discovery for users entirely.
		rt := proto.Clone(resourceTypeUser).(*v2.ResourceType)
		rt.Annotations = annotations.New(&v2.SkipEntitlementsAndGrants{})
		resourceType = rt
	}

	return &userResourceType{
		resourceType: resourceType,
		client:       client,
		managers:     make(map[string]*jcapi1.Systemuserreturn),
		usersCache:   newUsersCache(client),
	}
}

func (o *userResourceType) Entitlements(_ context.Context, _ *v2.Resource, _ sdkResources.SyncOpAttrs) ([]*v2.Entitlement, *sdkResources.SyncOpResults, error) {
	return nil, nil, nil
}

// Grants emits the cross-type role grant. There is no syncRoles guard here:
// when roles aren't being synced, newUserBuilder annotates the user resource
// type SkipEntitlementsAndGrants and the SDK never calls Grants() at all
// (shouldSkipGrants -> shouldSkipEntitlementsAndGrants in the SDK's
// pkg/sync/syncer.go), so a guard would be unreachable.
func (o *userResourceType) Grants(ctx context.Context, resource *v2.Resource, _ sdkResources.SyncOpAttrs) ([]*v2.Grant, *sdkResources.SyncOpResults, error) {
	userID := resource.Id.Resource
	// Only admin users have role grants. System users won't be found in the admin users endpoint.
	adminUser, err := o.client.GetUserByID(ctx, userID)
	if err != nil && !strings.Contains(err.Error(), codes.NotFound.String()) {
		return nil, nil, fmt.Errorf("failed to get user for grant discovery: %w", err)
	}

	roleName := adminUser.GetRoleName()
	if roleName == "" {
		return nil, nil, nil
	}

	roleID := fmtRoleNameAsID(roleName)
	roleRes := &v2.Resource{
		Id: &v2.ResourceId{
			ResourceType: resourceTypeRole.Id,
			Resource:     roleID,
		},
		DisplayName: roleName,
	}

	// Create grant for the role entitlement.
	var annos annotations.Annotations
	grant := &v2.Grant{
		Id: fmtResourceGrant(roleRes.Id, resource.Id, roleID),
		Entitlement: &v2.Entitlement{
			Id:       fmtResource(roleRes.Id, roleID),
			Resource: roleRes,
		},
		Annotations: annos,
		Principal:   resource,
	}

	return []*v2.Grant{grant}, nil, nil
}

func (o *userResourceType) List(ctx context.Context, parentResourceID *v2.ResourceId, opts sdkResources.SyncOpAttrs) ([]*v2.Resource, *sdkResources.SyncOpResults, error) {
	l := ctxzap.Extract(ctx)

	skip, b, err := unmarshalSkipToken(&opts.PageToken)
	if err != nil {
		return nil, nil, err
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
		options := &client.Options{}
		systemUsers, nextPageToken, err := o.client.ListSystemUsers(ctx, options.WithSkip(skip))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list system users: %w", err)
		}

		// populate the users cache with the system users
		err = o.usersCache.SetSystemUsers(ctx, opts.Session, systemUsers)
		if err != nil {
			return nil, nil, err
		}

		for i := range systemUsers {
			ur, err := o.userResource(ctx, &systemUsers[i])
			if err != nil {
				return nil, nil, err
			}
			rv = append(rv, ur)
		}
		pageToken, err = marshalSkipToken(nextPageToken, b)
		if err != nil {
			return nil, nil, err
		}
	case "list-admin-users":
		options := &client.Options{}
		adminUsers, nextPageToken, err := o.client.ListAdminUsers(ctx, options.WithSkip(skip))
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list admin users: %w", err)
		}

		adminEmails := make([]string, len(adminUsers))
		for i := range adminUsers {
			adminEmails[i] = adminUsers[i].GetEmail()
		}
		systemUsersMap, err := o.usersCache.GetSystemUsersByEmailList(ctx, opts.Session, adminEmails)
		if err != nil {
			return nil, nil, err
		}
		for i := range adminUsers {
			adminEmail := adminUsers[i].GetEmail()
			adminUser, err := o.adminUserResource(&adminUsers[i])
			if err != nil {
				return nil, nil, err
			}

			// Check if the admin user is also a system user, if so we'll use that user instead
			if systemUser, ok := systemUsersMap[adminEmail]; ok {
				if systemUser != nil {
					continue
				} else {
					return nil, nil, err
				}
			}

			l.Debug("admin user not found as system user, creating", zap.String("email", adminEmail))
			rv = append(rv, adminUser)
		}
		pageToken, err = marshalSkipToken(nextPageToken, b)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("unknown page state encountered during user list pagination: %s", b.Current().ResourceTypeID)
	}

	return rv, &sdkResources.SyncOpResults{NextPageToken: pageToken}, nil
}

func (o *userResourceType) adminUserResource(user *jcapi1.Userreturn) (*v2.Resource, error) {
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
		return nil, fmt.Errorf("failed to create resource for admin user: %w", err)
	}

	return r, nil
}

func (o *userResourceType) userResource(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.Resource, error) {
	trait, err := o.userTrait(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to get user trait during user resource creation: %w", err)
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

func (o *userResourceType) userTrait(ctx context.Context, user *jcapi1.Systemuserreturn) (*v2.UserTrait, error) {
	profile, err := structpb.NewStruct(map[string]interface{}{
		"id": user.GetId(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to construct user profile during user trait creation: %w", err)
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
		// TODO (luisina) > use a cache to store manager by ID, to avoid making multiple requests to the API
		manager, err := o.client.GetSystemUserByID(ctx, managerID)
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
