package handler

import (
	"fmt"
	"github.com/Dubjay18/Ushort/dto"
	"github.com/Dubjay18/Ushort/shortner"
	"github.com/Dubjay18/Ushort/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateShortUrl(c *gin.Context) {
	var creationRequest dto.UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewShortener := shortener.NewGenerator()

	shortUrl := NewShortener.GenerateShortLink(creationRequest.LongUrl)

	err := store.StoreService.Save(shortUrl, creationRequest.LongUrl)
	if err != nil {
		fmt.Printf("Error saving URL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	host := c.Request.Host
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"

	}
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": fmt.Sprintf("%s://%s/%s", protocol, host, shortUrl),
	})

}
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl, err := store.StoreService.Get(shortUrl)
	if err != nil {
		fmt.Printf("Error retrieving URL: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(302, initialUrl)
}
