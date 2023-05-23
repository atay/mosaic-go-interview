package commands

import "fmt"

type BasicArythemticCommand struct {
	Action string `json:"action"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func (c BasicArythemticCommand) GetCacheKey() string {
	return fmt.Sprintf("%s-%s-%s", c.Action, fmt.Sprint(c.X), fmt.Sprint(c.Y))
}
