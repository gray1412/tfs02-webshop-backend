package router

import (
	"project-backend/brand/controller"

	"github.com/gorilla/mux"
)

func BrandRouter(r *mux.Router) {
	r = r.PathPrefix("/brands").Subrouter()
	r.Methods("GET").Path("").HandlerFunc(controller.GetAllBrand)
}
