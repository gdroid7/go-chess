package chess

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/t/util"
)

const (
	MIN_CHAR_LEN_OF_PIECE = 4
	MAX_CHAR_LEN_OF_PIECE = 5
	MIN_MAX_POSITION_LEN  = 2
)

var PIECES = []string{"KING", "PAWN", "QUEEN"}

type Chessboard struct {
	Size int
}

func NewChessboard(size int) *Chessboard {
	return &Chessboard{Size: size}
}

func (cb Chessboard) ReadMove(args []string) ([]string, error) {

	if len(args) < 2 {
		return args, errors.New("Not enough arguments")
	}

	return args, nil
}

func (cb Chessboard) ValidateMove(pieceType string, position string) (bool, error) {

	if len(pieceType) < MIN_CHAR_LEN_OF_PIECE || len(pieceType) > MAX_CHAR_LEN_OF_PIECE {
		return false, errors.New(fmt.Sprintf("Invalid piece %s", pieceType))
	}

	if len(position) < MIN_MAX_POSITION_LEN || len(position) > MIN_MAX_POSITION_LEN {
		return false, errors.New(fmt.Sprintf("Invalid position %s", position))
	}

	//Can replace with a map search for piece name
	if s := strings.Join(PIECES, ""); !strings.Contains(s, pieceType) {
		return false, errors.New(fmt.Sprintf("Invalid piece \"%s\"", pieceType))
	}

	//Only get first 2 characters : max possibility
	posArray := strings.Split(position, "")

	if posIndex, err := strconv.Atoi(posArray[1]); err != nil {
		return false, errors.New(fmt.Sprintf("Invalid position \"%s\"", position))
	} else if posIndex < 1 || posIndex > cb.Size {
		return false, errors.New(fmt.Sprintf("Invalid position \"%s\", out of bound", position))
	}

	return true, nil
}

//Print a chessboard on cli with coordinates and cell names
//for reference
func (cb Chessboard) PrintChessboard() bool {

	fmt.Println(strings.Repeat("_", 12*cb.Size))

	for i := (cb.Size - 1); i >= 0; i-- {

		fmt.Print("|")

		for j := (cb.Size - 1); j >= 0; j-- {
			fmt.Printf(" [%d][%d]:%s%d |", i, (cb.Size - j - 1), util.ToChar(cb.Size-j-1), i+1)
		}

		fmt.Printf("\n%s", strings.Repeat("_", 12*cb.Size))

		if i != 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n\n")
	return true
}

func (cb Chessboard) GetCurrCoordinates(pos string) (int, int, error) {

	if len(pos) != MIN_MAX_POSITION_LEN {
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

func (cb Chessboard) GetPieceWithStrategy(pieceType string) (*Piece, error) {
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

func (cb Chessboard) HandleInput() (string, string, error) {

	var pieceType, position string

	//Read command line aruments except the first one
	//which is file namechessboard
	args, err := cb.ReadMove(os.Args[1:])

	if err != nil {
		return pieceType, position, errors.New(fmt.Sprintf("Error : %s\n", err.Error()))
	}

	//Capitalise ,Trim inputs
	pieceType, position = util.SanitizeInputs(args)

	//Check for out of bound moves invalid pieces
	if ok, err := cb.ValidateMove(pieceType, position); !ok {
		return pieceType, position, errors.New(fmt.Sprintf("Error : %s\n", err.Error()))
	}

	return pieceType, position, nil
}
