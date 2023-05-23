package handlers

import (
	"net/http"
)

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	BasicArythmeticHandler(w, r)
}
