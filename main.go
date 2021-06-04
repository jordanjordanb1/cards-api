package main

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
	"github.com/gin-gonic/gin"

	"github.com/jordanjordanb1/cards-api/api/decks"
)

var Cache *ristretto.Cache

func main() {
	createCache()

	// Initialize router
	router := gin.Default()

	// API route
	v1 := router.Group("/api")

	// Decks route
	decks.RouterRegister(v1.Group("/decks"))

	router.Run() // listen and serve on 0.0.0.0:8080
}

// Creates an in-memory cache
func createCache() {
	fmt.Printf("Creating ristretto cache...\n")

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Ristretto cache initialized\n")

	// Sets global
	Cache = cache
}
