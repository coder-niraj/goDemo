package services

import (
	repos "example/hello/Repos"
	"example/hello/helpers"
	"example/hello/models"
	"fmt"

	"gorm.io/gorm"
)

func UserRegisterService(userData models.UserRegisterDTO) (*models.UserDTO, string, error) {
	saveUser, err := repos.CreateUser(userData)
	if err != nil {
		return nil, "", err
	}
	jwt, jwtErr := helpers.CreateJWT(saveUser)
	if jwtErr != nil {
		return nil, "", err
	}
	return &saveUser, jwt, nil
}
func UserLoginService(userData models.UserLoginDTO) (*models.UserDTO, string, error) {
	user, err := repos.GetUserByEmail(userData.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "", err
		}
		return nil, "", err
	}
	passCompare := helpers.ComparePassword(user.Password, userData.Password)
	if passCompare == false {
		return nil, "", fmt.Errorf("invalid password")
	}
	jwt, err := helpers.CreateJWT(*user)
	if err != nil {
		return nil, "", err
	}
	return user, jwt, nil
}
