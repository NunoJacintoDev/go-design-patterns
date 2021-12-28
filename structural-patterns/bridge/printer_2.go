package bridge

import (
	"errors"
	"fmt"
	"io"
)

type Printer2 struct {
	Writer io.Writer
}

func (p *Printer2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to Printer2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}
