package connector

import (
	"fmt"

	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

func fmtResourceId(resourceTypeID string, id string) *v2.ResourceId {
	return &v2.ResourceId{
		ResourceType: resourceTypeID,
		Resource:     id,
	}
}

func fmtResource(resourceID *v2.ResourceId, role string) string {
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
