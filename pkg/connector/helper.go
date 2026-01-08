package connector

import (
	"errors"
	"fmt"
	"sync"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

var (
	userCache                sync.Map
	errUserNotFoundForEmail  = errors.New("user not found for email")
	errMultipleUsersForEmail = errors.New("multiple users found for email")
)

func fmtResourceId(resourceTypeID string, id string) *v2.ResourceId {
	return &v2.ResourceId{
		ResourceType: resourceTypeID,
		Resource:     id,
	}
}

func fmtResource(resourceID *v2.ResourceId, _ string) string {
	return fmt.Sprintf(
		"%s:%s",
		resourceID.ResourceType,
		resourceID.Resource,
	)
}

func fmtResourceGrant(resourceID *v2.ResourceId, principalId *v2.ResourceId, permission string) string {
	return fmt.Sprintf(
		"%s-grant:%s:%s:%s:%s",
		resourceID.ResourceType,
		resourceID.Resource,
		principalId.ResourceType,
		principalId.Resource,
		permission,
	)
}
