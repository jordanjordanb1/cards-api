package decks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Used to init routes
func RouterRegister(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		deck := NewDeck()

		c.JSON(http.StatusCreated, gin.H{"deck": deck})
	})
}
