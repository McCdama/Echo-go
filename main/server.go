package main

import (
	h "Echo-go/handle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Salut! ")
	})
	// http://localhost:1323/users/Mohed
	e.GET("/users/:id", h.GetUser)

	// go run main/server.go --> Listening on http://localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}
