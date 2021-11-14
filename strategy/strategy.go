package strategy

import (
	"errors"
	"fmt"
	"strings"

	"example.com/t/util"
)

type FindPossibleMoves func(pos string) bool

type Piece struct {
	Name              string
	FindPossibleMoves func(cb [][]int, x int, y int) string
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
		return piece, errors.New("Piece & Strategy doesn't exist")
	}
}

func PawnWalk() func(cb [][]int, x int, y int) string {
	return func(cb [][]int, x int, y int) string {
		return fmt.Sprintf(fmt.Sprintf("%s%d", util.ToChar(y), x+2))
	}
}

func KingWalk() func(cb [][]int, x int, y int) string {
	return func(cb [][]int, x int, y int) string {
		moves := [][]int{}
		xMove := []int{0, 1, 1, 1, 0, -1, -1, -1}
		yMove := []int{-1, -1, 0, 1, 1, 1, 0, -1}

		for i := 0; i < 8; i++ {

			if (x+(xMove[i])) >= 0 && (x+(xMove[i])) < 8 && y+(yMove[i]) >= 0 && y+(yMove[i]) < 8 {
				moves = append(moves, []int{x + (xMove[i]), y + (yMove[i])})
			}

			// if y > 0 && y <= 8 {

			// 	fmt.Println(x-1, y-1)
			// 	fmt.Println(x, y-1)
			// 	fmt.Println(x+1, y-1)
			// 	fmt.Println(x+1, y)

			// 	if y < (8-1) && y < 7 {
			// 		fmt.Println(x+1, y+1)
			// 		fmt.Println(x, y+1)
			// 		fmt.Println(x-1, y+1)
			// 		fmt.Println(x-1, y)
			// 		break
			// 	}
			// 	break
			// }

		}

		return CoordinatesForHumans(moves)
	}
}

func QueenWalk() func(cb [][]int, x int, y int) string {
	return func(cb [][]int, x int, y int) string {

		moves := [][]int{}
		// //LEFT HORIZ 1
		// for i := x + 1; i <= y; i++ {
		// 	fmt.Println(i, y-i)
		// }

		//LEFT HORIZ 1
		for i := 1; i <= y; i++ {
			// fmt.Println(x, y-i)
			moves = append(moves, []int{x, y - i})
			// fmt.Println(CoordinatesForHumans(x, y-i))
		}

		//LEFT TOP DIAG
		for i := 1; i <= y; i++ {
			// fmt.Println(x+i, y-i)
			// fmt.Println(CoordinatesForHumans(x+i, y-i))
			moves = append(moves, []int{x + i, y - i})

		}

		//TOP
		//iterate for max len - current pos
		for i := 1; i <= len(cb)-(x+1); i++ {
			// fmt.Println(x+i, y)
			// fmt.Println(CoordinatesForHumans(x+i, y))
			moves = append(moves, []int{x + i, y})

		}

		//RIGHT TOP DIAG
		//iterate until y and add 1 to x and y every iteration
		for i := 1; i <= y; i++ {
			// fmt.Println(x+i, y+i)
			// fmt.Println(CoordinatesForHumans(x+i, y+i))
			moves = append(moves, []int{x + i, y + i})

		}

		//RIGHT HORIZ
		for i := 1; i <= x; i++ {
			// fmt.Println(x, y+i)
			// fmt.Println(CoordinatesForHumans(x, y+i))
			moves = append(moves, []int{x, y + i})

		}

		//RIGHT DOWN HORIZ
		for i := 1; i <= x; i++ {
			// fmt.Println(x-i, y+i)
			// fmt.Println(CoordinatesForHumans(x-i, y+i))
			moves = append(moves, []int{x - i, y + i})

		}

		//DOWN
		for i := 1; i <= x; i++ {
			// fmt.Println(x-i, y)
			// fmt.Println(CoordinatesForHumans(x-i, y))
			moves = append(moves, []int{x - i, y})

		}
		return CoordinatesForHumans(moves)
	}
}

func CoordinatesForHumans(moves [][]int) string {

	var result string
	for i := 0; i < len(moves); i++ {
		result += fmt.Sprintf("%s%d, ", util.ToChar(moves[i][1]), moves[i][0]+1)
	}
	return strings.TrimRight(result, ", ")
}
