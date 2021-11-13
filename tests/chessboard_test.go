package tests

import (
	"testing"

	"example.com/t/util"
)

func TestCreateChessboard(t *testing.T) {

	cb := util.CreateEmptyChessboard(0)

	if len(cb) > 0 {
		t.Fatalf("CreateChessboard failed,expected cb to be nil got %v", cb)
	}

	size := 2
	cb = util.CreateEmptyChessboard(size)

	if len(cb) != size {
		t.Fatalf("CreateChessboard failed,expected chessboard of size %dx%d got %dx%d", size, size, len(cb), len(cb))
	}
}
