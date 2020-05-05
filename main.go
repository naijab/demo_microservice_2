package main

import (
	"core/config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"post_service/post"
	"post_service/post/model"
	"post_service/post/repository"
	"post_service/post/service"
)

type Config struct {
	Database *gorm.DB
}

func main() {
	c := Config{}
	c.initDatabase()
	c.initEcho()
}

func (c *Config) initDatabase() {
	db := config.DatabaseConnection{}
	if err := db.NewDatabaseConnection(); err == nil {
		c.Database = db.Database
		c.Database.AutoMigrate(&model.Post{})
	}
}

func (c *Config) initEcho() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Repository
	postRepo := repository.NewPostRepository(c.Database)

	// Service
	postService := service.NewPostService(postRepo)

	// Controller
	postCon := post.NewPostController(postService, postRepo)

	// Routes
	e.GET("/api/post", postCon.GetAllPost)
	e.GET("/api/post/:id", postCon.GetPostById)

	// Start server
	e.Logger.Fatal(e.Start(":1332"))
}