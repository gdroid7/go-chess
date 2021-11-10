package main

import (
	"testing"

	"example.com/t/util"
)

func TestCreateChessboard(t *testing.T) {
	row, col := 8, 8
	cb := util.CreateEmptyChessboard(row, col)
	if len(cb) < row && len(cb[0]) < col {
		t.Fatalf("CreateChessboard failed,expected chessboard of size %dx%d got %dx%d", 8, 8, len(cb), len(cb[0]))
	}
}
