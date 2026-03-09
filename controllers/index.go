package controllers

import (
	"encoding/json"
	repos "example/hello/Repos"
	"example/hello/helpers"
	"example/hello/middlewares"
	"example/hello/models"
	"fmt"
	"net/http"
)

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	userData := req.Context().Value(middlewares.RegisterUserKey).(models.UserRegister)
	saveUser, err := repos.CreateUser(userData)
	if err != nil {
		middlewares.JsonMiddlewareError(res, "error occurred")
		fmt.Println(err)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	jwt, err := helpers.CreateJWT(saveUser)

	json.NewEncoder(res).Encode(models.Response{Success: true, Data: saveUser, Token: jwt, Message: "user registered successful"})
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	resp := models.Response{Success: true, Message: "delivered"}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
func getAllUsers(res http.ResponseWriter, req *http.Request) {
	resp := models.Response{Success: true, Message: "delivered"}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
func getOneUser(res http.ResponseWriter, req *http.Request) {
	resp := models.Response{Success: true, Message: "delivered"}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
