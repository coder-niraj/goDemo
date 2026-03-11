package models

type UserDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRegisterDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ErrorResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
