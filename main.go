package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jordanjordanb1/cards-api/cache"
	"github.com/jordanjordanb1/cards-api/decks"
)

func main() {
	// Creates in-memory cache
	cache.Create()

	// Initialize router
	router := gin.Default()

	// API route
	v1 := router.Group("/api")

	// Decks route
	decks.RouterRegister(v1.Group("/decks"))

	router.Run() // listen and serve on 0.0.0.0:8080
}
