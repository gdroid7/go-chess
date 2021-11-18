package chess

import "example.com/t/util"

//Behaviour
type ChessboardMovement interface {
	GetPossibleMoves(cb [][]int, x int, y int) [][]int
}

//Concrete type
type Piece struct {
	Name               string
	ChessboardMovement ChessboardMovement
}

type Pawn Piece
type King Piece
type Queen Piece

func (p Pawn) GetPossibleMoves(cb [][]int, x int, y int) [][]int {
	moves := [][]int{}
	if x >= len(cb)-1 {
		return moves
	}
	return append(moves, []int{x + 1, y})
}

func (k King) GetPossibleMoves(cb [][]int, x int, y int) [][]int {
	moves := [][]int{}
	xMove := []int{0, 1, 1, 1, 0, -1, -1, -1}
	yMove := []int{-1, -1, 0, 1, 1, 1, 0, -1}

	for i := 0; i < 8; i++ {
		if (x+(xMove[i])) >= 0 && (x+(xMove[i])) < 8 && y+(yMove[i]) >= 0 && y+(yMove[i]) < 8 {
			moves = append(moves, []int{x + (xMove[i]), y + (yMove[i])})
		}
	}

	return moves
}

func (q Queen) GetPossibleMoves(cb [][]int, x int, y int) [][]int {
	moves := [][]int{}
	xMove := []int{
		0, 0, 0, 0, 0, 0, 0, //LEFT
		1, 2, 3, 4, 5, 6, 7, //TOP LEFT DIAG
		1, 2, 3, 4, 5, 6, 7, //TOP
		1, 2, 3, 4, 5, 6, 7, //TOP RIGHT DIAG
		0, 0, 0, 0, 0, 0, 0, // RIGHT
		-1, -2, -3, -4, -5, -6, -7, // RIGHT DOWN DIAG
		-1, -2, -3, -4, -5, -6, -7, // DOWN
		-1, -2, -3, -4, -5, -6, -7} // DOWN LEFT DIAG

	yMove := []int{
		-1, -2, -3, -4, -5, -6, -7, //LEFT
		-1, -2, -3, -4, -5, -6, -7, //TOP LEFT DIAG
		0, 0, 0, 0, 0, 0, 0, //TOP
		1, 2, 3, 4, 5, 6, 7, //TOP RIGHT DIAG
		1, 2, 3, 4, 5, 6, 7, // RIGHT
		1, 2, 3, 4, 5, 6, 7, // RIGHT DOWN DIAG
		0, 0, 0, 0, 0, 0, 0, // DOWN
		-1, -2, -3, -4, -5, -6, -7} // DOWN LEFT DIAG

	//LEFT TOP DIAG
	for i := 0; i < (len(cb)*len(cb))-len(cb); i++ {
		if (x+(xMove[i])) >= 0 && (x+(xMove[i])) < 8 && (y+(yMove[i])) >= 0 && (y+(yMove[i])) < 8 {
			moves = append(moves, []int{x + (xMove[i]), y + (yMove[i])})
		}
	}

	return moves
}

func (p *Piece) GetMoves(cb [][]int, x int, y int) string {
	return util.CoordinatesForHumans(p.ChessboardMovement.GetPossibleMoves(cb, x, y))
}

//Create new piece concrete object with specified interface implementation
func NewPiece(name string, cm ChessboardMovement) *Piece {
	return &Piece{
		Name:               name,
		ChessboardMovement: cm,
	}
}
