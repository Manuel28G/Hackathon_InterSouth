package main

import (
	"fmt"

	"../Sudoku/utils"
)

func main() {
	fmt.Println("Starting program...")

	board := Utils.ReadFile(Utils.PathTofile)

	board.FillInnerBox()

	board.FillEmptyPositions()

	board.Solve()

	//fmt.Println("Sudoku resuelto:")
	//fmt.Println(board.BoardToString())

}
