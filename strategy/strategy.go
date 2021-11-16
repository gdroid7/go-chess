package strategy

import (
	"errors"
	"fmt"
	"strings"

	"example.com/t/util"
)

type FindPossibleMoves func(pos string) bool

type Piece struct {
	Name string
	//pasing cb to this function if we need to check the values of indexes for blocking places
	//otherwise chessboard size is enough to print output
	FindPossibleMoves func(cb [8][8]int, x int, y int) string
}

func GetPieceWithStrategy(pieceType string) (Piece, error) {

	var piece Piece

	switch pieceType {

	case "PAWN":
		piece.Name = pieceType
		piece.FindPossibleMoves = PawnWalk()
		return piece, nil
	case "KING":
		piece.Name = pieceType
		piece.FindPossibleMoves = KingWalk()
		return piece, nil
	case "QUEEN":
		piece.Name = pieceType
		piece.FindPossibleMoves = QueenWalk()
		return piece, nil
	default:
		piece.FindPossibleMoves = nil
		return piece, errors.New(fmt.Sprintf("%s Piece & Strategy doesn't exist", pieceType))
	}
}

func PawnWalk() func(cb [8][8]int, x int, y int) string {
	return func(cb [8][8]int, x int, y int) string {

		if x >= len(cb)-1 {
			return ""
		}
		return fmt.Sprintf(fmt.Sprintf("%s%d", util.ToChar(y), x+2))
	}
}

func KingWalk() func(cb [8][8]int, x int, y int) string {
	return func(cb [8][8]int, x int, y int) string {

		moves := [][]int{}
		xMove := []int{0, 1, 1, 1, 0, -1, -1, -1}
		yMove := []int{-1, -1, 0, 1, 1, 1, 0, -1}

		for i := 0; i < 8; i++ {
			if (x+(xMove[i])) >= 0 && (x+(xMove[i])) < 8 && y+(yMove[i]) >= 0 && y+(yMove[i]) < 8 {
				moves = append(moves, []int{x + (xMove[i]), y + (yMove[i])})
			}
		}

		return CoordinatesForHumans(moves)
	}
}

func QueenWalk() func(cb [8][8]int, x int, y int) string {
	return func(cb [8][8]int, x int, y int) string {

		// 	moves := [][]int{}
		// 	//LEFT HORIZ
		// 	for i := 1; i <= y; i++ {
		// 		moves = append(moves, []int{x, y - i})
		// 	}

		// 	//LEFT TOP DIAG
		// 	for i := 1; x < len(cb)-1 && y < len(cb)-1 &&
		// 		i <= len(cb)-(x) && i <= len(cb)-(y+1); i++ {
		// 		moves = append(moves, []int{x + i, y - i})
		// 	}
		// 	//TOP
		// 	//iterate for max len - current pos
		// 	for i := 1; i <= len(cb)-(x+1); i++ {
		// 		moves = append(moves, []int{x + i, y})
		// 	}

		// 	//RIGHT TOP DIAG
		// 	//iterate until y and add 1 to x and y every iteration
		// 	for i := 1; x < len(cb)-1 && y < len(cb)-1 &&
		// 		i <= len(cb)-(x+1) && i <= len(cb)-(y+1); i++ {
		// 		moves = append(moves, []int{x + i, y + i})
		// 	}

		// 	//RIGHT HORIZ
		// 	for i := 1; i <= len(cb)-(y+1); i++ {
		// 		moves = append(moves, []int{x, y + i})
		// 	}

		// 	//RIGHT DOWN DIAG
		// 	for i := 1; x > 0 && y < len(cb)-i && i <= x; i++ {
		// 		moves = append(moves, []int{x - i, y + i})
		// 	}

		// 	//DOWN
		// 	for i := 1; i <= x; i++ {
		// 		moves = append(moves, []int{x - i, y})
		// 	}

		// 	//LEFT DOWN DIAG
		// 	for i := 1; y > 0 && x > 0 && i <= x; i++ {
		// 		moves = append(moves, []int{x - i, y - i})
		// 	}

		// 	return CoordinatesForHumans(moves)
		// }

		//Alternate bruth force solution
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

		return CoordinatesForHumans(moves)
	}
}

func CoordinatesForHumans(moves [][]int) string {
	var result string
	for i := 0; i < len(moves); i++ {
		result += fmt.Sprintf("%s%d,", util.ToChar(moves[i][1]), moves[i][0]+1)
	}
	return strings.TrimRight(result, ",")
}
