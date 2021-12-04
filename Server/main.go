package server

import (
	a "Echo-go/admin"
	c "Echo-go/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Defining of the admin router group
	adminGroup := e.Group("/admin")

	// http://localhost:8008/admin
	// Router for "/admin" path
	adminGroup.GET("", a.Admin())

	e.GET("/user/signin", c.SignInForm()).Name = "userSignInForm"
	e.POST("/user/signin", c.SignIn())

	e.Logger.Fatal(e.Start(":8008"))
}
