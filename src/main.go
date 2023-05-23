package main

import (
	"fmt"
	"log"
	"net/http"

	"mosaic-go-interview/src/handlers"
)

func handleBasicArythmetic() {
	http.HandleFunc("/add", handlers.AddHandler)
	http.HandleFunc("/subtract", handlers.SubtractHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)
	http.HandleFunc("/divide", handlers.DivideHandler)

}

func main() {
	handleBasicArythmetic()
	fmt.Println("Server listening on http://localhost/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
