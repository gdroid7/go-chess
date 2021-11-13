package strategy

import (
	"errors"
)

type FindPossibleMoves func(pos string) bool

type Piece struct {
	Name              string
	FindPossibleMoves func(x int, y int)
	CurrentPos        string
}

func GetPieceWithStrategy(pieceType string) (Piece, error) {

	var piece Piece

	switch pieceType {
	case "PAWN":
		piece.Name = "PAWN"
		piece.FindPossibleMoves = PawnWalk()
		return piece, nil
	case "KING":
		piece.Name = "KING"
		piece.FindPossibleMoves = KingWalk()
		return piece, nil
	case "QUEEN":
		piece.Name = "QUEEN"
		piece.FindPossibleMoves = QueenWalk()
		return piece, nil
	default:
		piece.FindPossibleMoves = nil
		return piece, errors.New("Piece & Strategy doesn't exist")
	}
}

func PawnWalk() func(x int, y int) {
	return func(x int, y int) {

	}
}

func KingWalk() func(x int, y int) {
	return func(x int, y int) {

	}
}

func QueenWalk() func(x int, y int) {
	return func(x int, y int) {

	}
}
