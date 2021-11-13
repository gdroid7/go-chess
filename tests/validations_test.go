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
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("No input"), err.Error())
	}

	_, err = util.ReadMove([]string{"", ""})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("No input"), err.Error())
	}

	_, err = util.ReadMove([]string{"A", "B"})

	if err == nil {
		t.Fatalf("TestReadCommandLineVars failed,expected %s error got %s", errors.New("Invalid input"), err.Error())
	}

	_, err = util.ReadMove([]string{"Kind", "P1"})

	if err != nil {
		t.Fatalf("TestReadCommandLineVars failed,expected nil got %s", err.Error())
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

func TestValidateMove(t *testing.T) {

	cbLen := 2

	ok, err := util.ValidateMove("BI", "", cbLen)

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
	}

	ok, err = util.ValidateMove("BI", "A1", cbLen)

	if ok && err != nil {
		t.Fatalf("TestValidateMove failed, expected status false got %t, expected error nil got %s", ok, err.Error())
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

	if x != 1 || y != 0 {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected result x=1 y=0, got  x=%d y=%d", x, y))
	}
}
