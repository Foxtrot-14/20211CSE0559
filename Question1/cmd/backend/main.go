package main

import (
	"backend/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Println("Server starting at : 127.0.0.1:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
