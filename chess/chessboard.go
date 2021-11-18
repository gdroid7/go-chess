package chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"example.com/t/constants"
	"example.com/t/util"
)

func ReadMove(args []string) ([]string, error) {

	if len(args) < 2 {
		return args, errors.New("Not enough arguments")
	}

	return args, nil
}

func ValidateMove(pieceType string, position string, cbLen int) (bool, error) {

	if len(pieceType) < constants.MIN_CHAR_LEN_OF_PIECE || len(pieceType) > constants.MAX_CHAR_LEN_OF_PIECE {
		return false, errors.New(fmt.Sprintf("Invalid piece %s", pieceType))
	}

	if len(position) < constants.MIN_MAX_POSITION_LEN || len(position) > constants.MIN_MAX_POSITION_LEN {
		return false, errors.New(fmt.Sprintf("Invalid position %s", position))
	}

	//Can replace with a map search for piece name
	if s := strings.Join(constants.PIECES, ""); !strings.Contains(s, pieceType) {
		return false, errors.New(fmt.Sprintf("Invalid piece \"%s\"", pieceType))
	}

	//Only get first 2 characters : max possibility
	posArray := strings.Split(position, "")

	if posIndex, err := strconv.Atoi(posArray[1]); err != nil {
		return false, errors.New(fmt.Sprintf("Invalid position \"%s\"", position))
	} else if posIndex < 1 || posIndex > cbLen {
		return false, errors.New(fmt.Sprintf("Invalid position \"%s\", out of bound", position))
	}

	return true, nil
}

//Print a chessboard on cli with coordinates and cell names
//for reference
func PrintChessboard(size int) bool {

	fmt.Println(strings.Repeat("_", 12*size))

	for i := (size - 1); i >= 0; i-- {

		fmt.Print("|")

		for j := (size - 1); j >= 0; j-- {
			fmt.Printf(" [%d][%d]:%s%d |", i, (size - j - 1), util.ToChar(size-j-1), i+1)
		}

		fmt.Printf("\n%s", strings.Repeat("_", 12*size))

		if i != 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n\n")
	return true
}

func GetCurrCoordinates(pos string) (int, int, error) {

	if len(pos) != constants.MIN_MAX_POSITION_LEN {
		return 0, 0, errors.New(fmt.Sprintf("Invalid position %s", pos))
	}

	posArray := strings.Split(pos, "")

	y := util.ToInt(byte(posArray[0][0]))

	x, err := strconv.Atoi(posArray[1])

	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("Invalid position %s", posArray[1]))
	}
	//x-1 to map cell numbers to indices
	return x - 1, y, nil
}

func GetPieceWithStrategy(pieceType string) (*Piece, error) {
	switch pieceType {
	case "PAWN":
		return NewPiece(pieceType, Pawn{}), nil
	case "KING":
		return NewPiece(pieceType, King{}), nil
	case "QUEEN":
		return NewPiece(pieceType, Queen{}), nil
	default:
		return nil, errors.New("Invalid piece")
	}
}
