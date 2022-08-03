package router

import (
	 "project-backend/product/controller"

	"github.com/gorilla/mux"
)

func ProductRouter(r *mux.Router) {
	r = r.PathPrefix("/products").Subrouter()

	r.Methods("GET").Path("").HandlerFunc(controller.SearchProduct)
	r.Methods("GET").Path("/{alias}").HandlerFunc(controller.GetProductByAlias)
	r.Methods("POST").Path("").HandlerFunc(controller.CreateProduct)
	// r.Methods("PUT").Path("/{id:[0-9]+}").HandlerFunc(controller.UpdateOne)
	// r.Methods("DELETE").Path("/{id:[0-9]+}").HandlerFunc(controller.DeleteOne)

}
