package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = ":4224"

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// store.SaveUrlMapping("https://apophis.ir", "https://ap.ir", "h4312fs")
	// log.Println(store.GetUrlMapping("https://apophis.ir"))

	log.Printf("INF: server is running on port %s\n", PORT)
	if err := r.Run(fmt.Sprintf("0.0.0.0%s", PORT)); err != nil {
		log.Fatalf("server error: failed to run server: %s", err.Error())
	}
}
