package main

import (
	"go-design-patterns/behavioral-patterns/command"
	"time"
)

func main() {
	queue := command.CommandQueue{}
	queue.AddCommand(command.NewOutputCommand("First message"))
	queue.AddCommand(command.NewTimePassedCommand(time.Now()))
	queue.AddCommand(command.NewOutputCommand("Third message"))
	queue.AddCommand(command.NewOutputCommand("Fourth message"))
	queue.AddCommand(command.NewOutputCommand("Fifth message"))
}
