package cache

import (
	"fmt"

	"github.com/dgraph-io/ristretto"
)

var Cache *ristretto.Cache

// Creates an in-memory cache
func Create() {
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

	Cache = cache
}
