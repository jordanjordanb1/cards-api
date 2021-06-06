package decks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Used to init routes
func RouterRegister(router *gin.RouterGroup) {
	router.POST("/", CreateDeck)
	router.GET("/:deckId", FindDeck)
}

func CreateDeck(c *gin.Context) {
	deck := NewDeck()

	deck.Save()

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   deck,
	})
}

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
