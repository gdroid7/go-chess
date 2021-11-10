package main

import (
	"fmt"

	"example.com/t/util"
)

var Chessboard [][]int

func main() {
	Chessboard := util.CreateEmptyChessboard(8, 8)
	fmt.Println(Chessboard)
}
