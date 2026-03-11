package main

import (
	database "example/hello/database"
	"example/hello/routes"
	"net/http"
)

type one struct {
	name string
}
type two struct {
	name one
}

func main() {
	database.Connection()
	router := routes.NewRouter()
	http.ListenAndServe(":8000", router)
}
