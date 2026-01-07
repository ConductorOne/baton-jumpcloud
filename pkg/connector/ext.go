package connector

import "github.com/conductorone/baton-sdk/pkg/uhttp"

type ExtensionClient struct {
	client *uhttp.BaseHttpClient
	apiKey string
	orgId  string
}

func (ec *ExtensionClient) UserList() *UserListRequest {
	return &UserListRequest{
		client: ec.client,
		apiKey: ec.apiKey,
		orgId:  ec.orgId,
	}
}

func (ec *ExtensionClient) UserGet(userID string) *UserGetRequest {
	return &UserGetRequest{
		client: ec.client,
		apiKey: ec.apiKey,
		orgId:  ec.orgId,
		userID: userID,
	}
}
