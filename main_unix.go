//go:build !windows

package main

func fixWindowsConsole() {
	// Rien à faire ici, Linux et Mac parlent déjà l'UTF-8 nativement !
}
