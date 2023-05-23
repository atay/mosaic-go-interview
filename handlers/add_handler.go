package handlers

import (
	"mosaic-go-interview/services"
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r, services.Add)
}
