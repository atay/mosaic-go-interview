package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"mosaic-go-interview/src/response"
)

func getParams(params url.Values) (int, int, error) {
	x, err := strconv.Atoi(params.Get("x"))
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(params.Get("y"))
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func sendResponse(w http.ResponseWriter, response response.ArthmeticResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func BasicArythmeticHandler(w http.ResponseWriter, r *http.Request, service func(int, int) (int, error)) {
	x, y, err := getParams(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	action := r.URL.Path[1:]

	result, err := service(x, y)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := response.ArthmeticResponse{
		Action: action,
		X:      x,
		Y:      y,
		Answer: result,
		Cached: false,
	}

	sendResponse(w, response)
}
