package handlers

import (
	"fmt"
	"mosaic-go-interview/cache"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDivideHandler(t *testing.T) {
	cacheService := cache.NewMemoryCacheService()

	testCases := []struct {
		x              interface{}
		y              interface{}
		expectedStatus int
		expectedBody   string
	}{
		{10, 5, http.StatusOK, `{"action":"divide","x":10,"y":5,"answer":2,"cached":false}`},  // Positive numbers
		{5, -3, http.StatusOK, `{"action":"divide","x":5,"y":-3,"answer":-1,"cached":false}`}, // Positive and negative numbers
		{0, -2, http.StatusOK, `{"action":"divide","x":0,"y":-2,"answer":0,"cached":false}`},  // Zero and negative number
		{0, -2, http.StatusOK, `{"action":"divide","x":0,"y":-2,"answer":0,"cached":true}`},   // chceck if cached
		{2, 0, http.StatusBadRequest, `{"error":"Cannot divide by zero"}`},                    // Divide by zero error
		{0, 0, http.StatusBadRequest, `{"error":"Cannot divide by zero"}`},                    // Divide by zero error
		{2, "abc", http.StatusBadRequest, `{"error":"Invalid operands"}`},                     // Invalid string operand
		{"abc", 5, http.StatusBadRequest, `{"error":"Invalid operands"}`},                     // Invalid string operand
		{"abc", "def", http.StatusBadRequest, `{"error":"Invalid operands"}`},                 // Invalid string operands
		{"", "", http.StatusBadRequest, `{"error":"Invalid operands"}`},                       // Invalid string operands
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("GET", fmt.Sprintf("/divide?x=%v&y=%v", tc.x, tc.y), nil)
		if err != nil {
			t.Fatal(err)
		}

		res := httptest.NewRecorder()

		DivideHandler(cacheService, res, req)

		if res.Code != tc.expectedStatus {
			t.Errorf("For x=%v, y=%v: Expected status code %d, but got %d", tc.x, tc.y, tc.expectedStatus, res.Code)
		}

		if strings.TrimSpace(res.Body.String()) != tc.expectedBody {
			t.Errorf("For x=%v, y=%v: Expected response body '%s', but got '%s'", tc.x, tc.y, tc.expectedBody, res.Body.String())
		}
	}
}
