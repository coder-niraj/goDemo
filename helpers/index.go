package helpers

import (
	repos "example/hello/Repos"
	"example/hello/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var JwtSecret = []byte("YOUR_SECRET_KEY")

func CreateJWT(data models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"email":      data.Email,
		"id":         data.Id,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsUserExist(email string) (bool, error) {
	userData, err := repos.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if userData != nil {
		return true, nil
	}
	return false, nil
}
