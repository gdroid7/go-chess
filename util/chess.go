package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"example.com/t/constants"
	"example.com/t/types"
)

var CharMap map[string]int

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
	BuildCharMap(size)
	return cb
}

func BuildCharMap(size int) {

	CharMap = make(map[string]int, size)

	for i := 0; i < size; i++ {
		CharMap[ToChar(i)] = i
	}

}

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

	// arg1Kind, arg2Kind := reflect.TypeOf(pieceType).Kind(), reflect.TypeOf(position).Kind()

	// if arg1Kind != reflect.String && arg2Kind != reflect.String {
	// 	return false, errors.New(fmt.Sprintf("Invalid input format %v %v", pieceType, position))
	// }

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

func PrintChessboard(Chessboard types.Chessboard) bool {

	fmt.Println(strings.Repeat("_", 12*len(Chessboard)))

	for i := (len(Chessboard) - 1); i >= 0; i-- {

		fmt.Print("|")

		for j := (len(Chessboard) - 1); j >= 0; j-- {
			fmt.Printf(" [%d][%d]:%s%d |", i, (len(Chessboard) - j - 1), ToChar(len(Chessboard)-j-1), i+1)
		}

		fmt.Printf("\n%s", strings.Repeat("_", 12*len(Chessboard)))

		if i != 0 {
			fmt.Println()
		}
	}
	fmt.Print("\n\n")
	return true
}

func FindCoordinates(pos string) (int, int, error) {

	if pos == "" {
		return 0, 0, errors.New(fmt.Sprintf("Invalid position %s", pos))
	}

	posArray := strings.Split(pos, "")

	y, ok := CharMap[posArray[0]]

	if !ok {
		return 0, 0, errors.New(fmt.Sprintf("Invalid position %s", posArray[0]))
	}

	x, err := strconv.Atoi(posArray[1])

	if err != nil {
		return 0, 0, errors.New(fmt.Sprintf("Invalid position %s", posArray[1]))
	}
	//x-1 to map cell numbers to indices
	return x - 1, y, nil
}
