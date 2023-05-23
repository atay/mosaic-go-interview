package services

import "mosaic-go-interview/commands"

func Add(command commands.BasicArythemticCommand) (int, error) {
	return command.X + command.Y, nil
}
