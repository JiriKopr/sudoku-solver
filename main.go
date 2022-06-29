package main

import (
	"fmt"
	"strings"
	. "sudoku/node"
)

func main() {
	testValues := []string{}

	for i := 0; i < 81; i++ {
		testValues = append(testValues, fmt.Sprintf("%d", i))
	}

	current := CreateSudokuFromString(strings.Join(testValues, ";"))

	for first := current; current != nil; first = first.Neighbourhood.Bottom {
		for ; current != nil; current = current.Neighbourhood.Right {
			fmt.Print(current.Value)
		}

        fmt.Print("\n")

		current = first.Neighbourhood.Bottom
	}

}
