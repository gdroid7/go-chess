package tests

import (
	"errors"
	"fmt"
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

	ok, err := util.ValidateMove("", "", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BI got nil", ok)
	}

	ok, err = util.ValidateMove("KI", "A0", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece KI got nil", ok)
	}

	ok, err = util.ValidateMove("KING", "A", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	ok, err = util.ValidateMove("KING", "A11", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid position A got nil", ok)
	}

	ok, err = util.ValidateMove("KGNI", "A1", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid1 piece KNIN got nil", ok)
	}

	ok, err = util.ValidateMove("123", "123", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid input format 123 123, got nil", ok)
	}

	ok, err = util.ValidateMove("BISHOP", "A1", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error to be Invalid piece BISHOP, got nil", ok)
	}

	ok, err = util.ValidateMove("QUEEN", "AA", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"AA\" got nil", ok)
	}

	ok, err = util.ValidateMove("QUEEN", "A0", cbLen)

	if err == nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error Invalid position \"A99\" got nil", ok)
	}

	ok, err = util.ValidateMove("BISHOP", "ZZ1", cbLen)

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	ok, err = util.ValidateMove("BISHOP", "A3", cbLen)

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	ok, err = util.ValidateMove("KING", "A2", cbLen)

	if !ok {
		t.Fatalf("TestValidateMove failed, expected status true got %t, expected error nil got %s", ok, err.Error())
	}

}

func TestFindCoordinatesTest(t *testing.T) {

	_, _, err := util.FindCoordinates("")

	if err == nil {
		t.Fatal("TestFindCoordinatesTest failed, expected error got nil")
	}

	x, y, err := util.FindCoordinates("A1")

	if err != nil {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected result x=1 y=0 , got error %s", err.Error()))
	}

	if x != 0 || y != 0 {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected result x=0 y=0, got  x=%d y=%d", x, y))
	}

	x, y, err = util.FindCoordinates("Z1")

	if err == nil {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected error, got nil"))
	}
}
