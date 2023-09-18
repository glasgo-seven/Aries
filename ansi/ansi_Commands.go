package Aries

import (
	"fmt"
)



func ANSI_ClearScreen() string {
	return ESC + CLEAR_SCREEN_TO_END + ESC + SCREEN_HOME
}



func ANSI_Colored16(color string) string {
	return ESC + color + STYLE
}

func ANSI_Colored16_R() string {
	return ESC + STYLE_COLOR_RESET + STYLE
}

func ANSI_Colored256_FG(color string) string {
	return ESC + STYLE_COLOR_256_FG + color + STYLE
}

func ANSI_Colored256_BG(color string) string {
	return ESC + STYLE_COLOR_256_BG + color + STYLE
}

func ANSI_ColoredRGB_FG(r string, g string, b string) string {
	return ESC + STYLE_COLOR_RGB_FG + r + ";" + g + ";" + b + STYLE
}

func ANSI_ColoredRGB_BG(r string, g string, b string) string {
	return ESC + STYLE_COLOR_RGB_BG + r + ";" + g + ";" + b + STYLE
}

func ANSI_CommandExecute(command string, x string) {
	fmt.Print(ESC + x + command)
}



func ANSI_AsAlert(message string) string {
	return ANSI_Colored16(STYLE_COLOR_FG_YELLOW) + message + ANSI_Colored16_R()
}

func ANSI_Alert(message string) {
	fmt.Print(ANSI_AsAlert("[*]\t" + message))
}

func ANSI_AlertLine(message string) {
	fmt.Println(ANSI_AsAlert("[*]\t" + message))
}

func ANSI_AsError(message string) string {
	return ANSI_Colored16(STYLE_COLOR_FG_RED) + message + ANSI_Colored16_R()
}

func ANSI_Error(message string) {
	fmt.Print(ANSI_AsError("[!]\t" + message))
}

func ANSI_ErrorLine(message string) {
	fmt.Println(ANSI_AsError("[!]\t" + message))
}

