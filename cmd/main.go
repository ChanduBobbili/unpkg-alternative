package main

import (
	"fmt"
	"log"
	"net/http"
	router "unpkg-alternative/http/router"
)

func main() {
	// Initialize HTTP router
	router := router.NewRouter()
	fmt.Println("Server running at http://localhost:8080")
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
