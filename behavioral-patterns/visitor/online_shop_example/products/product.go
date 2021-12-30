package products

import "go-design-patterns/behavioral-patterns/visitor/online_shop_example/interfaces"

type Product struct {
	Price float32
	Name  string
}

func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) Accept(v interfaces.Visitor) {
	v.Visit(p)
}
func (p *Product) GetName() string {
	return p.Name
}
