package routes

import (
	"task/handlers"
	"task/middlewares"

	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	router.HandleFunc("/api/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	userRouter := router.PathPrefix("/api/user").Subrouter()
	userRouter.Use(middlewares.Authenticate)
	userRouter.HandleFunc("/payment-methods", handlers.AddCreditCard).Methods("POST")
	userRouter.HandleFunc("/payment-methods/{id}", handlers.DeleteCreditCard).Methods("DELETE")
	userRouter.HandleFunc("/products", handlers.ListProducts).Methods("GET")
	userRouter.HandleFunc("/purchase", handlers.BuyProducts).Methods("POST")
	userRouter.HandleFunc("/orders", handlers.GetUserOrders).Methods("GET")

	adminRouter := router.PathPrefix("/api/admin").Subrouter()
	adminRouter.Use(middlewares.Authenticate, middlewares.RequireAdmin)
	adminRouter.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	adminRouter.HandleFunc("/products/{id}", handlers.UpdateProduct).Methods("PUT")
	adminRouter.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE")
	adminRouter.HandleFunc("/product-sales", handlers.GetProductSales).Methods("GET")
}
