package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ErrorResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
