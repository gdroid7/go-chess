package tests

import (
	"errors"
	"testing"

	"example.com/t/chess"
	"example.com/t/util"
)

func TestReadMove(t *testing.T) {

	chessboard := chess.NewChessboard(8)

	_, err := chessboard.ReadMove([]string{})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got nil", errors.New("No input"))
	}

	_, err = chessboard.ReadMove([]string{"King", "P1"})

	if err != nil {
		t.Fatalf("TestReadCommandLineVars failed,expected nil got %s", err.Error())
	}
}

func TestSanitizeInput(t *testing.T) {

	pType, pos := util.SanitizeInputs([]string{"kIng", "a1"})

	if pType != "KING" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "KING", pType)
	}

	if pos != "A1" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "A1", pos)
	}

}

func TestValidateMove(t *testing.T) {

	chessboard := chess.NewChessboard(8)

	if ok, err := chessboard.ValidateMove("", ""); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BI got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("KI", "A0"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece KI got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("KING", "A"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("KING", "A11"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("KGNI", "A1"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid1 piece KNIN got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("123", "123"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid input format 123 123, got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("BISHOP", "A1"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BISHOP, got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("QUEEN", "AA"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"AA\" got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("QUEEN", "A0"); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"A99\" got nil", ok)
	}

	if ok, err := chessboard.ValidateMove("BISHOP", "ZZ1"); ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	if ok, err := chessboard.ValidateMove("BISHOP", "A3"); ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	if ok, err := chessboard.ValidateMove("KING", "A2"); !ok {
		t.Fatalf("TestValidateMove failed, expected status true got %t, expected error nil got %s", ok, err.Error())
	}
}
