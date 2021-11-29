package main

import (
	h "Echo-go/handlers"
	s "Echo-go/storage"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/vue-app")
	s.Init(db, err)

	e.File("/", "public/index.html")

	// http://localhost:8008/tasks
	e.GET("/tasks", h.GetTasks(db))

	// http://localhost:8008/tasks
	// Postman, “PUT” -> click “Body” tab -> Check “raw” -> select JSON (application/json) as type
	e.PUT("/tasks", h.PutTask(db))

	e.DELETE("/tasks/:id", h.DeleteTask(db))

	// go run main/server.go --> Listening on http://localhost:8008
	e.Logger.Fatal(e.Start(":8008"))

}
