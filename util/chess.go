package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"example.com/t/types"
)

func CreateEmptyChessboard(r int) types.Chessboard {
	//create empty chess board with size r x c
	cb := make([][]int, r, r)
	//initialise with cell names and int value to represent availibility (future)
	for v := range cb {
		//initialise map of key : value (A1:0)
		cb[v] = make([]int, r)
		//iterate over num of c
		// for i := 0; i < r; i++ {
		// 	cb[v][i] = i
		// }
	}
	return cb
}

func ReadMove(args []string) (string, string, error) {

	if len(args) < 2 {
		return "", "", errors.New("Not enough arguments")
	}

	args[0] = strings.Trim(args[0], ",")
	args[1] = strings.Trim(args[1], ",")

	arg1Kind, arg2Kind := reflect.TypeOf(args[0]).Kind(), reflect.TypeOf(args[1]).Kind()

	if arg1Kind != reflect.String && arg2Kind != reflect.String {
		return "", "", errors.New("Invalid input")
	}

	if len(args[0]) < 4 || len(args[1]) < 2 {
		return "", "", errors.New("Invalid input")
	}

	return args[0], args[1], nil
}

func ValidateMove(cb types.Chessboard, t string, p string) (bool, error) {

	if _, ok := types.PIECES[strings.ToUpper(t)]; !ok {
		return false, errors.New(fmt.Sprintf("Invalid piece \"%s\"", t))
	}

	posArray := strings.Split(p, "")

	posIndex, err := strconv.Atoi(posArray[1])

	if err != nil {
		return false, errors.New(fmt.Sprintf("Invalid position %s", p))
	}

	r := len(cb)

	if posIndex > r {
		return false, errors.New(fmt.Sprintf("Invalid position %s, out of bound", p))
	}

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
