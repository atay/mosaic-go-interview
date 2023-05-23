package services

import "mosaic-go-interview/src/commands"

func Subtract(command commands.BasicArythemticCommand) (int, error) {
	return command.X - command.Y, nil
}
