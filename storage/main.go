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

	// test
	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(version)

	defer db.Close()
}
