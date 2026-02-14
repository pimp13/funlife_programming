package store

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const CacheDuration = 24 * time.Hour

type StoreService struct {
	redisClient *redis.Client
}

var (
	ctx = context.Background()
	// storeService = &StoreService{}
	storeService = StoreService{}
)

func initStoreService() *StoreService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("redis error: failed to ping redis: %s\n", err.Error())
		return nil
	}
	if pong == "PONG" {
		log.Println("INF: redis started successfully.")
	}
	storeService.redisClient = redisClient
	return &storeService
}

func SaveUrlMapping(shortUrl, longUrl, userId string) {
	s := initStoreService()
	key := shortUrl
	value := longUrl
	if err := s.redisClient.Set(ctx, key, value, CacheDuration).Err(); err != nil {
		log.Printf("ERR: Failed to save URL mapping: %v\n", err)
	}
}

func GetUrlMapping(shortUrl string) string {
	s := initStoreService()
	result := s.redisClient.Get(ctx, shortUrl)
	if result.Err() != nil {
		log.Printf("ERR: Failed to get URL mapping: %v\n", result.Err())
	}
	return result.Val()
}
