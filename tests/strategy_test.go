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

var Chessboard = util.CreateEmptyChessboard(8)

func TestPawnWalk(t *testing.T) {

	pieceType := "PAWN"

	piece := strategy.Piece{
		Name:              pieceType,
		FindPossibleMoves: strategy.PawnWalk(),
	}

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	moves := "C3"
	result := piece.FindPossibleMoves(Chessboard, 1, 2)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected %s, got %s", moves, result))
	}
}

func TestKingWalk(t *testing.T) {

	pieceType := "KING"

	piece := strategy.Piece{
		Name:              pieceType,
		FindPossibleMoves: strategy.KingWalk(),
	}

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	moves := "G1,G2,H2"
	result := piece.FindPossibleMoves(Chessboard, 0, 7)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestKingWalk failed, expected %s, got %s", moves, result))
	}
}

func TestQueenWalk(t *testing.T) {

	pieceType := "QUEEN"

	piece := strategy.Piece{
		Name:              pieceType,
		FindPossibleMoves: strategy.QueenWalk(),
	}

	if piece.Name != pieceType {
		t.Fatal(fmt.Sprintf("TestQueenWalk failed, expected name to be %s, got %s", pieceType, piece.Name))
	}

	moves := "D5,C5,B5,A5,D6,C7,B8,A9,E6,E7,E8,F6,G7,H8,I9,F5,G5,H5,I5,F4,G3,H2,I1,E4,E3,E2,E1"
	result := piece.FindPossibleMoves(Chessboard, 4, 4)

	if result != moves {
		t.Fatal(fmt.Sprintf("TestQueenWalk failed, expected %s, got %s", moves, result))
	}
}
