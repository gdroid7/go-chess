package tests

import (
	"errors"
	"testing"

	"example.com/t/util"
)

func TestReadMove(t *testing.T) {

	_, err := util.ReadMove([]string{})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got nil", errors.New("No input"))
	}

	_, err = util.ReadMove([]string{"King", "P1"})

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

	cbLen := 2

	if ok, err := util.ValidateMove("", "", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BI got nil", ok)
	}

	if ok, err := util.ValidateMove("KI", "A0", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece KI got nil", ok)
	}

	if ok, err := util.ValidateMove("KING", "A", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	if ok, err := util.ValidateMove("KING", "A11", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	if ok, err := util.ValidateMove("KGNI", "A1", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid1 piece KNIN got nil", ok)
	}

	if ok, err := util.ValidateMove("123", "123", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid input format 123 123, got nil", ok)
	}

	if ok, err := util.ValidateMove("BISHOP", "A1", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BISHOP, got nil", ok)
	}

	if ok, err := util.ValidateMove("QUEEN", "AA", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"AA\" got nil", ok)
	}

	if ok, err := util.ValidateMove("QUEEN", "A0", cbLen); err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"A99\" got nil", ok)
	}

	if ok, err := util.ValidateMove("BISHOP", "ZZ1", cbLen); ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	if ok, err := util.ValidateMove("BISHOP", "A3", cbLen); ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	if ok, err := util.ValidateMove("KING", "A2", cbLen); !ok {
		t.Fatalf("TestValidateMove failed, expected status true got %t, expected error nil got %s", ok, err.Error())
	}
}
