package connector

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/conductorone/baton-jumpcloud/pkg/jcapi1"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

var (
	userCache                sync.Map
	errUserNotFoundForEmail  = errors.New("user not found for email")
	errMultipleUsersForEmail = errors.New("multiple users found for email")
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

func fetchUserByEmail(ctx context.Context, client *jcapi1.APIClient, email string) (*jcapi1.Systemuserreturn, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if u, ok := userCache.Load(email); ok {
		return u.(*jcapi1.Systemuserreturn), nil
	}

	list, resp, err := client.SystemusersApi.SystemusersList(ctx).Filter(fmt.Sprintf("email:$eq:%s", email)).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if len(list.Results) == 0 {
		return nil, errUserNotFoundForEmail
	}

	if len(list.Results) != 1 {
		return nil, errMultipleUsersForEmail
	}

	userCache.Store(email, &list.Results[0])

	return &list.Results[0], nil
}
