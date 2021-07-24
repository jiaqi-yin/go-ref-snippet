package main

import "fmt"

func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	arr := [5]int{21, 12, 210, 100, 62}
	slice := arr[:]
	fmt.Println("Before sorting: ", slice)
	bubbleSort(slice)
	fmt.Println("After sorting: ", slice)
}
