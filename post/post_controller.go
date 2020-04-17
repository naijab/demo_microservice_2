package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	if id == "1" {
		p := &Post{
			Id:     1,
			Title:  "Hello World",
			Detail: "Lorem 2021",
		}
		res := &HttpResponse{
			Message: "Successful",
			Data:    &p,
		}
		return c.JSON(http.StatusOK, res)
	} else {
		res := &HttpResponse{
			Message: "Post Not Found",
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, res)
	}
}
