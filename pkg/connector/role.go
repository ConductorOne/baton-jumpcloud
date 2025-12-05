package connector

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
)

type roleResourceType struct {
	resourceType *v2.ResourceType
	client       jc1Func
	ext          *ExtensionClient
}

func (o *roleResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *roleResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Resource, *sdkResources.SyncOpResults, error) {
	var rv []*v2.Resource

	var annos annotations.Annotations
	annos.Update(&v2.RoleTrait{})

	for _, role := range jcRoleNames {
		rv = append(rv, &v2.Resource{
			Id: &v2.ResourceId{
				ResourceType: "role",
				Resource:     fmtRoleNameAsID(role),
			},
			DisplayName: role,
			Annotations: annos,
		})
	}
	return rv, nil, nil
}

// Jumpcloud doens't seem to publish API defs of Roles - just their string names in the support docs.
// https://support.jumpcloud.com/s/article/JumpCloud-Roles
var jcRoleNames = []string{
	"Administrator with Billing",
	"Administrator",
	"Manager",
	"Command Runner with Billing",
	"Command Runner",
	"Help Desk",
	"Read Only",
}

var replaceID = regexp.MustCompile("[^a-z0-9]+")

func fmtRoleNameAsID(roleName string) string {
	return "role_" + replaceID.ReplaceAllString(strings.ToLower(strings.TrimSpace(roleName)), "_")
}

func (o *roleResourceType) Entitlements(
	ctx context.Context,
	resource *v2.Resource,
	opts sdkResources.SyncOpAttrs,
) ([]*v2.Entitlement, *sdkResources.SyncOpResults, error) {
	var rv []*v2.Entitlement

	rv = append(rv, roleEntitlement(resource))

	return rv, nil, nil
}

func roleEntitlement(resource *v2.Resource) *v2.Entitlement {
	return &v2.Entitlement{
		Id:          fmtResource(resource.Id, resource.Id.GetResource()),
		Resource:    resource,
		DisplayName: fmt.Sprintf("%s Role Member", resource.DisplayName),
		Description: fmt.Sprintf("Member of %s role", resource.DisplayName),
		GrantableTo: []*v2.ResourceType{resourceTypeUser},
		Purpose:     v2.Entitlement_PURPOSE_VALUE_PERMISSION,
		Slug:        "member",
	}
}

func (o *roleResourceType) Grants(
	_ context.Context,
	_ *v2.Resource,
	_ sdkResources.SyncOpAttrs,
) ([]*v2.Grant, *sdkResources.SyncOpResults, error) {
	// Grants for roles are now discovered from User.Grants() to avoid caching all users.
	// This improves performance and memory usage for large organizations.
	return nil, nil, nil
}

func newRoleBuilder(client jc1Func, ext *ExtensionClient) *roleResourceType {
	return &roleResourceType{
		resourceType: resourceTypeRole,
		client:       client,
		ext:          ext,
	}
}
