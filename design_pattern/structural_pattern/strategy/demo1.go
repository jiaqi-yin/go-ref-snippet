package main

import "fmt"

type Context struct {
	Strategy
}

type Strategy interface {
	Do()
}

type Strategy1 struct{}

func (s1 *Strategy1) Do() {
	fmt.Println("do strategy1")
}

type Strategy2 struct{}

func (s2 *Strategy2) Do() {
	fmt.Println("do strategy2")
}

func main() {
	context := Context{}

	strategy1 := &Strategy1{}
	context.Strategy = strategy1
	context.Do()

	strategy2 := &Strategy2{}
	context.Strategy = strategy2
	context.Do()
}
