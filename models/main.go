package models

import (
	"database/sql"
)

type Task struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Regdate   string `json:"regdate"`
}

type TaskCollection struct {
	Task []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM TASKS"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Firstname, &task.Lastname, &task.Email, &task.Regdate)
		if err2 != nil {
			panic(err2)
		}
		result.Task = append(result.Task, task)
	}
	return result
}

func PutTask(db *sql.DB, email string) (int64, error) {
	sql := "INSERT INTO tasks(email) VALUES(?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	// Replace the '?' in our prepared statement with 'firstname'
	result, err2 := stmt.Exec(email)
	if err2 != nil {
		panic(err)
	}
	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
