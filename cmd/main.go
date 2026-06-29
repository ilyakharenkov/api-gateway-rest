package main

import (
	"fmt"
	"inventoiry-service/api"
	"inventoiry-service/internal/repository"
	"inventoiry-service/internal/service"
	"net/http"
)

func main() {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := api.NewProductHttpHandler(productService)

	mux := http.NewServeMux()
	productHandler.RegisterRoutes(mux)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err)
	}
}
