package Aries

import (
	// "fmt"
	"os/exec"
	// "syscall"
	"strings"
)

var (
	// user32 = syscall.NewLazyDLL("User32.dll")
	// getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	// SM_CXSCREEN = 0
	// SM_CYSCREEN = 1

	app string = "powershell"
	cmdGetWidth string = "$Host.UI.RawUI.WindowSize.Width"
	cmdGetHeight string = "$Host.UI.RawUI.WindowSize.Height"
	args string = ""
)



func trimOfNewLine(str string) string {
	return strings.Trim(str, "\n")
}

// func GetSystemMetrics(nIndex int) int {
//	index := uintptr(nIndex)
//	ret, _, _ := getSystemMetrics.Call(index)
//	return int(ret)
// }

// func GetScreenSize() (int, int) {
//	return GetSystemMetrics(SM_CXSCREEN), GetSystemMetrics(SM_CYSCREEN)
// }

func GetTerminalWidth() string {
	out, _ := exec.Command(app, cmdGetWidth, args).Output()
	// fmt.Println(err)
	return trimOfNewLine(string(out))
}

func GetTerminalHeight() string {
	out, _ := exec.Command(app, cmdGetHeight, args).Output()
	// fmt.Println(err)
	return trimOfNewLine(string(out))
}

func GetTerminalSize() (string, string) {
	return GetTerminalWidth(), GetTerminalHeight()
}