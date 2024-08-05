package Aries

import (

)

const nil_	string = "";

func Ternary(statement bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if statement {
		return ifTrue
	} else {
		return ifFalse
	}
}
