package main

import "fmt"

type Product interface {
	SetName(string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "Product 1 name = " + p1.name
}

type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "Product 2 name = " + p2.name
}

type ProductFactory interface {
	Create() Product
}

type Product1Factory struct{}

func (p1f *Product1Factory) Create() Product {
	return &Product1{}
}

type Product2Factory struct{}

func (p2f *Product2Factory) Create() Product {
	return &Product2{}
}

func main() {
	product1Factory := &Product1Factory{}
	p1 := product1Factory.Create()
	p1.SetName("p1")
	fmt.Println(p1.GetName())

	product2Factory := &Product2Factory{}
	p2 := product2Factory.Create()
	p2.SetName("p2")
	fmt.Println(p2.GetName())
}
