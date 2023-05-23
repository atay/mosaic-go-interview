package services

import (
	"mosaic-go-interview/src/commands"
	"testing"
)

func TestMultiply(t *testing.T) {
	testCases := []struct {
		command  commands.BasicArythemticCommand
		expected int
	}{
		{commands.BasicArythemticCommand{Action: "multiply", X: 2, Y: 5}, 10},   // positive numbers
		{commands.BasicArythemticCommand{Action: "multiply", X: 0, Y: 5}, 0},    // one operand is zero
		{commands.BasicArythemticCommand{Action: "multiply", X: -3, Y: -4}, 12}, // negative numbers
		{commands.BasicArythemticCommand{Action: "multiply", X: 7, Y: -2}, -14}, // mixed signs
	}

	for _, tc := range testCases {
		result, err := Multiply(tc.command)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d * %d, expected result %d, but got %d", tc.command.X, tc.command.Y, tc.expected, result)
		}
	}
}
