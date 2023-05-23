package handlers

import (
	"mosaic-go-interview/src/cache"
	"net/http"
)

func SubtractHandler(cacheService cache.CacheService, w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(cacheService, w, r)
}
