package main

import (
	"log"
	"net/http"
	"project_marketplace/services/market_service/internal/middleware"
	"project_marketplace/services/market_service/internal/product"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Apply the CORS middleware
	r.Use(middleware.CORS)

	// Define your routes
	r.HandleFunc("/api/products", product.GetProducts).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
