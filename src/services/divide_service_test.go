package services

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		x           int
		y           int
		expected    int
		expectedErr error
	}{
		{10, 2, 5, nil}, // positive numbers
		{0, 5, 0, nil},  // numerator is zero
		{10, 0, 0, errors.New("Cannot divide by zero")}, // denominator is zero
		{-15, -3, 5, nil}, // negative numbers
		{7, 2, 3, nil},    // non-divisible numbers
		{12, 4, 3, nil},   // evenly divisible numbers
	}

	for _, tc := range testCases {
		result, err := Divide(tc.x, tc.y)

		if tc.expectedErr != nil && err != nil && tc.expectedErr.Error() != err.Error() {
			t.Errorf("For %d / %d, expected error '%v', but got '%v'", tc.x, tc.y, tc.expectedErr, err)
		}

		if err == nil && result != tc.expected {
			t.Errorf("For %d / %d, expected result %d, but got %d", tc.x, tc.y, tc.expected, result)
		}
	}
}
