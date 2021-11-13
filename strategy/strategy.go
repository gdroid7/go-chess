package strategy

import (
	"errors"
	"fmt"

	"example.com/t/util"
)

type FindPossibleMoves func(pos string) bool

type Piece struct {
	Name              string
	FindPossibleMoves func(x int, y int) string
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

func PawnWalk() func(x int, y int) string {
	return func(x int, y int) string {
		return fmt.Sprintf(fmt.Sprintf("%s%d", util.ToChar(y), x+2))
	}
}

func KingWalk() func(x int, y int) string {
	return func(x int, y int) string {
		return fmt.Sprintf("")
	}
}

func QueenWalk() func(x int, y int) string {
	return func(x int, y int) string {
		return fmt.Sprintf("")
	}
}
