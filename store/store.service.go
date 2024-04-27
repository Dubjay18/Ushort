// Package store provides the storage layer for the application.
package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
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
	redisClient *redis.Client
}

// Top level declarations for the StoreService and Redis context.
var (
	StoreService = &StorageService{}    // Singleton instance of StorageService.
	ctx          = context.Background() // Context for Redis operations.
)

// CacheDuration is the duration for which the data will be cached in Redis.
const CacheDuration = 6 * time.Hour

// Init method initializes the storage service and returns a pointer to the storage service.
func (s StorageService) Init() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis-18402.c10.us-east-1-2.ec2.redns.redis-cloud.com:18402",
		Password: "AOJ3POcDntTGXwofilBM1uUTdvuYAyZV",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	fmt.Printf("\nRedis client: %v\n", redisClient)
	StoreService.redisClient = redisClient
	return StoreService
}

// Save method saves the short URL and original URL in the Redis cache.
func (s StorageService) Save(shortUrl string, originalUrl string) error {
	err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
		return err
	}
	return nil
}

// Get method retrieves the original URL from the Redis cache using the short URL.
func (s StorageService) Get(shortUrl string) (string, error) {
	val, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s\n", err, shortUrl))
		return "", err
	}
	return val, nil
}
