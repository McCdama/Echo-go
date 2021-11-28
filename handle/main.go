package handle

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func Show(c echo.Context) error {
	team := c.QueryParam("team")
	leader := c.QueryParam("leader")
	return c.String(http.StatusOK, "Team: "+team+", Leader: "+leader)
}