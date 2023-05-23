package services

import (
	"mosaic-go-interview/commands"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		command  commands.BasicArythemticCommand
		expected int
	}{
		{commands.BasicArythemticCommand{Action: "add", X: 2, Y: 5}, 7},    // positive numbers
		{commands.BasicArythemticCommand{Action: "add", X: -3, Y: -4}, -7}, // negative numbers
		{commands.BasicArythemticCommand{Action: "add", X: -3, Y: 0}, -3},  // zero number
	}

	for _, tc := range testCases {
		result, err := Add(tc.command)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != tc.expected {
			t.Errorf("For %d + %d, expected result %d, but got %d", tc.command.X, tc.command.Y, tc.expected, result)
		}
	}
}
