package tests

import (
	"fmt"
	"testing"

	"example.com/t/chess"
)

func TestPawnWalk(t *testing.T) {
	chessboard := chess.NewChessboard(8)

	pieceType := "PAWN"

	piece := chess.NewPiece(pieceType, chess.Pawn{})

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	x, y := 0, 0
	moves := "A2"
	result := piece.GetMoves(*chessboard, x, y)
	fmt.Println(x, y)
	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 7
	moves = ""
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 0, 7
	moves = "H2"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 0, 3
	moves = "D2"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 0
	moves = ""
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 4
	moves = ""
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}

}
