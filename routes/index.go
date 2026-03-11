package routes

import (
	"example/hello/controllers"
	"example/hello/middlewares"
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	registerHandler := http.HandlerFunc(controllers.RegisterHandler)
	loginHandler := http.HandlerFunc(controllers.LoginHandler)
	home := http.HandlerFunc(controllers.Home)
	router.Handle("/auth/login",
		middlewares.Logger(
			middlewares.LoginValidator(
				loginHandler,
			),
		),
	)
	router.Handle("/auth/register",
		middlewares.Logger(
			middlewares.RegisterValidator(
				registerHandler,
			),
		),
	)
	router.Handle("/home",
		middlewares.Logger(
			middlewares.AuthMiddleware(
				home,
			),
		),
	)
	return router
}
