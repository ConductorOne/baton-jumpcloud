package connector

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

type UserListRequest struct {
	client *uhttp.BaseHttpClient
	apiKey string
	orgId  string
	skip   int32
}

type listUserResponse struct {
	TotalCount int64               `json:"totalCount"`
	Results    []jcapi1.Userreturn `json:"results"`
}

func (ulr *UserListRequest) Skip(skip int32) *UserListRequest {
	ulr.skip = skip
	return ulr
}

func (ulr *UserListRequest) Execute(ctx context.Context) ([]jcapi1.Userreturn, *http.Response, error) {
	// curl --request PUT \
	// --url https://console.jumpcloud.com/api/users/{id} \
	// --header 'content-type: application/json' \
	// --header 'x-api-key: REPLACE_KEY_VALUE' \
	// --header 'x-org-id: ' \
	// --data '{"email":"user@example.com","enableMultiFactor":true,"firstname":"string","growthData":{},"lastWhatsNewChecked":"2019-08-24","lastname":"string","roleName":"string"}'
	//
	// JumpCloud doesn't export the List endpoint in their
	// OpenAPI spec, but they do export the PUT...
	// .... so.. here we go!

	qp := url.Values{}
	if ulr.skip != 0 {
		qp.Set("skip", strconv.FormatInt(int64(ulr.skip), 10))
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "console.jumpcloud.com",
		Path:     "/api/users",
		RawQuery: qp.Encode(),
	}

	var reqOpts []uhttp.RequestOption
	reqOpts = append(reqOpts, uhttp.WithAcceptJSONHeader())
	reqOpts = append(reqOpts, uhttp.WithHeader("x-api-key", ulr.apiKey))
	if ulr.orgId != "" {
		reqOpts = append(reqOpts, uhttp.WithHeader("x-org-id", ulr.orgId))
	}

	req, err := ulr.client.NewRequest(ctx, http.MethodGet, u, reqOpts...)
	if err != nil {
		return nil, nil, err
	}

	rv := &listUserResponse{}
	resp, err := ulr.client.Do(req, uhttp.WithJSONResponse(rv))
	if err != nil {
		return nil, resp, err
	}

	return rv.Results, resp, nil
}

type UserGetRequest struct {
	client *uhttp.BaseHttpClient
	apiKey string
	orgId  string
	userID string
}

// Execute fetches an admin user by ID.
// Uses uhttp.Do() with WithJSONResponse for automatic JSON unmarshaling and error handling.
func (ugr *UserGetRequest) Execute(ctx context.Context) (*jcapi1.Userreturn, *http.Response, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   "console.jumpcloud.com",
		Path:   "/api/users/" + url.PathEscape(ugr.userID),
	}

	var reqOpts []uhttp.RequestOption
	reqOpts = append(reqOpts, uhttp.WithAcceptJSONHeader())
	reqOpts = append(reqOpts, uhttp.WithHeader("x-api-key", ugr.apiKey))
	if ugr.orgId != "" {
		reqOpts = append(reqOpts, uhttp.WithHeader("x-org-id", ugr.orgId))
	}

	req, err := ugr.client.NewRequest(ctx, http.MethodGet, u, reqOpts...)
	if err != nil {
		return nil, nil, err
	}

	rv := &jcapi1.Userreturn{}
	resp, err := ugr.client.Do(req, uhttp.WithJSONResponse(rv))
	if err != nil {
		return nil, resp, err
	}

	return rv, resp, nil
}
