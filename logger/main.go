package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UseLogRec(e *echo.Echo) echo.Echo {
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return *e
}
