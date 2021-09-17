package main

import "fmt"

type ICooler interface {
	turnon()
	work()
	turnoff()
}

type Cooler struct{}

func (c *Cooler) turnon() {
	fmt.Println("Turned on")
}

func (c *Cooler) work() {}

func (c *Cooler) turnoff() {
	fmt.Println("Turned off")
}

func doWork(cooler ICooler) {
	cooler.turnon()
	cooler.work()
	cooler.turnoff()
}

type Fan struct {
	Cooler
}

func (f *Fan) work() {
	fmt.Println("Fan is working")
}

type AirCon struct {
	Cooler
}

func (a *AirCon) work() {
	fmt.Println("AirCon is working")
}

func main() {
	fan := &Fan{}
	doWork(fan)

	airCon := &AirCon{}
	doWork(airCon)
}
