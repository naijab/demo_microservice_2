package post

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetPost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/post/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	res := rec.Result()
	defer res.Body.Close()

	if assert.NoError(t, GetPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
