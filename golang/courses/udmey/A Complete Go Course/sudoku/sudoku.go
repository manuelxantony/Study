package main

import (
	"fmt"
)

func main() {
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}
	displayBoard(puzz)
	solvePuzzle(puzz)
	fmt.Println("Puzzle Solved")
	displayBoard(puzz)
}

func solvePuzzle(puzz [][]int) bool {
	row, column := getEmptySpace(puzz)
	if row == -1 {
		return true
	}

	for i := 1; i <= 9; i++ {
		if isNumValid(puzz, i, row, column){
			puzz[row][column] = i
			
			fmt.Println()
			displayBoard(puzz)
			fmt.Println()
			
			if solvePuzzle(puzz) {
				return true
			}

			puzz[row][column] = 0
		}
	}
	return false
}

func displayBoard(puzz [][]int) {
	col_len := len(puzz[0])
	row_len := len(puzz)
	for i := 0; i < row_len; i++ {
		for j := 0; j < col_len; j ++{
			fmt.Print(puzz[i][j], " ")
		}
		fmt.Println()
	}

}

func getEmptySpace(puzz [][]int) (int, int) {
	col_len := len(puzz[0])
	row_len := len(puzz)

	for i := 0; i < row_len; i++ {
		for j := 0; j < col_len; j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNumValid(puzz [][]int, guess int, row int, column int) bool {
	for index := range puzz {
		if ((puzz[row][index] == guess && column != index) || (puzz[index][column] == guess && row != index)) {
			// guess found
			return false
		}
	}
	// row - (row %3) + cycle
	// col - (col%3) + cycle
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			current_row := row - (row%3) + k
			current_column := column - (column%3) + l
			if puzz[current_row][current_column] == guess && 
			((current_row != row) || (current_column != column)) {	
				return false
			}
		}
	}
	return true
}