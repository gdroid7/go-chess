package tests

import (
	"fmt"
	"testing"

	"example.com/t/strategy"
)

func TestPawnWalk(t *testing.T) {

	pieceType := "PAWN"

	piece := strategy.Piece{
		Name:              pieceType,
		FindPossibleMoves: strategy.PawnWalk(),
	}

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	x, y := 0, 0
	moves := "A2"
	result := piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 7
	moves = ""
	result = piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 0, 7
	moves = "H2"
	result = piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 0, 3
	moves = "D2"
	result = piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 0
	moves = ""
	result = piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 4
	moves = ""
	result = piece.FindPossibleMoves(chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

}
