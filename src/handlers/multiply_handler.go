package handlers

import (
	"mosaic-go-interview/cache"
	"net/http"
)

func MultiplyHandler(cacheService cache.CacheService, w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(cacheService, w, r)
}
