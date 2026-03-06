package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connection() {
	fmt.Println("database")
	dsn := "root:svipl@tcp(127.0.0.1:3306)/demo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("database error")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database error")
		panic(err)
	}
	fmt.Println("database connected")
	DB = db
}
