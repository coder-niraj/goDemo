package models

import (
	database "example/hello/DataBase"
	database "goHello/DataBase"
)
type user struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:email`
}

func getUsers() ([]user,error){
	rows,err  = database.DB.Query("select * from users")
}