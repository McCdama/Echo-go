package handle

import (
	"io"
	"net/http"
	"os"

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

func Save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "Name: "+name+", Email: "+email)
}

func SaveData(c echo.Context) error {
	// Get Name
	name := c.FormValue("name")

	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")

}
