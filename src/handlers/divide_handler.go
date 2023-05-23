package handlers

import (
	"mosaic-go-interview/src/services"
	"net/http"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r, services.Divide)
}
