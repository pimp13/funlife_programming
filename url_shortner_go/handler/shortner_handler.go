package handler

import (
	"fmt"
	"net/http"
	"url_shortner_go/shortner"
	"url_shortner_go/store"

	"github.com/gin-gonic/gin"
)

type URLCreationRequestBody struct {
	URL    string `json:"url" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

const BASE_URL = "http://127.0.0.1:4224"

func CreateShortURL(c *gin.Context) {
	var crationRequest URLCreationRequestBody
	if err := c.ShouldBindJSON(&crationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":     false,
			"errors": err.Error(),
		})
		return
	}

	shortUrl := shortner.GenerateShortURL(crationRequest.URL, crationRequest.UserId)
	if err := store.SaveUrlMapping(shortUrl, crationRequest.URL, crationRequest.UserId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"ok":      true,
		"message": "url shortner is ok",
		"data": map[string]string{
			"short_url": fmt.Sprintf("%s/%s", BASE_URL, shortUrl),
		},
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("shorturl")
	initURL := store.RetrieveInitialUrl(shortURL)
	c.Redirect(http.StatusPermanentRedirect, initURL)
}
