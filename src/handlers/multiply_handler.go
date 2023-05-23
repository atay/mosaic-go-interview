package handlers

import (
	"mosaic-go-interview/src/services"
	"net/http"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r, services.Multiply)
}
