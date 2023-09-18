package Aries

import (

)

func Bold(str string) string {
	return ANSI_FormatString(str, STYLE_BOLD, STYLE_BOLD_RESET)
}
