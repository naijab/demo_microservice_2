module post_service

go 1.14

require (
	core v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.12
	github.com/labstack/echo/v4 v4.1.16
	github.com/stretchr/testify v1.4.0
)

replace core => ../core
