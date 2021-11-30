package handlers

import (
	"Echo-go/models"
	m "Echo-go/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, m.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task models.Task
		// Map incoming JSON body to the new task
		c.Bind(&task)
		// Add task
		id, err := m.PutTask(db, task.Email)
		// Return JSON
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
				"email":   task.Email,
			})
		} else {
			return err
		}

	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := m.DeleteTask(db, id)
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
