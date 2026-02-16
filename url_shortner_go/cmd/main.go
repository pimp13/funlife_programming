package main

import (
	"fmt"
	"log"
	"net/http"
	"url_shortner_go/handler"
	"url_shortner_go/store"

	"github.com/gin-gonic/gin"
)

const PORT = ":4224"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "Welcome to the URL Shortner API",
		})
	})

	r.POST("/create", handler.CreateShortURL)
	r.GET("/:shorturl", handler.HandleShortURLRedirect)

	store.InitStoreService()

	log.Printf("INF: server is running on port %s\n", PORT)
	if err := r.Run(fmt.Sprintf("0.0.0.0%s", PORT)); err != nil {
		log.Fatalf("server error: failed to run server: %s", err.Error())
	}
}
