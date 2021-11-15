package tests

import (
	"fmt"
	"testing"

	"example.com/t/strategy"
	"example.com/t/util"
)

func TestGetPieceWithStrategy(t *testing.T) {

	pieceTye := ""
	_, err := strategy.GetPieceWithStrategy(pieceTye)

	if err == nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error got %s", err))
	}

	pieceTye = "PAWN"
	piece, err := strategy.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "KING"
	piece, err = strategy.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "QUEEN"
	piece, err = strategy.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "BISHOP"
	piece, err = strategy.GetPieceWithStrategy(pieceTye)

	if piece.Name != "" || err == nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be Piece & Strategy doesn't exist got nil"))
	}
}

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
	x, y := 0, 7 // H1
	expected := "G1,F1,E1,D1,C1,B1,A1,G2,F3,E4,D5,C6,B7,A8,H2,H3,H4,H5,H6,H7,H8"
	result := piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 7, 7 // H8
	expected = "G8,F8,E8,D8,C8,B8,A8,H7,H6,H5,H4,H3,H2,H1,G7,F6,E5,D4,C3,B2,A1"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 4, 4 // E5
	expected = "D5,C5,B5,A5"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 7, 0 // A8
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 2, 1 // B3
	expected = "A3"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 0, 0 // A1
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	// TOP
	x, y = 0, 0 // A1
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	//RIGHT DOWN DIAG TEST
	x, y = 7, 0 // A8
	expected = "B7,C6,D5,E4,F3,G2,H1"
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 0, 0 // A1
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 0, 7 // H1
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 0, 3 // D1
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

	x, y = 7, 7 // H8
	expected = ""
	result = piece.FindPossibleMoves(Chessboard, x, y)

	if result != expected {
		t.Fatal(fmt.Printf("Failed \n %s \n %s", result, expected))
	}

}
