package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
)

// RoleResourceTypeID must equal resourceTypeRole.Id below. Exported so the connector
// can gate cross-type role-grant emission on WillSyncResourceType without duplicating
// the string literal.
const RoleResourceTypeID = "role"

var (
	// newUserBuilder clones this and adds SkipEntitlements, or
	// SkipEntitlementsAndGrants when role isn't synced.
	resourceTypeUser = &v2.ResourceType{
		Id:          "user",
		DisplayName: "User",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
		Description: "JumpCloud User: The User account is the core identity for your employees, and is the account type that is used to authenticate against resources",
	}
	resourceTypeGroup = &v2.ResourceType{
		Id:          "group",
		DisplayName: "Group",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_GROUP},
	}
	resourceTypeApp = &v2.ResourceType{
		Id:          "app",
		DisplayName: "App",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_APP},
	}
	// resourceTypeRole skips Grants(): roleBuilder.Grants() is a no-op because the role
	// grants are emitted from the user builder, which reads each admin user's role in
	// the pass it already makes rather than re-scanning every user once per role.
	resourceTypeRole = &v2.ResourceType{
		Id:          RoleResourceTypeID,
		DisplayName: "Role",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
		Annotations: annotations.New(&v2.SkipGrants{}),
	}
)
