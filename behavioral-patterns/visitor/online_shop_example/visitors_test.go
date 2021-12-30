package online_shop_example

import (
	"fmt"
	"go-design-patterns/behavioral-patterns/visitor/online_shop_example/interfaces"
	item "go-design-patterns/behavioral-patterns/visitor/online_shop_example/products"
	"testing"
)

func Test(t *testing.T) {
	products := make([]interfaces.Visitable, 3)
	products[0] = &item.Rice{
		Product: item.Product{
			Price: 32.0,
			Name:  "Some rice",
		}}
	products[1] = &item.Pasta{
		Product: item.Product{
			Price: 40.0,
			Name:  "Some pasta",
		}}
	products[2] = &item.Fridge{
		Product: item.Product{
			Price: 50,
			Name:  "A fridge",
		}}

	//Print the sum of prices
	priceVisitor := &PriceVisitor{}
	for _, p := range products {
		p.Accept(priceVisitor)
	}
	var expectedTotal float32 = 142.00
	if priceVisitor.Sum != expectedTotal {
		t.Errorf("Price Visitor with wrong total price! expected: %f\n actual: %f\n", expectedTotal, priceVisitor.Sum)
	}

	//Print the products list
	nameVisitor := &NamePrinterVisitor{}
	for _, p := range products {
		p.Accept(nameVisitor)
	}
	expectedStringList := "A fridge\nSome pasta\nSome rice\n"
	if nameVisitor.ProductList != expectedStringList {
		t.Errorf("Price Visitor with wrong total price! expected: %s\n actual: %s\n", expectedStringList, nameVisitor.ProductList)

	}

	fmt.Printf("Total: %f\n", priceVisitor.Sum)
	fmt.Printf("\nProduct list:\n-------------\n%s", nameVisitor.ProductList)

}
