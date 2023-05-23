package handlers

import (
	"mosaic-go-interview/src/cache"
	"net/http"
)

func DivideHandler(cacheService cache.CacheService, w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(cacheService, w, r)
}
