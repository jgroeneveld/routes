package routes

import (
	"bytes"
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	rd, err := ParseRoutes(strings.NewReader(routesFile))
	require.NoError(t, err)

	w := &bytes.Buffer{}
	err = Write("routes", rd, w)
	require.NoError(t, err)

	assert.Equal(t, expected, w.String())
}

const routesFile = `
import "hallo"
import "miau"

GET		/api/user/:id	web.ShowUser
PATCH	/api/user/:id	web.UpdateUser
POST	/api/users		checkWriteRightsMiddleware(web.CreateUser)
`

const expected = `package routes

import (
	"github.com/julienschmidt/httprouter"

	"hallo"
	"miau"
)

func Router() *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/user/:id", web.ShowUser)
	router.PATCH("/api/user/:id", web.UpdateUser)
	router.POST("/api/users", checkWriteRightsMiddleware(web.CreateUser))

	return router
}

func APIUserByIDPath(ID string) {}

func APIUsersPath() {}
`
