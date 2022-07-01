package main

import (
	"fmt"
	. "sudoku/node"
)

func main() {
	input := []int{
		0, 0, 0, 0, 5, 0, 9, 2, 0,
		1, 0, 0, 0, 4, 2, 7, 6, 3,
		9, 0, 2, 0, 0, 7, 0, 0, 5,

		0, 0, 0, 0, 0, 3, 1, 5, 7,
		0, 5, 0, 6, 0, 9, 0, 8, 0,
		0, 0, 0, 5, 7, 0, 0, 0, 0,

		5, 0, 0, 0, 9, 8, 6, 0, 2,
		0, 2, 7, 3, 0, 1, 0, 0, 9,
		0, 4, 9, 7, 0, 0, 8, 3, 0,
	}

	current := CreateSudoku(input)

	current.PrintBoard()

	current.Solve()

	current.PrintBoard()

	for start := current; start != nil; start = start.Neighbourhood.Right {
		fmt.Printf("%d: %v \n", start.Value, start.TakenValues)
	}
}
