package main

import (
	"fmt"
	"os"

	"example.com/t/types"

	"example.com/t/util"
)

var Chessboard types.Chessboard

func init() {
	Chessboard = util.CreateEmptyChessboard(8)
	util.PrintChessboard(Chessboard)
}

func main() {

	piece, pos, err := util.ReadMove(os.Args[1:])

	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	if ok, err := util.ValidateMove(Chessboard, piece, pos); !ok {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	fmt.Println(piece, pos)
	return
}
