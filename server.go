package main

import (
	"github.com/gin-gonic/gin"
	"fmt"

	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Homepage",
		})
	})

	r.POST("/", func(c *gin.Context) {
		fmt.Println("POST RECEIVED")
	})
	r.Run(":8080")
}