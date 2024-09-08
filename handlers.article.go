package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":    "Home Page",
			"payloads": articles,
		},
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

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
