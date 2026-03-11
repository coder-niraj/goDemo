package controllers

import (
	"example/hello/helpers"
	"example/hello/middlewares"
	"example/hello/models"
	"example/hello/services"
	"net/http"
)

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	userData, ok := req.Context().Value(middlewares.RegisterUserKey).(models.UserRegisterDTO)
	saveUser, token, err := services.UserRegisterService(userData)
	if err != nil || !ok {
		middlewares.JsonMiddlewareError(res, "error occurred")
		return
	}
	helpers.JSONFormat(res, http.StatusCreated, models.Response{
		Success: true,
		Data:    saveUser,
		Token:   token,
		Message: "user registered successful",
	})
}
func LoginHandler(res http.ResponseWriter, req *http.Request) {
	userData, ok := req.Context().Value(middlewares.LoginUserKey).(models.UserLoginDTO)
	if !ok {
		middlewares.JsonMiddlewareError(res, "invelid request")
		return
	}
	getUserData, token, err := services.UserLoginService(userData)
	if err != nil {
		middlewares.JsonMiddlewareError(res, "error occurred")
		return
	}
	helpers.JSONFormat(res, http.StatusOK, models.Response{
		Success: true,
		Message: "user login successful",
		Data:    getUserData,
		Token:   token,
	})
}
func Home(res http.ResponseWriter, req *http.Request) {
	helpers.JSONFormat(res, http.StatusOK, models.Response{
		Success: true,
		Message: "user login successful",
		Data:    nil,
		Token:   "",
	})
}
