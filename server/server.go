package server

import (
	"fmt"
	"net/http"
	auth "project-backend/auth/router"
	order "project-backend/order/router"
	product "project-backend/product/router"
	brand "project-backend/brand/router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RunServer() {
	defer fmt.Println("Server stopped!")

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(router)

	auth.AuthRouter(router)
	product.ProductRouter(router)
	order.OrderRouter(router)
	brand.BrandRouter(router)
	//Orders
	// router.Methods("GET").Path("/orders").HandlerFunc(controller.GetAllOrders)
	// router.Methods("POST").Path("/orders").HandlerFunc(controller.AddOrder)
	// router.Methods("GET").Path("/orders/{id:[0-9]+}").HandlerFunc(controller.GetOrder)
	// router.Methods("PUT").Path("/orders/{id:[0-9]+}").HandlerFunc(controller.UpdateOrder)
	// router.Methods("DELETE").Path("/orders/{id:[0-9]+}").HandlerFunc(controller.DeleteOrder)
	fmt.Println("Server opened at port 8081")

	err := http.ListenAndServe(":8081", handler)
	if err != nil {
		panic(err)
	}
}
