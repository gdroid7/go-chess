package tests

import (
	"fmt"
	"testing"

	"example.com/t/strategy"
	"example.com/t/util"
)

func TestQueenWalk(t *testing.T) {

	var Chessboard = util.CreateEmptyChessboard(8)

	pieceType := "QUEEN"

	piece := strategy.Piece{
		Name:              pieceType,
		FindPossibleMoves: strategy.QueenWalk(),
	}

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestQueenWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	// LEFT HOTIZ TEST
	x, y := 0, 0 // H1
	expected := "A2,A3,A4,A5,A6,A7,A8,B2,C3,D4,E5,F6,G7,H8,B1,C1,D1,E1,F1,G1,H1"
	result := piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("TestQueenWalk failed expected \n %s  got \n %s", expected, result))
	}

	x, y = 7, 7 // H8
	expected = "G8,F8,E8,D8,C8,B8,A8,H7,H6,H5,H4,H3,H2,H1,G7,F6,E5,D4,C3,B2,A1"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("TestQueenWalk failed expected \n %s  got \n %s", expected, result))
	}

	x, y = 4, 4 // E5
	expected = "D5,C5,B5,A5,D6,C7,B8,E6,E7,E8,F6,G7,H8,F5,G5,H5,F4,G3,H2,E4,E3,E2,E1,D4,C3,B2,A1"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 7, 0 // A8
	expected = "B8,C8,D8,E8,F8,G8,H8,B7,C6,D5,E4,F3,G2,H1,A7,A6,A5,A4,A3,A2,A1"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 0, 7 // B3
	expected = "G1,F1,E1,D1,C1,B1,A1,G2,F3,E4,D5,C6,B7,A8,H2,H3,H4,H5,H6,H7,H8"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

}
