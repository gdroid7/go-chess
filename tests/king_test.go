package tests

import (
	"fmt"
	"testing"

	"example.com/t/chess"
)

func TestKingWalk(t *testing.T) {

	chessboard := chess.NewChessboard(8)

	pieceType := "KING"

	piece := chess.NewPiece(pieceType, chess.King{})

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	x, y := 0, 0
	moves := "A2,B2,B1"
	result := piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 0, 7
	moves = "G1,G2,H2"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 7
	moves = "G8,H7,G7"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 3
	moves = "C8,E8,E7,D7,C7"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 7, 0
	moves = "B8,B7,A7"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

	x, y = 3, 3
	moves = "C4,C5,D5,E5,E4,E3,D3,C3"
	result = piece.GetMoves(*chessboard, x, y)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}

}
