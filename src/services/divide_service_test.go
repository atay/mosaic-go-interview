package services

import (
	"errors"
	"mosaic-go-interview/commands"
	"testing"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		command     commands.BasicArythemticCommand
		expected    int
		expectedErr error
	}{
		{commands.BasicArythemticCommand{Action: "divide", X: 10, Y: 2}, 5, nil},                                 // positive numbers
		{commands.BasicArythemticCommand{Action: "divide", X: 0, Y: 5}, 0, nil},                                  // numerator is zero
		{commands.BasicArythemticCommand{Action: "divide", X: 10, Y: 0}, 0, errors.New("Cannot divide by zero")}, // denominator is zero
		{commands.BasicArythemticCommand{Action: "divide", X: -15, Y: -3}, 5, nil},                               // negative numbers
		{commands.BasicArythemticCommand{Action: "divide", X: 7, Y: 2}, 3, nil},                                  // non-divisible numbers
		{commands.BasicArythemticCommand{Action: "divide", X: 12, Y: 4}, 3, nil},                                 // evenly divisible numbers
	}

	for _, tc := range testCases {
		result, err := Divide(tc.command)

		if tc.expectedErr != nil {
			if err == nil || tc.expectedErr.Error() != err.Error() {
				t.Errorf("For %d / %d, expected error '%v', but got '%v'", tc.command.X, tc.command.Y, tc.expectedErr, err)
			}
		} else {
			if err != nil {
				t.Errorf("For %d / %d, unexpected error: %v", tc.command.X, tc.command.Y, err)
			} else if result != tc.expected {
				t.Errorf("For %d / %d, expected result %d, but got %d", tc.command.X, tc.command.Y, tc.expected, result)
			}
		}
	}
}
