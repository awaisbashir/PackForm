package router

import (
	"go-post/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/orders/", middleware.GetOrders).Methods("GET", "OPTIONS")

	return router
}
