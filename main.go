package main

import (
	"fmt"
	"os"

	"example.com/t/strategy"

	"example.com/t/util"
)

var Chessboard [][]int

func init() {
	Chessboard = util.CreateEmptyChessboard(8)
	util.PrintChessboard(Chessboard)
}

func main() {
	//Read command line aruments except the first one
	//which is file name
	args, err := util.ReadMove(os.Args[1:])

	if err != nil {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	//Capitalise ,Trim inputs
	args = util.SanitizeInputs(args)

	//Check for out of bound moves invalid pieces
	if ok, err := util.ValidateMove(args[0], args[1], len(Chessboard)); !ok {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	pieceType := args[0]

	piece, err := strategy.GetPieceWithStrategy(pieceType)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error finding specified piece %s", pieceType))
	}

	position := args[1]

	x, y, err := util.FindCoordinates(position)

	fmt.Println(piece.FindPossibleMoves(Chessboard, x, y))

	return
}
