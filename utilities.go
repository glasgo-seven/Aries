package Aries

import (

)

func Ternary(statement bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if statement {
		return ifTrue
	} else {
		return ifFalse
	}
}
