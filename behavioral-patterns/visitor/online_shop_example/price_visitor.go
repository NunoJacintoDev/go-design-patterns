package online_shop_example

import (
	"go-design-patterns/behavioral-patterns/visitor/online_shop_example/interfaces"
)

type PriceVisitor struct {
	Sum float32
}

func (pv *PriceVisitor) Visit(p interfaces.Retriever) {
	pv.Sum += p.GetPrice()
}
