package services

import "testing"

func TestSubtract(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{5, 2, 3},    // positive numbers
		{0, 5, -5},   // subtract from zero
		{-4, -2, -2}, // negative numbers
		{-2, 3, -5},  // subtracting a positive from a negative
		{10, -5, 15}, // subtracting a negative from a positive
		{0, 0, 0},    // both operands are zero
	}

	for _, tc := range testCases {
		result, err := Subtract(tc.x, tc.y)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d - %d, expected result %d, but got %d", tc.x, tc.y, tc.expected, result)
		}
	}
}
