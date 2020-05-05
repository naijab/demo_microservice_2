package post

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"post_service/post/repository"
	"post_service/post/service"
)

type IPostController interface {
	GetPostById(c echo.Context) error
	GetAllPost(c echo.Context) error
}

type Controller struct {
	postService service.IPostService
	postRepo    repository.IPostRepository
}

// Create New Instance of PostService
func NewPostController(postService service.IPostService, postRepo repository.IPostRepository) *Controller {
	return &Controller{postService, postRepo}
}

func (con *Controller) GetPostById(c echo.Context) error {
	id := c.Param("id")
	p, e := con.postService.GetById(id)
	if e != nil {
		res := &HttpResponse{
			Message: e.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, res)
	}
	res := &HttpResponse{
		Message: "Successful",
		Data:    &p,
	}
	return c.JSON(http.StatusOK, res)
}

func (con *Controller) GetAllPost(c echo.Context) error {
	p, e := con.postService.GetAll()
	if e != nil {
		res := &HttpResponse{
			Message: e.Error(),
			Data:    nil,
		}
		return c.JSON(http.StatusNotFound, res)
	}
	res := &HttpResponse{
		Message: "Successful",
		Data:    &p,
	}
	return c.JSON(http.StatusOK, res)
}

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
