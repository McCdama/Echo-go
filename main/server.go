package main

import (
	h "Echo-go/handle"
	l "Echo-go/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	l.UseLogRec(e)

	// http://localhost:1323/
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Salut!")
	})
	// http://localhost:1323/users/Mohed
	e.GET("/users/:id", h.GetUser)

	// http://localhost:1323/show?team=Peaky_Blinders&leader=Tommy_Shelby
	e.GET("/show", h.Show)

	// curl(.exe) -F "name=Mohed Rah" -F "email=mccdama@gmail.com" http://localhost:1323/save
	e.POST("/save", h.Save)

	// curl(.exe) -F "name=Mohed Rah" -F "avatar=@avatar\favicon.ico" http://localhost:1323/savedata
	e.POST("/savedata", h.SaveData)

	// curl(.exe) -v -F "name=Mohed Rah" -F "email=mccdama@gmail.com" http://localhost:1323/users
	e.POST("/users", h.Payload)

	// Bind JUST the Body
	// THIS SHOULD NOT WORK
	// curl -v -X POST http://localhost:1323/usersBody -H 'Content-Type: application/json' -F 'name=Mohed Rah' -F 'email=mccdama@gmail.com'

	// --> This works
	// curl -v -X POST http://localhost:1323/usersBody -H 'Content-Type: application/json' -d '{"date":"2021-11-30"}'
	e.POST("/usersBody", h.CreateDiary)

	// // http://localhost:1323/users
	// g := e.Group("/admin")
	// m.BasicAuth(g, e)

	// go run main/server.go --> Listening on http://localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}
