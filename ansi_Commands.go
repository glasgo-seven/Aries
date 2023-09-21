package Aries

import (
	"fmt"
)



func ANSI_ClearScreen() {
	fmt.Print(ESC + CLEAR_SCREEN_TO_END + ESC + SCREEN_HOME)
}



func ANSI_SetColor16(color string) {
	fmt.Print(ESC + color + STYLE)
}
func ANSI_SetColor16_string(color string) string {
	return ESC + color + STYLE
}

func ANSI_ResetColor16() {
	fmt.Print(ESC + STYLE_COLOR_RESET + STYLE)
}
func ANSI_ResetColor16_string() string {
	return ESC + STYLE_COLOR_RESET + STYLE
}

func ANSI_SetColor256_FG(color string) {
	fmt.Print(ESC + STYLE_COLOR_256_FG + color + STYLE)
}
func ANSI_SetColor256_FG_string(color string) string {
	return ESC + STYLE_COLOR_256_FG + color + STYLE
}

func ANSI_SetColor256_BG(color string) {
	fmt.Print(ESC + STYLE_COLOR_256_BG + color + STYLE)
}
func ANSI_SetColor256_BG_string(color string) string {
	return ESC + STYLE_COLOR_256_BG + color + STYLE
}

func ANSI_SetColorRGB_FG(r string, g string, b string) {
	fmt.Print(ESC + STYLE_COLOR_RGB_FG + r + ";" + g + ";" + b + STYLE)
}
func ANSI_SetColorRGB_FG_string(r string, g string, b string) string {
	return ESC + STYLE_COLOR_RGB_FG + r + ";" + g + ";" + b + STYLE
}

func ANSI_SetColorRGB_BG(r string, g string, b string) {
	fmt.Print(ESC + STYLE_COLOR_RGB_BG + r + ";" + g + ";" + b + STYLE)
}
func ANSI_SetColorRGB_BG_string(r string, g string, b string) string {
	return ESC + STYLE_COLOR_RGB_BG + r + ";" + g + ";" + b + STYLE
}

func ANSI_ResetColor() {
	fmt.Print(ESC + STYLE_COLOR_RESET + STYLE)
}
func ANSI_ResetColor_string() string {
	return ESC + STYLE_COLOR_RESET + STYLE
}



func ANSI_FormatString(str string, format string, reset string) string {
	return ESC + format + str +  ESC + reset
}
func ANSI_CommandExecute(command string, x string) {
	fmt.Print(ESC + x + command)
}



func ANSI_AsAlert(message string) string {
	return ANSI_SetColor16_string(STYLE_COLOR_FG_YELLOW) + message + ANSI_ResetColor_string()
}

func ANSI_Alert(message string) {
	fmt.Print(ANSI_AsAlert("[*]\t" + message))
}

func ANSI_AlertLine(message string) {
	fmt.Println(ANSI_AsAlert("[*]\t" + message))
}

func ANSI_AsError(message string) string {
	return ANSI_SetColor16_string(STYLE_COLOR_FG_RED) + message + ANSI_ResetColor_string()
}

func ANSI_Error(message string) {
	fmt.Print(ANSI_AsError("[!]\t" + message))
}

func ANSI_ErrorLine(message string) {
	fmt.Println(ANSI_AsError("[!]\t" + message))
}



func ANSI_Bold(str string) string {
	return ANSI_FormatString(str, STYLE_BOLD, STYLE_BOLD_RESET)
}

func ANSI_Dim(str string) string {
	return ANSI_FormatString(str, STYLE_DIM, STYLE_DIM_RESET)
}

func ANSI_Italic(str string) string {
	return ANSI_FormatString(str, STYLE_ITALIC, STYLE_ITALIC_RESET)
}

func ANSI_UnderLine(str string) string {
	return ANSI_FormatString(str, STYLE_UNDER, STYLE_UNDER_RESET)
}

func ANSI_Blink(str string) string {
	return ANSI_FormatString(str, STYLE_BLINK, STYLE_BLINK_RESET)
}

func ANSI_InverseColor(str string) string {
	return ANSI_FormatString(str, STYLE_INVERSE, STYLE_INVERSE_RESET)
}

func ANSI_Hidden(str string) string {
	return ANSI_FormatString(str, STYLE_HIDDEN, STYLE_HIDDEN_RESET)
}

func ANSI_Strike(str string) string {
	return ANSI_FormatString(str, STYLE_STRIKE, STYLE_STRIKE_RESET)
}
