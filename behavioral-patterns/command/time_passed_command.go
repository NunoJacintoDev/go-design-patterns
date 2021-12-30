package command

import (
	"fmt"
	"time"
)

type timePassedCommand struct {
	start time.Time
}

func NewTimePassedCommand(start time.Time) *timePassedCommand {
	return &timePassedCommand{start: start}
}

func (t *timePassedCommand) Execute() {
	fmt.Println(time.Since(t.start).String())
}
