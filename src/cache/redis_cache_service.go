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
	ttl    time.Duration
}

func NewRedisCacheService(redisUrl string) *RedisCacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // Set password if required
		DB:       0,  // Use default database
	})

	return &RedisCacheService{
		client: client,
		ttl:    60 * time.Second,
	}
}

func (c *RedisCacheService) Set(key string, value int) error {
	ctx := context.Background()

	err := c.client.Set(ctx, key, value, c.ttl).Err()
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

	_, err = c.client.Expire(ctx, key, c.ttl).Result()
	if err != nil {
		return 0, false, fmt.Errorf("failed to refresh cache TTL: %w", err)
	}

	return value, true, nil // Cache hit
}
