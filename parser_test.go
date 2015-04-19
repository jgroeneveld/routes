package routes

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRoutes(t *testing.T) {
	input := `
		import "github.com/jgroeneveld/myproj/users"

		GET		/api/v1/users		users.GetUsers
		POST	/api/v1/users 		users.PostUser

		// single users
		get		/api/v1/user/:id 	users.GetUserForID
	`

	rd, err := ParseRoutes(strings.NewReader(input))
	require.NoError(t, err)

	assert.Equal(
		t,
		[]Import{
			"github.com/jgroeneveld/myproj/users",
		},
		rd.Imports,
	)

	assert.Equal(
		t,
		[]Route{
			{"GET", "/api/v1/users", "users.GetUsers"},
			{"POST", "/api/v1/users", "users.PostUser"},
			{"GET", "/api/v1/user/:id", "users.GetUserForID"},
		},
		rd.Routes,
	)
}

func TestParseRoutes_MalformedImport(t *testing.T) {
	input := ` import "asd def" `

	_, err := ParseRoutes(strings.NewReader(input))

	assert.Error(t, err, "should detect malformed import")
}

func TestParseRoutes_MalformedRoute(t *testing.T) {
	input := ` GET /users `

	_, err := ParseRoutes(strings.NewReader(input))

	assert.Error(t, err, "should detect malformed route")
}

func TestDowncaseFirstCamel(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"idForUser", "idForUser"},
		{"IDForUser", "idForUser"},
		{"UserName", "userName"},
		{"OAuthToken", "oAuthToken"},
	}

	for _, tst := range tests {
		assert.Equal(t, tst.expected, downcaseFirstCamel(tst.in), "for %s", tst.in)
	}
}
