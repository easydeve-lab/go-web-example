package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "assets")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handleIndex)
	router.GET("/ping", handlePing)

	router.Run(":8080")
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func handleIndex(c *gin.Context) {
	// Check if the scheme is available (via X-Forwarded-Proto or TLS)
	var scheme string
	if proto := c.GetHeader("X-Forwarded-Proto"); proto != "" {
		// If X-Forwarded-Proto is present, use it (common with proxies)
		scheme = proto
	} else if c.Request.TLS != nil {
		// If the connection is secure (HTTPS), use https
		scheme = "https"
	} else {
		// Otherwise, default to http
		scheme = "http"
	}

	// Get host (with port if necessary)
	host := c.Request.Host

	path := c.Request.URL.Path

	// Combine to form the base URL
	baseURL := fmt.Sprintf("%s://%s%s", scheme, host, path)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":   "Index - QuickStart Bootstrap Template",
		"baseURL": baseURL,
	})
}
