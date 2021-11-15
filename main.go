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
	pieceType, position := util.SanitizeInputs(args)

	//Check for out of bound moves invalid pieces
	if ok, err := util.ValidateMove(pieceType, position, len(Chessboard)); !ok {
		fmt.Printf("Error : %s\n", err.Error())
		return
	}

	//Get Piece sturct with moving strategy
	piece, err := strategy.GetPieceWithStrategy(pieceType)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error finding specified piece %s", pieceType))
	}

	//Get coordinates from cell position
	x, y, err := util.FindCoordinates(position)

	fmt.Printf("Input : %s %s\n", pieceType, position)

	fmt.Printf("Possible moves : %s\n", piece.FindPossibleMoves(Chessboard, x, y))

	return
}
