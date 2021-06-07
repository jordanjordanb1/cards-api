package decks

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Used to init routes
func RouterRegister(router *gin.RouterGroup) {
	router.POST("/", CreateDeck)
	router.GET("/:deckId", FindDeck)
}

// Handles POST route to create a new deck
func CreateDeck(c *gin.Context) {
	var deck Deck

	cardSelection := c.Query("cards")
	shuffleQuery := c.Query("shuffle")

	shuffle, _ := strconv.ParseBool(shuffleQuery)

	// Checks if query was set
	if cardSelection != "" {
		// Splits query string into slice
		cardSelectionArray := strings.Split(cardSelection, ",")

		deck = NewCustomDeck(shuffle, cardSelectionArray)
	} else {
		// No query, create normal deck
		deck = NewDeck(shuffle)
	}

	// Saves deck to cache
	deck.Save()

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   deck,
	})
}

// Handles GET route to fetch deck by it's id
func FindDeck(c *gin.Context) {
	deckId := c.Param("deckId")

	deck, error := GetDeckById(deckId)

	// If deck not found, handle accordingly
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": error.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   deck,
	})
}
