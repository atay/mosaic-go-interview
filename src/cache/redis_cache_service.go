package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCacheService struct {
	client *redis.Client
}

func NewRedisCacheService() *RedisCacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // Set password if required
		DB:       0,  // Use default database
	})

	return &RedisCacheService{
		client: client,
	}
}

func (c *RedisCacheService) Set(key string, value int, expiration time.Duration) error {
	ctx := context.Background()

	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set cache value: %w", err)
	}

	return nil
}

func (c *RedisCacheService) Get(key string) (int, bool, error) {
	ctx := context.Background()

	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, false, nil // Cache miss
	} else if err != nil {
		return 0, false, fmt.Errorf("failed to get cache value: %w", err)
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		return 0, false, fmt.Errorf("failed to convert cache value to int: %w", err)
	}

	return value, true, nil // Cache hit
}
