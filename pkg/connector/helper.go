package connector

import v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"

func fmtResourceId(resourceTypeID string, id string) *v2.ResourceId {
	return &v2.ResourceId{
		ResourceType: resourceTypeID,
		Resource:     id,
	}
}
