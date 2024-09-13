package main

import (
	"fmt"
	"github.com/Dubjay18/Ushort/database"
	"github.com/Dubjay18/Ushort/handler"
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
	// Note that store initialization happens here
	dbInstance := database.New("shorturl", "localhost", "27017")

	r := gin.Default()
	r.Use(CORS())
	r.GET("/", func(c *gin.Context) {
		handler.IndexPageHandler(c)
	})

	r.POST("/create", func(c *gin.Context) {
		handler.CreateShortUrl(c, dbInstance)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
