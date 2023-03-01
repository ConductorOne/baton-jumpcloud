package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/pagination"
)

const defaultLimit = 100

func parsePageToken(token string, resourceID *v2.ResourceId) (*pagination.Bag, string, error) {
	b := &pagination.Bag{}
	err := b.Unmarshal(token)
	if err != nil {
		return nil, "", err
	}

	if b.Current() == nil {
		b.Push(pagination.PageState{
			ResourceTypeID: resourceID.ResourceType,
			ResourceID:     resourceID.Resource,
		})
	}

	page := b.PageToken()

	return b, page, nil
}

func newPaginationToken(limit int, nextPageToken string) *pagination.Token {
	if limit == 0 || limit > defaultLimit {
		limit = defaultLimit
	}

	return &pagination.Token{
		Size:  limit,
		Token: nextPageToken,
	}
}
