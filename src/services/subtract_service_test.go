package services

import (
	"mosaic-go-interview/commands"
	"testing"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		command  commands.BasicArythemticCommand
		expected int
	}{
		{commands.BasicArythemticCommand{Action: "subtract", X: 5, Y: 2}, 3},    // positive numbers
		{commands.BasicArythemticCommand{Action: "subtract", X: 0, Y: 5}, -5},   // subtract from zero
		{commands.BasicArythemticCommand{Action: "subtract", X: -4, Y: -2}, -2}, // negative numbers
		{commands.BasicArythemticCommand{Action: "subtract", X: -2, Y: 3}, -5},  // subtracting a positive from a negative
		{commands.BasicArythemticCommand{Action: "subtract", X: 10, Y: -5}, 15}, // subtracting a negative from a positive
		{commands.BasicArythemticCommand{Action: "subtract", X: 0, Y: 0}, 0},    // both operands are zero
	}

	for _, tc := range testCases {
		result, err := Subtract(tc.command)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d - %d, expected result %d, but got %d", tc.command.X, tc.command.Y, tc.expected, result)
		}
	}
}
