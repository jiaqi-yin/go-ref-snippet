package main

import "fmt"

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 1. Original array
	var chessMap [11][11]int
	chessMap[1][2] = 1 // Black
	chessMap[2][3] = 2 // White
	// 2. Output array
	fmt.Println("---- Original array ----")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	// 3. Transform to sparse array
	var sparseArr []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	fmt.Println("---- Sparse array ----")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	// 4. Restore to original array
	var chessMap2 [11][11]int
	sparseArr = sparseArr[1:]
	for _, valNode := range sparseArr {
		chessMap2[valNode.row][valNode.col] = valNode.val
	}

	fmt.Println("---- Restored original array ----")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}
