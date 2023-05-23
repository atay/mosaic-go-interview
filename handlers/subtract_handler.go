package handlers

import (
	"mosaic-go-interview/services"
	"net/http"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r, services.Subtract)
}
