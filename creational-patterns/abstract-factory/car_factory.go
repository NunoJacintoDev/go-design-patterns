package abstract_factory

import (
	"errors"
	"fmt"
)

type CarType = int

const (
	LuxuryCarType CarType = iota
	FamilyCarType
)

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
