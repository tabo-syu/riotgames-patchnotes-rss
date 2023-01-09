package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	riotgames "github.com/tabo-syu/riotgames-game-articles"
)

func main() {
	r := gin.Default()
	r.GET("/league-of-legends/:locale", func(c *gin.Context) {
		articles, err := riotgames.NewLOLWebsiteArticles(c.Param("locale"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, articles.LOLPatchNotes())
	})

	r.GET("/valorant/:locale", func(c *gin.Context) {
		articles, err := riotgames.NewValorantWebsiteArticles(c.Param("locale"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, articles.PatchNotes())
	})

	r.Run()
}
