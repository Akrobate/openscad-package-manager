//go:build windows

package main

import (
	"syscall"
)

func fixWindowsConsole() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleOutputCP := kernel32.NewProc("SetConsoleOutputCP")
	// 65001 = UTF-8 chez Microsoft
	_, _, _ = setConsoleOutputCP.Call(uintptr(65001))
}
