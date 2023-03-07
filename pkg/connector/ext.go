package connector

import "net/http"

type ExtensionClient struct {
	client *http.Client
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
