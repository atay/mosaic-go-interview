package handlers

import (
	"net/http"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r)
}
