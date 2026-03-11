package repos

import (
	database "example/hello/database"
	"example/hello/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data models.UserRegisterDTO) (models.UserDTO, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.UserDTO{}, err
	}
	user := models.UserDTO{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(encPass),
	}
	dbErr := database.DB.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, dbErr
}

func GetUsers() ([]models.UserDTO, error) {
	var users []models.UserDTO
	err := database.DB.Find(&users).Error
	return users, err
}
func GetUser(id int) (models.UserDTO, error) {
	var user models.UserDTO
	err := database.DB.Where("id=?", id).First(&user).Error
	return user, err
}
func GetUserByEmail(email string) (*models.UserDTO, error) {
	var user models.UserDTO
	err := database.DB.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func EditUser(data models.UserDTO) error {
	encPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return database.DB.Model(&data).Updates(models.UserDTO{
		Name:     data.Name,
		Email:    data.Email,
		Password: string(encPass),
	}).Error
}
func DeleteUser(id string) error {
	return database.DB.Delete(&models.UserDTO{}, id).Error
}
