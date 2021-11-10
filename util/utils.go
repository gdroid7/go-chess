package util

func CreateEmptyChessboard(r int, c int) [][]int {
	cb := make([][]int, r)
	for i := range cb {
		cb[i] = make([]int, c)
	}
	return cb
}
