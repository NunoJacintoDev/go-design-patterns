package online_shop_example

import (
	"fmt"
	"go-design-patterns/behavioral-patterns/visitor/online_shop_example/interfaces"
)

type NamePrinterVisitor struct {
	ProductList string
}

func (n *NamePrinterVisitor) Visit(p interfaces.Retriever) {
	n.ProductList = fmt.Sprintf("%s\n%s", p.GetName(), n.ProductList)
}
