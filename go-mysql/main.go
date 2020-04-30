package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go Mysql")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('Ivan')")

	if err != nil {
		panic(err.Error())
	}

	defer isert.Close()

	fmt.Println("Successfully inserted into users table")
}
