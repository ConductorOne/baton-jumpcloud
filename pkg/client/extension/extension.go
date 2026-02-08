package extension

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/conductorone/baton-jumpcloud/pkg/client/jcapi1"
	"github.com/conductorone/baton-sdk/pkg/uhttp"
)

const defaultBaseURL = "https://console.jumpcloud.com"

type UserListRequest struct {
	Client  *uhttp.BaseHttpClient
	ApiKey  string
	OrgId   string
	BaseURL string
	Page    int32
}

type listUserResponse struct {
	TotalCount int64               `json:"totalCount"`
	Results    []jcapi1.Userreturn `json:"results"`
}

func (ulr *UserListRequest) Skip(skip int32) *UserListRequest {
	ulr.Page = skip
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
	if ulr.Page != 0 {
		qp.Set("skip", strconv.FormatInt(int64(ulr.Page), 10))
	}

	baseURL := ulr.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return nil, nil, err
	}

	u := &url.URL{
		Scheme:   parsedBase.Scheme,
		Host:     parsedBase.Host,
		Path:     "/api/users",
		RawQuery: qp.Encode(),
	}

	var reqOpts []uhttp.RequestOption
	reqOpts = append(reqOpts, uhttp.WithAcceptJSONHeader())
	reqOpts = append(reqOpts, uhttp.WithHeader("x-api-key", ulr.ApiKey))
	if ulr.OrgId != "" {
		reqOpts = append(reqOpts, uhttp.WithHeader("x-org-id", ulr.OrgId))
	}

	req, err := ulr.Client.NewRequest(ctx, http.MethodGet, u, reqOpts...)
	if err != nil {
		return nil, nil, err
	}

	rv := &listUserResponse{}
	resp, err := ulr.Client.Do(req, uhttp.WithJSONResponse(rv))
	if err != nil {
		return nil, resp, err
	}

	return rv.Results, resp, nil
}

type UserGetRequest struct {
	Client  *uhttp.BaseHttpClient
	ApiKey  string
	OrgId   string
	BaseURL string
	UserID  string
}

// Execute fetches an admin user by ID.
// Uses uhttp.Do() with WithJSONResponse for automatic JSON unmarshaling and error handling.
func (ugr *UserGetRequest) Execute(ctx context.Context) (*jcapi1.Userreturn, *http.Response, error) {
	baseURL := ugr.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return nil, nil, err
	}

	u := &url.URL{
		Scheme: parsedBase.Scheme,
		Host:   parsedBase.Host,
		Path:   "/api/users/" + url.PathEscape(ugr.UserID),
	}

	var reqOpts []uhttp.RequestOption
	reqOpts = append(reqOpts, uhttp.WithAcceptJSONHeader())
	reqOpts = append(reqOpts, uhttp.WithHeader("x-api-key", ugr.ApiKey))
	if ugr.OrgId != "" {
		reqOpts = append(reqOpts, uhttp.WithHeader("x-org-id", ugr.OrgId))
	}

	req, err := ugr.Client.NewRequest(ctx, http.MethodGet, u, reqOpts...)
	if err != nil {
		return nil, nil, err
	}

	rv := &jcapi1.Userreturn{}
	resp, err := ugr.Client.Do(req, uhttp.WithJSONResponse(rv))
	if err != nil {
		return nil, resp, err
	}

	return rv, resp, nil
}
