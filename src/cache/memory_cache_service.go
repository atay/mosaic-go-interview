package cache

import "time"

type MemoryCacheService struct {
	cache map[string]int
}

func NewMemoryCacheService() *MemoryCacheService {
	return &MemoryCacheService{
		cache: make(map[string]int),
	}
}

func (m *MemoryCacheService) Set(key string, value int, expiration time.Duration) error {
	m.cache[key] = value
	return nil
}

func (m *MemoryCacheService) Get(key string) (int, bool, error) {
	value, ok := m.cache[key]
	if !ok {
		return 0, false, nil // Cache miss
	}
	return value, true, nil // Cache hit
}
