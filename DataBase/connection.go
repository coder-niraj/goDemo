package database

import (
	"example/hello/schema"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	fmt.Println("database")
	dsn := "root:svipl@tcp(127.0.0.1:3306)/demo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database error")
	}
	db.AutoMigrate(&schema.User{})
	fmt.Println("database connected")
	DB = db
}
