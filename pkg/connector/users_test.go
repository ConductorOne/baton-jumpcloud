package connector

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/conductorone/baton-jumpcloud/pkg/client"
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
	"github.com/conductorone/baton-sdk/pkg/annotations"
	sdkResources "github.com/conductorone/baton-sdk/pkg/types/resource"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// TestUserGrants_RoleSyncFilter covers ConductorOne/baton-linear#55: the user
// builder emits role grants as a sync optimization, but must not do so when
// the customer's sync filter excludes the role resource type.
func TestUserGrants_RoleSyncFilter(t *testing.T) {
	userResource := &v2.Resource{
		Id: fmtResourceId(resourceTypeUser.Id, "user-1"),
	}

	t.Run("role type filtered out -> no grants, no API call", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer srv.Close()

		c, err := client.NewClient(context.Background(), "api-key", "", srv.URL)
		require.NoError(t, err)

		builder := newUserBuilder(c, false)
		require.False(t, builder.syncRoles)

		grants, _, err := builder.Grants(context.Background(), userResource, sdkResources.SyncOpAttrs{})
		require.NoError(t, err)
		require.Empty(t, grants)
		require.False(t, called, "role grant discovery must not call the API when roles are filtered out")
	})

	t.Run("role type synced -> role grant emitted", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{
				"id":       "user-1",
				"roleName": "Administrator",
			})
		}))
		defer srv.Close()

		c, err := client.NewClient(context.Background(), "api-key", "", srv.URL)
		require.NoError(t, err)

		builder := newUserBuilder(c, true)
		require.True(t, builder.syncRoles)

		grants, _, err := builder.Grants(context.Background(), userResource, sdkResources.SyncOpAttrs{})
		require.NoError(t, err)
		require.Len(t, grants, 1)
		require.Equal(t, fmtRoleNameAsID("Administrator"), grants[0].Entitlement.Resource.Id.Resource)
		require.Equal(t, resourceTypeRole.Id, grants[0].Entitlement.Resource.Id.ResourceType)
	})
}

// TestNewUserBuilder_ResourceTypeAnnotations covers Step 4 of the
// WillSyncResourceType gating pattern: when the user builder has no
// entitlements/grants of its own to offer, mark the emitted resource type
// with SkipEntitlementsAndGrants so the SDK doesn't bother syncing them.
func TestNewUserBuilder_ResourceTypeAnnotations(t *testing.T) {
	t.Run("role type filtered out -> SkipEntitlementsAndGrants set", func(t *testing.T) {
		builder := newUserBuilder(nil, false)
		var skip v2.SkipEntitlementsAndGrants
		ok, err := annotationsContain(builder.resourceType.GetAnnotations(), &skip)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("role type synced -> no SkipEntitlementsAndGrants", func(t *testing.T) {
		builder := newUserBuilder(nil, true)
		var skip v2.SkipEntitlementsAndGrants
		ok, err := annotationsContain(builder.resourceType.GetAnnotations(), &skip)
		require.NoError(t, err)
		require.False(t, ok)
	})
}

func annotationsContain(annos []*anypb.Any, msg proto.Message) (bool, error) {
	as := annotations.Annotations(annos)
	return as.Pick(msg)
}
