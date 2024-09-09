package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

func showArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	article, err := getArticleByID(id)
	if err != nil {
		// If the article is not found, abort with an error
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	render(
		c,
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
		"article.html",
	)
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
