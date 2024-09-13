// Package handler contains the HTTP handlers for the application.
package handler

import (
	"fmt"
	"github.com/Dubjay18/Ushort/database"
	"github.com/Dubjay18/Ushort/dto"
	shortener "github.com/Dubjay18/Ushort/shortner"
	"github.com/Dubjay18/Ushort/store"
	"github.com/Dubjay18/Ushort/views"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"net/http"
)

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func IndexPageHandler(c *gin.Context) {
	fmt.Printf("IndexPageHandler\n")
	render(c, http.StatusOK, views.Index())
}

// CreateShortUrl is a handler function that creates a short URL from a given long URL.
func CreateShortUrl(c *gin.Context, st *database.Service) {
	// Bind the incoming JSON to a UrlCreationRequest struct.
	var creationRequest *dto.UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		// If JSON binding fails, try binding form data.
		longUrl := c.PostForm("url")
		if longUrl == "" {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "Invalid request",
				Error:   err.Error(),
			})
			return
		}
		// Create a new URL shortener.
		NewShortener := shortener.NewGenerator()

		// Generate a short URL from the long URL.
		shortUrl := NewShortener.GenerateShortLink(longUrl)

		// Save the short URL and the corresponding long URL in the store.
		_, err := st.Db.Collection("urls").InsertOne(nil, map[string]string{"shortUrl": shortUrl, "originalUrl": longUrl})
		if err != nil {
			fmt.Printf("Error saving URL: %v", err)
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{
				Message: "Error saving URL",
				Error:   err.Error(),
			})
			return
		}

		// Determine the protocol to use for the short URL.
		host := c.Request.Host
		protocol := "http"
		if c.Request.TLS != nil {
			protocol = "https"
		}
		// Respond with the short URL.
		c.JSON(200, dto.UrlCreationResponse{
			Message:  "Short URL created successfully",
			ShortUrl: fmt.Sprintf("%s://%s/%s", protocol, host, shortUrl),
		})
		return
	}

	// Create a new URL shortener.
	NewShortener := shortener.NewGenerator()

	// Generate a short URL from the long URL.
	shortUrl := NewShortener.GenerateShortLink(creationRequest.Url)

	// Save the short URL and the corresponding long URL in the store.
	_, err := st.Db.Collection("urls").InsertOne(nil, map[string]string{"shortUrl": shortUrl, "originalUrl": creationRequest.Url})
	if err != nil {
		fmt.Printf("Error saving URL: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Error saving URL",
			Error:   err.Error(),
		})
		return
	}

	// Determine the protocol to use for the short URL.
	host := c.Request.Host
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	// Respond with the short URL.
	c.JSON(200, dto.UrlCreationResponse{
		Message:  "Short URL created successfully",
		ShortUrl: fmt.Sprintf("%s://%s/%s", protocol, host, shortUrl),
	})
}

// HandleShortUrlRedirect is a handler function that redirects a short URL to its corresponding long URL.
func HandleShortUrlRedirect(c *gin.Context) {
	// Get the short URL from the request parameters.
	shortUrl := c.Param("shortUrl")

	// Retrieve the long URL from the store.
	initialUrl, err := store.StoreService.Get(shortUrl)
	if err != nil {
		fmt.Printf("Error retrieving URL: %v", err)
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "URL not found",
			Error:   err.Error(),
		})
		return
	}
	// Redirect to the long URL.
	c.Redirect(302, initialUrl)
}
