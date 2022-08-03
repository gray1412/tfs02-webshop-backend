package router

import (
	auth "project-backend/auth/controller"
	"project-backend/middleware"

	"github.com/gorilla/mux"
)

func AuthRouter(r *mux.Router) {
	r = r.PathPrefix("/auth").Subrouter()
	r.Methods("GET").Path("/").HandlerFunc(middleware.TokenAuth(auth.GetUser))
	r.Methods("POST").Path("/login").HandlerFunc(auth.Login)
	r.Methods("POST").Path("/signup").HandlerFunc(auth.SignUp)
}
