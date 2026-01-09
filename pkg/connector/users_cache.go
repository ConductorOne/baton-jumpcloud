package connector

import (
	"context"

	"github.com/conductorone/baton-jumpcloud/pkg/connector/client"
	"github.com/conductorone/baton-jumpcloud/pkg/connector/client/jcapi1"
	"github.com/conductorone/baton-sdk/pkg/session"
	"github.com/conductorone/baton-sdk/pkg/types/sessions"
)

type usersCache struct {
	client *client.Client
}

func newUsersCache(jumpcloudClient *client.Client) *usersCache {
	return &usersCache{
		client: jumpcloudClient,
	}
}

func (uc *usersCache) GetSystemUsersByEmailList(ctx context.Context, sessionStorage sessions.SessionStore, emails []string) (map[string]*jcapi1.Systemuserreturn, error) {
	systemUsers, err := session.GetManyJSON[*jcapi1.Systemuserreturn](ctx, sessionStorage, emails)
	if err != nil {
		return nil, err
	}
	return systemUsers, nil
}

func (uc *usersCache) SetSystemUsers(ctx context.Context, sessionStorage sessions.SessionStore, systemUsers []jcapi1.Systemuserreturn) error {
	if len(systemUsers) == 0 {
		return nil
	}

	const batchSize = 100

	// Process in batches of 100
	for i := 0; i < len(systemUsers); i += batchSize {
		end := i + batchSize
		if end > len(systemUsers) {
			end = len(systemUsers)
		}

		batch := systemUsers[i:end]
		systemUsersMap := make(map[string]*jcapi1.Systemuserreturn, len(batch))

		// Create map with email as key and pointer to user
		// Note: we need to take address of the slice element, not the loop variable
		for j := range batch {
			systemUsersMap[batch[j].GetEmail()] = &batch[j]
		}

		if err := session.SetManyJSON(ctx, sessionStorage, systemUsersMap); err != nil {
			return err
		}
	}

	return nil
}
