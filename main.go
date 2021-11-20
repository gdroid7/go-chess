package main

import (
	"fmt"
	"log"

	"example.com/t/chess"
)

func main() {
	chessboard := chess.NewChessboard(8)
	chessboard.PrintChessboard()
	run(*chessboard)
}

func run(cb chess.Chessboard) {

	pieceType, position, err := cb.HandleInput()

	if err != nil {
		log.Fatal(err)
		return
	}

	//Get Piece sturct with moving strategy
	piece, err := cb.GetPieceWithStrategy(pieceType)

	if err != nil {
		log.Fatal(err)
		return
	}

	//Get coordinates from cell position
	x, y, err := cb.GetCurrCoordinates(position)

	//Print ioutput
	fmt.Printf("Input : %s %s\n", pieceType, position)
	fmt.Printf("Possible moves : %s\n", piece.GetMoves(cb, x, y))

	return
}
