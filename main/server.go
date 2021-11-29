package main

import (
	_ "Echo-go/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// http://localhost:8008/tasks
	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })

	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })

	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// go run main/server.go --> Listening on http://localhost:8008
	e.Logger.Fatal(e.Start(":8008"))

}
