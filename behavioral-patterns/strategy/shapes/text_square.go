package shapes

import (
	"bytes"
	"go-design-patterns/behavioral-patterns/strategy"
	"io"
)

type TextSquare struct {
	strategy.Output
}

func (t *TextSquare) Print() error {
	r := bytes.NewReader([]byte("Square"))
	io.Copy(t.Writer, r)
	return nil
}
