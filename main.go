package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"example.com/t/chess"
	"example.com/t/util"
)

var size = 8
var chessboard [][]int

func main() {
	chessboard = make([][]int, size)
	//Print chessboad for reference
	chess.PrintChessboard(size)
	//Run program
	run(chessboard)
}

func run(cb [][]int) {

	pieceType, position, err := handleInput()

	if err != nil {
		log.Fatal(err)
		return
	}

	//Get Piece sturct with moving strategy
	piece, err := chess.GetPieceWithStrategy(pieceType)

	if err != nil {
		log.Fatal(err)
		return
	}

	//Get coordinates from cell position
	x, y, err := chess.GetCurrCoordinates(position)

	//Print ioutput
	fmt.Printf("Input : %s %s\n", pieceType, position)
	fmt.Printf("Possible moves : %s\n", piece.GetMoves(cb, x, y))

	return
}

func handleInput() (string, string, error) {

	var pieceType, position string

	//Read command line aruments except the first one
	//which is file namechessboard
	args, err := chess.ReadMove(os.Args[1:])

	if err != nil {
		return pieceType, position, errors.New(fmt.Sprintf("Error : %s\n", err.Error()))
	}

	//Capitalise ,Trim inputs
	pieceType, position = util.SanitizeInputs(args)

	//Check for out of bound moves invalid pieces
	if ok, err := chess.ValidateMove(pieceType, position, size); !ok {
		return pieceType, position, errors.New(fmt.Sprintf("Error : %s\n", err.Error()))
	}

	return pieceType, position, nil
}
