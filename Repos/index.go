package repos

import (
	database "example/hello/database"
	"example/hello/models"
)

func CreateUser(data models.UserRegister) (models.User, error) {
	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}
	err := database.DB.Create(&user).Error
	return user, err
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}
func GetUser(id int) (models.User, error) {
	var user models.User
	err := database.DB.Where("id=?", id).First(&user).Error
	return user, err
}
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email=?", email).First(&user).Error
	return &user, err
}
func EditUser(data models.User) error {
	return database.DB.Model(&data).Updates(models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}).Error
}
func DeleteUser(id string) error {
	return database.DB.Delete(&models.User{}, id).Error
}
