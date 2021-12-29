package shapes

import (
	"fmt"
	"go-design-patterns/behavioral-patterns/strategy"
	"os"
)

const (
	TEXT_STRATEGY  = "text"
	IMAGE_STRATEGY = "image"
)

func NewPrinter(s string) (strategy.PrintStrategy, error) {
	switch s {
	case TEXT_STRATEGY:
		return &TextSquare{
			Output: strategy.Output{
				LogWriter: os.Stdout,
			},
		}, nil
	case IMAGE_STRATEGY:
		return &ImageSquare{
			Output: strategy.Output{
				LogWriter: os.Stdout,
			},
		}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found\n", s)
	}
}
