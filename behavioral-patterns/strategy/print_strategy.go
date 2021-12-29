package strategy

import (
	"io"
)

type PrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}

type Output struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (d *Output) SetLog(w io.Writer) {
	d.LogWriter = w
}
func (d *Output) SetWriter(w io.Writer) {
	d.Writer = w
}
