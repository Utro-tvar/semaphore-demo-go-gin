package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func showIndexPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Home Page",
		},
	)
}

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", showIndexPage)

	router.Run()
}
