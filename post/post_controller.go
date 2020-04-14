package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPost(c echo.Context) error {
	p := &Post{
		Title:  "Hello World",
		Detail: "Lorem 2021",
	}
	return c.JSON(http.StatusOK, p)
}
