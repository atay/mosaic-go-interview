package cache

import (
	"time"
)

type CacheService interface {
	Set(key string, value int, expiration time.Duration) error
	Get(key string) (int, bool, error)
}
