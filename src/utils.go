package main

import "fmt"

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}
