package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Server() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hey there!")
	})
	e.Logger.Fatal(e.Start(":8008"))
}
