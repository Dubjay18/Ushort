package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"

	"time"
)

type Storage interface {
	Init() *StorageService
	Save(shortUrl string, originalUrl string, userId string) error
	Get(shortUrl string) (string, error)
}

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the StoreService and Redis context
var (
	StoreService = &StorageService{}
	ctx          = context.Background()
)

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
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

// Save the short URL and original URL in the Redis cache
func (s StorageService) Save(shortUrl string, originalUrl string, userId string) error {
	err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
		return err
	}
	return nil
}

// Retrieve the original URL from the Redis cache
func (s StorageService) Get(shortUrl string) (string, error) {
	val, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving key url | Error: %v - shortUrl: %s\n", err, shortUrl))
		return "", err
	}
	return val, nil
}

func GetStore() Storage {
	storeInit := StorageService{}
	storeInit.Init()
	return storeInit
}
