package main

import (
	"encoding/json"
	database "example/hello/DataBase"
	"fmt"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("here api called")
	res.Header().Set("Content-Type", "application/json")
	database.Connection()
	name := User{
		Name: "niraj",
		Age:  10,
	}
	json.NewEncoder(res).Encode(name)
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8000")
	http.ListenAndServe(":8000", nil)
}
