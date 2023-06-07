package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mosaic-go-interview/cache"
	"mosaic-go-interview/handlers"
)

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
	redisUrl := os.Getenv("REDIS_URL")
	cacheService := cache.NewRedisCacheService(redisUrl)
	handleBasicArythmetic(cacheService)
	port := os.Getenv("HTTP_PORT")

	log.Printf("Server listening on http://localhost:%s/\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
