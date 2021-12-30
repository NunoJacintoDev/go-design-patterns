package products

import "go-design-patterns/behavioral-patterns/visitor/online_shop_example/interfaces"

type Fridge struct {
	Product
}

func (f *Fridge) GetPrice() float32 {
	return f.Product.Price + 20
}

func (f *Fridge) Accept(v interfaces.Visitor) {
	v.Visit(f)
}
