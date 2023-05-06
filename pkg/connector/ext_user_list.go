package connector

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
)

type UserListRequest struct {
	client *http.Client
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
	// OpenAPI spec, but they do epxort the PUT...
	// .... so.. here we go!

	qp := url.Values{}
	if ulr.skip != 0 {
		qp.Set("skip", strconv.FormatInt(int64(ulr.skip), 10))
	}

	u := url.URL{
		Scheme:   "https",
		Host:     "console.jumpcloud.com",
		Path:     "/api/users",
		RawQuery: qp.Encode(),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Accept", "application/json")
	if ulr.orgId != "" {
		req.Header.Set("x-org-id", ulr.orgId)
	}
	req.Header.Set("x-api-key", ulr.apiKey)
	req = req.WithContext(ctx)

	resp, err := ulr.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("jumpcloud: invalid response code '%d' to list users", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	rv := &listUserResponse{}
	err = json.Unmarshal(data, rv)
	if err != nil {
		return nil, nil, err
	}
	return rv.Results, resp, nil
}
