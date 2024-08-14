package Aries

import (
	_ "fmt"
	"strings"
	"math"
)

const RUNES string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ChangeBase(_number int, _base byte) string {
	var intBase		int		= int(_base)
	if intBase > len(RUNES) {
		panic("Provided _base is bigger than 36 ( 0-9A-Z )")
	}

	var isNegative	bool	= _number < 0
	if isNegative {
		_number = -_number
	}

	var newNumber	string	= ""
	for _number != 0 {
		newNumber = string(RUNES[_number % intBase]) + newNumber
		_number /= intBase
	}

	if isNegative {
		return "-" + newNumber
	}
	return newNumber
}

func ChangeBase_CustomRunes(_number int, _runes string) string {
	var isNegative	bool	= _number < 0
	if isNegative {
		_number = -_number
	}

	var newNumber	string	= ""
	var intBase		int		= len(_runes)
	for _number != 0 {
		newNumber = string(_runes[_number % intBase]) + newNumber
		_number /= intBase
	}

	if isNegative {
		return "-" + newNumber
	}
	return newNumber
}

func Bin(_number int) string {
	return ChangeBase(_number, 2)
}
func Oct(_number int) string {
	return ChangeBase(_number, 8)
}
func Hex(_number int) string {
	return ChangeBase(_number, 16)
}


func Dec(_number string, _base byte) int {
	var isNegative bool = _number[0] == 45
	if isNegative {
		_number = _number[1:]
	}

	var size		int		= len(_number)
	var number		int		= 0
	for i := 0; i < size; i++ {
		runeValue := strings.Index(RUNES, string(_number[i]))
		number += runeValue * int(math.Pow(float64(_base), float64(size - i - 1)))
	}

	if isNegative {
		return -number
	}
	return number
}