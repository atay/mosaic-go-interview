package handlers

import (
	"net/http"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r)
}
