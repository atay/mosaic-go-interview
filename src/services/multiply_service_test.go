package services

import "testing"

func TestMultiply(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 5, 10},   // positive numbers
		{0, 5, 0},    // one operand is zero
		{-3, -4, 12}, // negative numbers
		{7, -2, -14}, // mixed signs
	}

	for _, tc := range testCases {
		result, err := Multiply(tc.x, tc.y)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d * %d, expected result %d, but got %d", tc.x, tc.y, tc.expected, result)
		}
	}
}
