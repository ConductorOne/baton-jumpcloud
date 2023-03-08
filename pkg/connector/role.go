package connector

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/ConductorOne/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

type roleResourceType struct {
	resourceType *v2.ResourceType
	ext          *ExtensionClient

	allUsers        []jcapi1.Userreturn
	fetchedAllUsers bool
	allUsersMtx     sync.Mutex
}

func (o *roleResourceType) ResourceType(_ context.Context) *v2.ResourceType {
	return o.resourceType
}

func (o *roleResourceType) List(
	ctx context.Context,
	resourceID *v2.ResourceId,
	token *pagination.Token,
) ([]*v2.Resource, string, annotations.Annotations, error) {
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
	return rv, "", nil, nil
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
	token *pagination.Token,
) ([]*v2.Entitlement, string, annotations.Annotations, error) {
	var rv []*v2.Entitlement

	rv = append(rv, roleEntitlement(ctx, resource))

	return rv, "", nil, nil
}

func roleEntitlement(ctx context.Context, resource *v2.Resource) *v2.Entitlement {
	return &v2.Entitlement{
		Id:          fmtResource(resource.Id, resource.Id.GetResource()),
		Resource:    resource,
		DisplayName: fmt.Sprintf("%s Role Member", resource.DisplayName),
		Description: fmt.Sprintf("Member of %s role", resource.DisplayName),
		GrantableTo: []*v2.ResourceType{resourceTypeAdminUser},
		Purpose:     v2.Entitlement_PURPOSE_VALUE_PERMISSION,
		Slug:        resource.DisplayName,
	}
}

func (o *roleResourceType) cacheAllUsers(ctx context.Context) ([]jcapi1.Userreturn, error) {
	o.allUsersMtx.Lock()
	defer o.allUsersMtx.Unlock()
	if o.fetchedAllUsers {
		return o.allUsers, nil
	}

	skip := int32(0)
	rv := []jcapi1.Userreturn{}
	for {
		users, resp, err := o.ext.UserList().Skip(skip).Execute(ctx)
		if err != nil {
			return nil, err
		}
		_ = resp.Body.Close()

		rv = append(rv, users...)
		if len(users) == 0 {
			break
		}
		skip += int32(len(users))
	}
	o.allUsers = rv
	o.fetchedAllUsers = true
	return rv, nil
}

func (o *roleResourceType) Grants(
	ctx context.Context,
	resource *v2.Resource,
	token *pagination.Token,
) ([]*v2.Grant, string, annotations.Annotations, error) {
	users, err := o.cacheAllUsers(ctx)
	if err != nil {
		return nil, "", nil, err
	}

	var rv []*v2.Grant
	for i := range users {
		user := &users[i]
		roleID := fmtRoleNameAsID(user.GetRoleName())
		if resource.Id.Resource != roleID {
			continue
		}
		rv = append(rv, roleGrant(resource, resourceTypeAdminUser.Id, user))
	}
	return rv, "", nil, nil
}

func roleGrant(resource *v2.Resource, resoureTypeId string, user *jcapi1.Userreturn) *v2.Grant {
	roleID := resource.Id.GetResource()
	ur := &v2.Resource{Id: &v2.ResourceId{ResourceType: resoureTypeId, Resource: user.GetId()}}

	var annos annotations.Annotations

	return &v2.Grant{
		Id: fmtResourceGrant(resource.Id, ur.Id, roleID),
		Entitlement: &v2.Entitlement{
			Id:       fmtResource(resource.Id, roleID),
			Resource: resource,
		},
		Annotations: annos,
		Principal:   ur,
	}
}

func newRoleBuilder(ext *ExtensionClient) *roleResourceType {
	return &roleResourceType{
		resourceType: resourceTypeRole,
		ext:          ext,
	}
}
