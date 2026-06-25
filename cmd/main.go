package main

import (
	"fmt"
	"inventoiry-service/api"
	"net/http"
)

func main() {
	productHandler := api.NewProductHttpHandler()

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/products/stock", func(w http.ResponseWriter, r *http.Request) {})

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}
}
