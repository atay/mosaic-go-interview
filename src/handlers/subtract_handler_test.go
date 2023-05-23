package handlers

import (
	"fmt"
	"mosaic-go-interview/src/cache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSubtractHandler(t *testing.T) {

	cacheService := cache.NewMemoryCacheService()

	testCases := []struct {
		x              interface{}
		y              interface{}
		expectedStatus int
		expectedBody   string
	}{
		{10, 5, http.StatusOK, `{"action":"subtract","x":10,"y":5,"answer":5,"cached":false}`},  // Positive numbers
		{5, -3, http.StatusOK, `{"action":"subtract","x":5,"y":-3,"answer":8,"cached":false}`},  // Positive and negative numbers
		{0, -2, http.StatusOK, `{"action":"subtract","x":0,"y":-2,"answer":2,"cached":false}`},  // Zero and negative number
		{-5, 3, http.StatusOK, `{"action":"subtract","x":-5,"y":3,"answer":-8,"cached":false}`}, // Negative and positive numbers
		{-5, 3, http.StatusOK, `{"action":"subtract","x":-5,"y":3,"answer":-8,"cached":true}`},  // check if cached
		{0, 0, http.StatusOK, `{"action":"subtract","x":0,"y":0,"answer":0,"cached":false}`},    // Both operands are zero
		{"abc", 5, http.StatusBadRequest, `{"error":"Invalid operands"}`},                       // Invalid string operand
		{2, "abc", http.StatusBadRequest, `{"error":"Invalid operands"}`},                       // Invalid string operand
		{"abc", "def", http.StatusBadRequest, `{"error":"Invalid operands"}`},                   // Invalid string operands
		{"", "", http.StatusBadRequest, `{"error":"Invalid operands"}`},                         // Invalid string operands
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", fmt.Sprintf("/subtract?x=%v&y=%v", tc.x, tc.y), nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		SubtractHandler(cacheService, res, req)

		if res.Code != tc.expectedStatus {
			t.Errorf("For x=%v, y=%v: Expected status code %d, but got %d", tc.x, tc.y, tc.expectedStatus, res.Code)
		}

		if strings.TrimSpace(res.Body.String()) != tc.expectedBody {
			t.Errorf("For x=%v, y=%v: Expected response body '%s', but got '%s'", tc.x, tc.y, tc.expectedBody, res.Body.String())
		}
	}
}
