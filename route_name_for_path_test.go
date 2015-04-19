package routes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteNameForPath(t *testing.T) {
	tests := []struct {
		path string
		name string
	}{
		{"/", "Root"},
		{"/users", "Users"},
		{"/user/:id", "UserByID"},
		{"/user/:id/info", "UserInfoByID"},
		{"/organization/:organization_id/user/:user_id", "OrganizationUserByOrganizationIDAndUserID"},
		{"/api/v1/login", "APIV1Login"},
		{"/api/v1/oauth_token", "APIV1OAuthToken"},
	}

	for _, tst := range tests {
		assert.Equal(t, tst.name, RouteNameForPath(tst.path), "for %q", tst.path)
	}
}
