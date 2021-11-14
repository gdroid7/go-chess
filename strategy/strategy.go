package strategy

import (
	"errors"
	"fmt"

	"example.com/t/util"
)

type FindPossibleMoves func(pos string) bool

type Piece struct {
	Name              string
	FindPossibleMoves func(cb [][]int, x int, y int) string
	CurrentPos        string
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

		xMove := []int{-1, 0, 1, 1, 1, 0, -1, -1}
		yMove := []int{-1, -1, -1, 0, 1, 1, 1, 0}

		for i := 0; i < 8; i++ {

			// fmt.Printf("i:%d,x:%d,y:%d,xMove[i]:%d x+(xMove[i])) : %d, y+(yMove[i]) : %d \n", i, x, y, xMove[i], x+(xMove[i]), (y + (yMove[i])))

			if (x+(xMove[i])) >= 0 && (x+(xMove[i])) <= 8 && (y+(yMove[i])) >= 0 && (y+(xMove[i])) <= 8 {
				fmt.Println(util.ToChar(y+(yMove[i])), x+(xMove[i])+1)
			}
		}

		return fmt.Sprintf("%d,%d", x, y)
	}
}

func QueenWalk() func(cb [][]int, x int, y int) string {
	return func(cb [][]int, x int, y int) string {

		//LEFT HORIZ
		for i := (y - 1); i >= 0; i-- {
			// fmt.Println(x, i)
		}

		//LEFT TOP DIAG
		for i := 1; i <= y; i++ {
			// fmt.Println(x+i, y-i)
		}

		//TOP
		for i := 1; i <= y; i++ {
			// fmt.Println(x+i, y)
		}

		//RIGHT HORIZ
		for i := 1; i <= x; i++ {
			// fmt.Println(x, y+i)
		}

		//RIGHT TOP DIAG
		for i := 1; i <= x; i++ {
			// fmt.Println(x+i, y+i)
		}

		//RIGHT DOWN HORIZ
		for i := 1; i <= x; i++ {
			// fmt.Println(x-i, y+i)
		}

		//DOWN
		for i := 1; i <= x; i++ {
			fmt.Println(x-i, y)
		}

		return fmt.Sprintf("%d,%d", x, y)
	}
}
