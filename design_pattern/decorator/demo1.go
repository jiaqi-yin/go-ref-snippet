package main

import "fmt"

type Component interface {
	Operate()
}

type Component1 struct{}

func (c1 *Component1) Operate() {
	fmt.Println("c1 operate")
}

type Decorator interface {
	Component
	Do()
}

type Decorator1 struct {
	c Component
}

func (d1 *Decorator1) Do() {
	fmt.Println("d1 decorator1 do")
}

func (d1 *Decorator1) Operate() {
	d1.Do()
	d1.c.Operate()
}

func main() {
	c1 := &Component1{}
	c1.Operate()

	d1 := &Decorator1{}
	d1.c = c1
	d1.Operate()
}
