package product

import (
	"encoding/json"
	"net/http"
)

// In your Go handler (for example, in GetProducts)
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow all origins (for testing)
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Allow GET method

	products := GetAllProducts() // Call the service to get products
	json.NewEncoder(w).Encode(products)
}
