package tests

import (
	"fmt"
	"testing"

	"example.com/t/chess"
)

func TestGetPieceWithStrategy(t *testing.T) {

	pieceTye := ""
	_, err := chess.GetPieceWithStrategy(pieceTye)

	if err == nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error got %s", err))
	}

	pieceTye = "PAWN"
	piece, err := chess.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "KING"
	piece, err = chess.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "QUEEN"
	piece, err = chess.GetPieceWithStrategy(pieceTye)

	if err != nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be nil got %T", err))
	}

	if piece.Name != pieceTye {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected piece name to be %s got %s", pieceTye, piece.Name))
	}

	pieceTye = "BISHOP"
	piece, err = chess.GetPieceWithStrategy(pieceTye)

	if piece != nil || err == nil {
		t.Fatal(fmt.Sprintf("TestGetPieceStrategy failed, expected error to be Piece & Strategy doesn't exist got nil"))
	}
}
