package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
)

// RoleResourceTypeID is the resource type ID for roles, exported so callers
// (e.g. cmd/main.go) can gate cross-type grant emission on
// cli.ConnectorOpts.WillSyncResourceType(RoleResourceTypeID) without
// duplicating the string literal.
const RoleResourceTypeID = "role"

var (
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
	resourceTypeRole = &v2.ResourceType{
		Id:          RoleResourceTypeID,
		DisplayName: "Role",
		Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
		Annotations: annotations.New(&v2.SkipGrants{}),
	}
)
