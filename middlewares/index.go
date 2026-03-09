package middlewares

import (
	"context"
	"encoding/json"
	repos "example/hello/Repos"
	"example/hello/helpers"
	"example/hello/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const RegisterUserKey contextKey = "registerUser"
const LoginUserKey contextKey = "loginUser"
const UserKey contextKey = "user"

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Print(req.Method, req.URL.Path)
		next.ServeHTTP(res, req)
	})
}
func JsonMiddlewareError(res http.ResponseWriter, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(models.ErrorResp{Success: false, Message: message})
}

func RegisterValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var rules models.UserRegister
		err := json.NewDecoder(req.Body).Decode(&rules)
		if err != nil {
			JsonMiddlewareError(res, "error occured")
			return
		} else if rules.Name == "" {
			JsonMiddlewareError(res, "name is required")
			return
		} else if rules.Email == "" {
			JsonMiddlewareError(res, "email is required")
			return
		} else if rules.Password == "" {
			JsonMiddlewareError(res, "password is required")
			return
		} else {
			exist, err := helpers.IsUserExist(rules.Email)
			if err != nil {
				JsonMiddlewareError(res, "error occurred")
				return
			}
			if exist {
				JsonMiddlewareError(res, "Already Registered")
				return

			}
			ctx := context.WithValue(req.Context(), RegisterUserKey, rules)
			next.ServeHTTP(res, req.WithContext(ctx))
		}
	})
}
func LoginValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var rules models.UserLogin
		err := json.NewDecoder(req.Body).Decode(&rules)
		if err != nil {
			JsonMiddlewareError(res, "error occured")
			return
		} else if rules.Email == "" {
			JsonMiddlewareError(res, "email is required")
			return
		} else if rules.Password == "" {
			JsonMiddlewareError(res, "password is required")
			return
		} else {
			ctx := context.WithValue(req.Context(), LoginUserKey, rules)
			next.ServeHTTP(res, req.WithContext(ctx))
		}
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			JsonMiddlewareError(res, "token is required")
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "bearer" {
			JsonMiddlewareError(res, "token is invalid")
			return
		}
		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return helpers.JwtSecret, nil
		})
		if err != nil || !token.Valid {
			JsonMiddlewareError(res, "unauthorized")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			JsonMiddlewareError(res, "invalid token")
			return
		}
		userID := int(claims["user_id"].(float64))
		dbUser, err := repos.GetUser(userID)
		if err != nil {
			JsonMiddlewareError(res, "invalid token")
			return
		}
		ctx := context.WithValue(req.Context(), UserKey, dbUser)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
