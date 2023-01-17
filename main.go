package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	riotgames "github.com/tabo-syu/riotgames-game-articles"
	"golang.org/x/exp/slices"
)

func main() {
	r := gin.Default()
	r.GET("/league-of-legends/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		if !slices.Contains(riotgames.LOLLocales, locale) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid locale",
			})
		}

		articles, err := riotgames.NewLOLWebsiteArticles(locale)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, articles.LOLPatchNotes())
	})

	r.GET("/valorant/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		if !slices.Contains(riotgames.ValorantLocales, locale) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid locale",
			})
		}

		articles, err := riotgames.NewValorantWebsiteArticles(locale)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, articles.PatchNotes())
	})

	r.Run()
}
