package main

import (
	"fmt"
	"github.com/Dubjay18/Ushort/handler"
	"github.com/Dubjay18/Ushort/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getVars() {
	// Get the environment variables

	err := godotenv.Load()

	if err != nil {
		println("Error loading .env file")
	}

}

func main() {
	getVars()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the U-Short API",
		})
	})

	r.POST("/create", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.StoreService.Init()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
