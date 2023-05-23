package services

import "mosaic-go-interview/commands"

func Multiply(command commands.BasicArythemticCommand) (int, error) {
	return command.X * command.Y, nil
}
