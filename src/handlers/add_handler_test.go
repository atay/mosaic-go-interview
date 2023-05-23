package handlers

import (
	"fmt"
	"mosaic-go-interview/src/cache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddHandler(t *testing.T) {

	cacheService := cache.NewMemoryCacheService()

	testCases := []struct {
		x              interface{}
		y              interface{}
		expectedStatus int
		expectedBody   string
	}{
		{2, 5, http.StatusOK, `{"action":"add","x":2,"y":5,"answer":7,"cached":false}`},    // Positive numbers
		{-3, 8, http.StatusOK, `{"action":"add","x":-3,"y":8,"answer":5,"cached":false}`},  // Negative and positive numbers
		{0, -2, http.StatusOK, `{"action":"add","x":0,"y":-2,"answer":-2,"cached":false}`}, // Zero and negative number
		{2, 0, http.StatusOK, `{"action":"add","x":2,"y":0,"answer":2,"cached":false}`},    // Positive number and zero
		{0, 0, http.StatusOK, `{"action":"add","x":0,"y":0,"answer":0,"cached":false}`},    // Both operands are zero
		{2, -5, http.StatusOK, `{"action":"add","x":2,"y":-5,"answer":-3,"cached":false}`}, // Positive number and negative number
		{2, -5, http.StatusOK, `{"action":"add","x":2,"y":-5,"answer":-3,"cached":true}`},  // check if cached
		{2, "abc", http.StatusBadRequest, `{"error":"Invalid operands"}`},                  // Invalid string operand
		{"abc", 5, http.StatusBadRequest, `{"error":"Invalid operands"}`},                  // Invalid string operand
		{"abc", "def", http.StatusBadRequest, `{"error":"Invalid operands"}`},              // Invalid string operands
		{"", "", http.StatusBadRequest, `{"error":"Invalid operands"}`},                    // Invalid string operands
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", fmt.Sprintf("/add?x=%v&y=%v", tc.x, tc.y), nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		AddHandler(cacheService, res, req)

		if res.Code != tc.expectedStatus {
			t.Errorf("For x=%v, y=%v: Expected status code %d, but got %d", tc.x, tc.y, tc.expectedStatus, res.Code)
		}

		if strings.TrimSpace(res.Body.String()) != tc.expectedBody {
			t.Errorf("For x=%v, y=%v: Expected response body '%s', but got '%s'", tc.x, tc.y, tc.expectedBody, res.Body.String())
		}
	}
}
