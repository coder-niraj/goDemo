package main

import (
	database "example/hello/database"
	"example/hello/routes"
	"net/http"
)

func main() {
	database.Connection()
	router := routes.NewRouter()
	http.ListenAndServe(":8000", router)
}
