package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/conductorone/baton-jumpcloud/pkg/client/extension"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi1"
	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi2"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"google.golang.org/grpc/codes"
)

type Client struct {
	_client1        *jcapi1.APIClient
	_client2        *jcapi2.APIClient
	extensionClient *uhttp.BaseHttpClient
	apiKey          string
	orgId           string
}

func NewClient(ctx context.Context, apiKey string, orgId string) (*Client, error) {
	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, nil), uhttp.WithUserAgent("baton-jumpcloud/0.1.0"))
	if err != nil {
		return nil, err
	}

	cc1 := jcapi1.NewConfiguration()
	cc1.HTTPClient = httpClient
	cc1.UserAgent = "baton-jumpcloud/0.1.0"

	cc2 := jcapi2.NewConfiguration()
	cc2.HTTPClient = httpClient
	cc2.UserAgent = "baton-jumpcloud/0.1.0"

	if orgId != "" {
		// optional, only needed by API keys linked to multi-tenant admins
		cc1.AddDefaultHeader("x-org-id", orgId)
		cc2.AddDefaultHeader("x-org-id", orgId)
	}

	client1 := jcapi1.NewAPIClient(cc1)
	client2 := jcapi2.NewAPIClient(cc2)

	baseHttpClient, err := uhttp.NewBaseHttpClientWithContext(ctx, httpClient)
	if err != nil {
		return nil, err
	}

	if client1 == nil || client2 == nil {
		return nil, uhttp.WrapErrors(
			codes.InvalidArgument,
			"client initialization failed: API client not configured",
			fmt.Errorf("API clients not properly initialized during connector setup"),
		)
	}

	return &Client{
		_client1:        client1,
		_client2:        client2,
		extensionClient: baseHttpClient,
		apiKey:          apiKey,
		orgId:           orgId,
	}, nil
}

func (jc *Client) client1(ctx context.Context) (context.Context, *jcapi1.APIClient) {
	return context.WithValue(ctx, jcapi1.ContextAPIKeys, map[string]jcapi1.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client1
}

func (jc *Client) client2(ctx context.Context) (context.Context, *jcapi2.APIClient) {
	return context.WithValue(ctx, jcapi2.ContextAPIKeys, map[string]jcapi2.APIKey{
		"x-api-key": {
			Key: jc.apiKey,
		},
	}), jc._client2
}

func (jc *Client) ListDirectories(ctx context.Context, opts *Options) ([]jcapi2.Directory, error) {
	ctx, client := jc.client2(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	directories, resp, err := client.DirectoriesApi.DirectoriesList(ctx).Limit(limit).Skip(page).Execute()
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		code := httpStatusToGRPCCode(resp.StatusCode)
		err := uhttp.WrapErrors(
			code,
			fmt.Sprintf("list directories: unexpected status code %d", resp.StatusCode),
			nil,
		)
		return nil, err
	}

	return directories, nil
}

func (jc *Client) GetUserByID(ctx context.Context, userID string) (*jcapi1.Userreturn, *http.Response, error) {
	userGetRequest := &extension.UserGetRequest{
		Client: jc.extensionClient,
		ApiKey: jc.apiKey,
		OrgId:  jc.orgId,
		UserID: userID,
	}

	user, resp, err := userGetRequest.Execute(ctx)
	if err != nil {
		return nil, resp, err
	}

	defer resp.Body.Close()

	return user, resp, nil
}

func (jc *Client) ListSystemUsers(ctx context.Context, opts *Options) ([]jcapi1.Systemuserreturn, *http.Response, string, error) {
	ctx, client := jc.client1(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	systemUsers, resp, err := client.SystemusersApi.SystemusersList(ctx).Skip(page).Limit(limit).Execute()
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list users")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(systemUsers.Results), page)

	return systemUsers.Results, resp, pageToken, nil
}

func (jc *Client) ListAdminUsers(ctx context.Context, opts *Options) ([]jcapi1.Userreturn, *http.Response, string, error) {
	page := opts.getPage()
	userListRequest := &extension.UserListRequest{
		Client: jc.extensionClient,
		ApiKey: jc.apiKey,
		OrgId:  jc.orgId,
		Page:   page,
	}

	users, resp, err := userListRequest.Skip(page).Execute(ctx)
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list admin users")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(users), page)

	return users, resp, pageToken, nil
}

func (jc *Client) GetSystemUserByID(ctx context.Context, userID string) (*jcapi1.Systemuserreturn, error) {
	ctx, client := jc.client1(ctx)

	user, resp, err := client.SystemusersApi.SystemusersGet(ctx, userID).Execute()
	if err != nil {
		return nil, wrapSDKError(err, resp, "failed to fetch manager details")
	}
	defer resp.Body.Close()

	return user, nil
}

func (jc *Client) ListGroups(ctx context.Context, opts *Options) ([]jcapi2.UserGroup, *http.Response, string, error) {
	ctx, client := jc.client2(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	groups, resp, err := client.UserGroupsApi.GroupsUserList(ctx).Skip(page).Limit(limit).Execute()
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list groups")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(groups), page)
	return groups, resp, pageToken, nil
}

func (jc *Client) ListGroupMembers(ctx context.Context, groupID string, opts *Options) ([]jcapi2.GraphConnection, *http.Response, string, error) {
	ctx, client := jc.client2(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	members, resp, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersList(ctx, groupID).Skip(page).Limit(limit).Execute()
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list group members")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(members), page)
	return members, resp, pageToken, nil
}

func (jc *Client) AddGroupMember(ctx context.Context, groupID string, memberID string) (*http.Response, error) {
	ctx, client := jc.client2(ctx)

	resp, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersPost(ctx, groupID).Body(jcapi2.GraphOperationUserGroupMember{
		Id:   memberID,
		Op:   "add",
		Type: "user",
	}).Execute()

	if err != nil {
		return nil, wrapSDKError(err, resp, "failed to add group member")
	}
	defer resp.Body.Close()

	return resp, nil
}

func (jc *Client) RemoveGroupMember(ctx context.Context, groupID string, memberID string) (*http.Response, error) {
	ctx, client := jc.client2(ctx)

	resp, err := client.UserGroupMembersMembershipApi.GraphUserGroupMembersPost(ctx, groupID).Body(jcapi2.GraphOperationUserGroupMember{
		Id:   memberID,
		Op:   "remove",
		Type: "user",
	}).Execute()
	if err != nil {
		return nil, wrapSDKError(err, resp, "failed to remove group member")
	}
	defer resp.Body.Close()

	return resp, nil
}

func (jc *Client) ListApplications(ctx context.Context, opts *Options) ([]jcapi1.Application, *http.Response, string, error) {
	ctx, client := jc.client1(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	applications, resp, err := client.ApplicationsApi.ApplicationsList(ctx).Skip(page).Limit(limit).Execute()
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list applications")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(applications.Results), page)
	return applications.Results, resp, pageToken, nil
}

func (jc *Client) ListApplicationAssociations(ctx context.Context, applicationID string, opts *Options) ([]jcapi2.GraphConnection, *http.Response, string, error) {
	ctx, client := jc.client2(ctx)

	limit := opts.getLimit()
	page := opts.getPage()
	targets := opts.getTargets()
	associations, resp, err := client.ApplicationsApi.GraphApplicationAssociationsList(ctx, applicationID).Skip(page).Limit(limit).Targets(targets).Execute()
	if err != nil {
		return nil, nil, "", wrapSDKError(err, resp, "failed to list application associations")
	}
	defer resp.Body.Close()

	pageToken := getNextPageToken(len(associations), page)
	return associations, resp, pageToken, nil
}
