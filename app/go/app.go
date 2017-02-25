package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	router := initRouter()

	// Group of API calls
	api := router.Group("/api")
	{
		api.GET("/healthcheck", healthcheck)
	}

	web := router.Group("/")
	{
		web.GET("/", index)
	}

	// By default it serves on :8080
	// PORT environment variable was defined.
	router.Run()
}

func initRouter() (router *gin.Engine) {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router = gin.Default()

	// Set html render options
	router.LoadHTMLGlob("templates/*")

	return
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func index(c *gin.Context) {
	// TODO actual posts
	var posts [0]int

	c.HTML(http.StatusOK, "index.html", gin.H{"posts": posts})

}