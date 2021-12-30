package command

import "fmt"

type Output struct {
	message string
}

func NewOutputCommand(message string) *Output {
	return &Output{
		message: message,
	}
}

func (c *Output) Execute() {
	fmt.Println(c.message)
}
