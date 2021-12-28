package bridge

import "fmt"

type Printer1 struct{}

func (p *Printer1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}
