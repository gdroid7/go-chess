package tests

import (
	"testing"

	"example.com/t/util"
)

func TestCreateChessboard(t *testing.T) {
	row := 8
	cb := util.CreateEmptyChessboard(row)

	if len(cb) < row && len(cb[0]) < row {
		t.Fatalf("CreateChessboard failed,expected chessboard of size %dx%d got %dx%d", 8, 8, len(cb), len(cb[0]))
	}
}
