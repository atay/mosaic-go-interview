package services

import "mosaic-go-interview/commands"

type CannotDivideByZeroError struct{}

func (e CannotDivideByZeroError) Error() string {
	return "Cannot divide by zero"
}

func (e CannotDivideByZeroError) UserMessage() string {
	return e.Error()
}

func Divide(command commands.BasicArythemticCommand) (int, error) {
	if command.Y == 0 {
		return 0, CannotDivideByZeroError{}
	}
	return command.X / command.Y, nil
}
