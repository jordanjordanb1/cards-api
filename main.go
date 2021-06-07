package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jordanjordanb1/cards-api/cache"
	"github.com/jordanjordanb1/cards-api/decks"
)

func main() {
	fmt.Print("Initializing server...\n")

	// Creates in-memory cache
	cache.Create()

	// Initialize router
	router := gin.Default()

	// API route
	v1 := router.Group("/api")

	// Decks route
	decks.RouterRegister(v1.Group("/decks"))

	error := router.Run() // listen and serve on 0.0.0.0:8080

	if error != nil {
		fmt.Printf("An error occured when Initializing server. Error: %v\n", error)
		return
	}
}
