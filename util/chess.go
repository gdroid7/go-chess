package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"example.com/t/constants"
	"example.com/t/types"
)

var MIN_POSITION_LEN = 2
var MAX_POSITION_LEN = 2

func CreateEmptyChessboard(size int) [][]int {

	if size < 1 {
		return nil
	}

	cb := make([][]int, size)
	//create empty chess board with size r x c
	//initialise with cell names and int value to represent availibility (future)
	for v := range cb {
		//initialise map of key : value (A1:0)
		//iterate over num of c
		for i := 0; i < size; i++ {
			cb[v] = append(cb[v], 0)
		}
	}

	MIN_POSITION_LEN = len(cb) % 10

	return cb
}

func ReadMove(args []string) ([]string, error) {

	if len(args) < 2 {
		return args, errors.New("Not enough arguments")
	}

	args[0] = strings.Trim(args[0], ",")
	args[1] = strings.Trim(args[1], ",")

	if len(args[0]) < constants.MIN_CHAR_LEN_OF_PIECE || len(args[1]) < 2 {
		return args, errors.New("Invalid input")
	}

	arg1Kind, arg2Kind := reflect.TypeOf(args[0]).Kind(), reflect.TypeOf(args[1]).Kind()

	if arg1Kind != reflect.String && arg2Kind != reflect.String {
		return args, errors.New("Invalid input")
	}

	return args, nil
}

func ValidateMove(pieceType string, position string, cbLen int) (bool, error) {

	if s := strings.Join(constants.PIECES, ""); !strings.Contains(s, pieceType) {
		return false, errors.New(fmt.Sprintf("Invalid piece \"%s\"", pieceType))
	}

	if len(position) < MIN_POSITION_LEN && len(position) > MAX_POSITION_LEN {
		return false, errors.New(fmt.Sprintf("Invalid position %s", position))
	}

	//Only get first 2 characters : max possibility
	posArray := strings.Split(position, "")

	posIndex, err := strconv.Atoi(posArray[1])

	if err != nil {
		return false, errors.New(fmt.Sprintf("Invalid position %s", position))
	}

	if posIndex > cbLen {
		return false, errors.New(fmt.Sprintf("Invalid position %s, out of bound", position))
	}

	//TODO
	//HANDLE Alphabets as columns

	return true, nil
}

func PrintChessboard(Chessboard types.Chessboard) {

	fmt.Print("\n------BOARD------\n\n")

	fmt.Println(strings.Repeat("_", 12*len(Chessboard)))

	for i := (len(Chessboard) - 1); i >= 0; i-- {

		fmt.Print("|")

		for j := (len(Chessboard) - 1); j >= 0; j-- {
			fmt.Printf(" [%d][%d]:%s |", i, (len(Chessboard) - j - 1), ToChar(len(Chessboard)-j-1, i))
		}

		fmt.Printf("\n%s", strings.Repeat("_", 12*len(Chessboard)))

		if i != 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n\n")
}

func FindCoordinates(pos string) (int, int, error) {

	return 0, 0, nil
}
