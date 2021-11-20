package tests

import (
	"fmt"
	"testing"

	"example.com/t/chess"
	"example.com/t/util"
)

func TestToInt(t *testing.T) {
	if v := util.ToInt(byte('A')); v != 0 {
		t.Fatal(fmt.Printf("TestToInt failed, expected value to be 0 , got %v", v))
	}
}

func TestFindCoordinatesTest(t *testing.T) {

	chessboard := chess.NewChessboard(8)

	_, _, err := chessboard.GetCurrCoordinates("")

	if err == nil {
		t.Fatal("TestFindCoordinatesTest failed, expected error got nil")
	}

	x, y, err := chessboard.GetCurrCoordinates("A1")

	if err != nil {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected result x=1 y=0 , got error %s", err.Error()))
	}

	if x != 0 || y != 0 {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected result x=0 y=0, got  x=%d y=%d", x, y))
	}

	x, y, err = chessboard.GetCurrCoordinates("Z1")

	if err != nil {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected error nil, got error %s ", err))
	}

	if x != 0 || y != 25 {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected x=0,y=25, got x=%d y=%d ", x, y))
	}

	x, y, err = chessboard.GetCurrCoordinates("AA")

	if err == nil {
		t.Fatal(fmt.Sprintf("TestFindCoordinatesTest failed, expected error, got nil"))
	}

}

func TestPrintChessboard(t *testing.T) {
	chessboard := chess.NewChessboard(8)
	if ok := chessboard.PrintChessboard(); !ok {
		t.Fatal("TestPrintChessboard failed, expected true, got false")
	}
}
