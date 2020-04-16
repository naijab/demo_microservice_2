package main

import (
	"post_service/post"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/post", post.GetPost)

	// Start server
	e.Logger.Fatal(e.Start(":1332"))
}
