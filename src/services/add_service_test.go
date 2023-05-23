package services

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 5, 7},    // positive numbers
		{-3, -4, -7}, // negative numbers
		{-3, 0, -3},  // zero number
	}

	for _, tc := range testCases {
		result, err := Add(tc.x, tc.y)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d + %d, expected result %d, but got %d", tc.x, tc.y, tc.expected, result)
		}
	}
}
