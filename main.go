package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	riotgames "github.com/tabo-syu/riotgames-game-articles"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		articles, err := riotgames.NewLOLWebsiteArticles(riotgames.JaJp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": articles.LOLPatchNotes(),
		})
	})

	r.Run()
}
