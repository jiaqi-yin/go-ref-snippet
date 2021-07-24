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

type productType int

const (
	p1 productType = iota
	p2
)

type productFactory struct{}

func (pf productFactory) Create(pt productType) Product {
	if pt == p1 {
		return &Product1{}
	}

	if pt == p2 {
		return &Product2{}
	}

	return nil
}

func main() {
	// product1 := &Product1{}
	// product1.SetName("p1")
	// fmt.Println(product1.GetName())

	// product2 := &Product2{}
	// product2.SetName("p2")
	// fmt.Println(product2.GetName())

	pf := productFactory{}

	product1 := pf.Create(p1)
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := pf.Create(p2)
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}
