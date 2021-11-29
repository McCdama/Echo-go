package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("Open Conn...")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/vue-app")

	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		panic("db nil")
	}

	// test
	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(version)

	defer db.Close()
	migrate(db)
}

func migrate(db *sql.DB) {
	sql :=
		`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT(6) UNSIGNED NOT NULL AUTO_INCREMENT,
			firstname VARCHAR(30) NOT NULL COLLATE utf8mb4_general_ci,
			lastname VARCHAR(30) NOT NULL COLLATE utf8mb4_general_ci,
			email VARCHAR(50) NULL DEFAULT NULL COLLATE utf8mb4_general_ci,
			reg_date TIMESTAMP NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
			PRIMARY KEY (id) USING BTREE
		)
		COLLATE='utf8mb4_general_ci'
		ENGINE=InnoDB
		;
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}
