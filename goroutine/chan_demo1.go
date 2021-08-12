package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "Tom"
	cat := Cat{"Tom", 1}
	allChan <- cat
	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)

	// Close channel before for-range
	// close(allChan)
	// for v := range allChan {
	// 	fmt.Println("v=", v)
	// }
}
