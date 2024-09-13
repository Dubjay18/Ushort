// Package store provides the storage layer for the application.
package store

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

// Storage interface defines the methods that our storage service must implement.
type Storage interface {
	Init() *StorageService
	Save(shortUrl string, originalUrl string) error
	Get(shortUrl string) (string, error)
}

// StorageService struct is a wrapper around the raw Redis client.
type StorageService struct {
	mongoClient *mongo.Client
}

// Top level declarations for the StoreService
var (
	StoreService = &StorageService{}    // Singleton instance of StorageService.
	ctx          = context.Background() // Context for Redis operations.
)

// CacheDuration is the duration for which the data will be cached in Redis.
const CacheDuration = 6 * time.Hour

// Init method initializes the storage service and returns a pointer to the storage service.
func (s StorageService) Init() *StorageService {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. ")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	StoreService.mongoClient = client
	return StoreService
}

// Save method saves the short URL and original URL in the Redis cache.
func (s StorageService) Save(shortUrl string, originalUrl string) error {
	_, err := s.mongoClient.Database("shortener").Collection("urls").InsertOne(ctx, map[string]string{"shortUrl": shortUrl, "originalUrl": originalUrl})

	if err != nil {
		fmt.Printf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)

		return err
	}
	return nil
}

// Get method retrieves the original URL from the Redis cache using the short URL.
func (s StorageService) Get(shortUrl string) (string, error) {
	var result map[string]string
	err := s.mongoClient.Database("shortener").Collection("urls").FindOne(ctx, map[string]string{"shortUrl": shortUrl}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result["originalUrl"], nil
}
