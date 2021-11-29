package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Root level middleware
func UseLogRec(e *echo.Echo) echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return *e
}

// Group level middleware
func BasicAuth(g *echo.Group, e *echo.Echo) {
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "Mohed" && password == "Test" {
			HandleFunc(e)
			return true, nil
		}
		return false, nil
	}))
}

// Root level middleware
func HandleFunc(e *echo.Echo) {
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)
}
