package product

// Simulated database
var products = []Product{
	{ID: 1, Name: "Smartphone", Price: "$399"},
	{ID: 2, Name: "Laptop", Price: "$999"},
	{ID: 3, Name: "Headphones", Price: "$199"},
}

// GetAllProducts retrieves all products
func GetAllProducts() []Product {
	return products
}
