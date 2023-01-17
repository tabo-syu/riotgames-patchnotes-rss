package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	riotgames "github.com/tabo-syu/riotgames-game-articles"
	"golang.org/x/exp/slices"
)

func main() {
	LOLCache := NewCache[riotgames.LOLArticle]()
	ValorantCache := NewCache[riotgames.ValorantArticle]()

	r := gin.Default()
	r.GET("/league-of-legends/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		if !slices.Contains(riotgames.LOLLocales, locale) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid locale",
			})

			return
		}

		articles := LOLCache.Get(locale)
		if articles == nil {
			res, err := riotgames.NewLOLWebsiteArticles(locale)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}

			articles = res.LOLPatchNotes()
			LOLCache.Set(locale, articles)
		}

		c.JSON(http.StatusOK, articles)
	})

	r.GET("/valorant/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		if !slices.Contains(riotgames.ValorantLocales, locale) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid locale",
			})

			return
		}

		articles := ValorantCache.Get(locale)
		if articles == nil {
			res, err := riotgames.NewValorantWebsiteArticles(locale)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
			}

			articles = res.PatchNotes()
			ValorantCache.Set(locale, articles)
		}

		c.JSON(http.StatusOK, articles)
	})

	r.Run()
}
