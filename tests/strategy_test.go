package tests

import (
	"fmt"
	"testing"

	"example.com/t/strategy"
	"example.com/t/util"
)

func TestGetPieceStrategy(t *testing.T) {

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

	pieceTye = "BISHOP"
	piece, err = strategy.GetPieceWithStrategy(pieceTye)

	if piece.Name != "" || err == nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be Piece & Strategy doesn't exist got nil"))
	}
}

func TestPawnWalk(t *testing.T) {

	piece := strategy.Piece{
		Name:              "PAWN",
		FindPossibleMoves: strategy.PawnWalk(),
	}

	if piece.Name != "PAWN" {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected name to be PAWN, got %s", piece.Name))
	}

	cb := util.CreateEmptyChessboard(8)
	result := piece.FindPossibleMoves(cb, 1, 2)

	if result != "C3" {
		t.Fatal(fmt.Sprintf("TestPawnWalk failed, expected C3, got %s", result))
	}
}
