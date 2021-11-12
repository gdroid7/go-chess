package types

var PIECES map[string]bool

func init() {
	PIECES = map[string]bool{
		"KING":  false,
		"PAWN":  false,
		"QUEEN": false,
	}
}

type Chessboard [][]int

type CanMove func(pos string) bool

type Piece struct {
	CanMove
}
