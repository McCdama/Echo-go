package handle

import (
	p "Echo-go/payload"
	"fmt"
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

func Payload(c echo.Context) error {
	u := new(p.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// OR
	// return c.XML(http.StatusCreated, u)
}

func PayloadBindBody(c echo.Context) error {
	var getData p.TestModel
	if err := (&echo.DefaultBinder{}).BindBody(c, &getData); err != nil {
		fmt.Print(err.Error())
	}
	return c.JSON(200, getData)
}

func PayloadBindPara(c echo.Context) error {
	// c.QueryParams() --> MAP
	var getData p.User
	e := c.QueryParam("email")
	n := c.QueryParam("name")
	if n == "" || e == "" {
		return c.JSON(200, getData.PanicDetail())
	} else {
		if err := (&echo.DefaultBinder{}).BindQueryParams(c, &getData); err != nil {
			panic(err)
		}
	}
	return c.JSON(200, getData.GetDetail())
}
