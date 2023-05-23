package main

import (
	"fmt"
	"log"
	"net/http"

	"mosaic-go-interview/src/cache"
	"mosaic-go-interview/src/handlers"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func handleBasicArythmetic(cacheService cache.CacheService) {
	handlersMap := map[string]func(cache.CacheService, http.ResponseWriter, *http.Request){
		"/add":      handlers.AddHandler,
		"/subtract": handlers.SubtractHandler,
		"/multiply": handlers.MultiplyHandler,
		"/divide":   handlers.DivideHandler,
	}

	for path, handler := range handlersMap {
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler(cacheService, w, r)
		})
	}
}

func main() {
	cacheService := cache.NewRedisCacheService()
	handleBasicArythmetic(cacheService)
	fmt.Println("Server listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
