package tests

import (
	"errors"
	"testing"

	"example.com/t/types"
	"example.com/t/util"
)

func TestReadCommandLineVars(t *testing.T) {

	_, _, err := util.ReadMove([]string{})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("No input"), err.Error())
	}

	_, _, err = util.ReadMove([]string{"", ""})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("No input"), err.Error())
	}

	_, _, err = util.ReadMove([]string{"A", "B"})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("Invalid input"), err.Error())
	}

	_, _, err = util.ReadMove([]string{"Kind", "P1"})

	if err != nil {
		t.Fatalf("TestReadCommandLineVars failed,expected nil got %s", err.Error())
	}
}

func TestValidateMove(t *testing.T) {

	cb := types.Chessboard{{0, 0}, {0, 0}}

	ok, err := util.ValidateMove(cb, "Bishop", "ZZ1")

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, error nil got %s", ok, err.Error())
	}

	ok, err = util.ValidateMove(cb, "Bishop", "A3")

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, error nil got %s", ok, err.Error())
	}

	ok, err = util.ValidateMove(cb, "King", "A2")

	if !ok {
		t.Fatalf("TestValidateMove failed, expected status true got %t, error nil got %s", ok, err.Error())
	}
}

func TestSanitizeInput(t *testing.T) {

	sanitizedInputs := util.SanitizeInputs([]string{"kIng", "BISHOP", "a1", "B2"})

	if sanitizedInputs[0] != "KING" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "KING", sanitizedInputs[0])
	}

	if sanitizedInputs[1] != "BISHOP" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "BISHOP", sanitizedInputs[1])
	}

	if sanitizedInputs[2] != "A1" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "a1", sanitizedInputs[2])
	}

	if sanitizedInputs[3] != "B2" {
		t.Fatalf("TestSanitizeInput failed, expected %s got %s", "B2", sanitizedInputs[3])
	}

}