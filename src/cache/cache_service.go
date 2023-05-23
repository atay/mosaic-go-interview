package cache

type CacheService interface {
	Set(key string, value int) error
	Get(key string) (int, bool, error)
}
