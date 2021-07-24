package main

import "fmt"

func BinarySearch(slice []int, leftIndex int, rightIndex int, findVal int) (bool, int) {
	if leftIndex > rightIndex {
		return false, -1
	}

	middleIndex := (leftIndex + rightIndex) / 2
	if slice[middleIndex] > findVal {
		return BinarySearch(slice, leftIndex, middleIndex-1, findVal)
	} else if slice[middleIndex] < findVal {
		return BinarySearch(slice, middleIndex+1, rightIndex, findVal)
	} else {
		return true, middleIndex
	}
}

func main() {
	slice := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	findVal := 99
	found, findValIndex := BinarySearch(slice, 0, len(slice)-1, findVal)
	if found {
		fmt.Printf("Found value %v at index %v\n", findVal, findValIndex)
	} else {
		fmt.Printf("%v not found\n", findVal)
	}

}
