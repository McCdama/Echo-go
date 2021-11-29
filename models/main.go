package models

import "database/sql"

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
