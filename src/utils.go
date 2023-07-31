package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J") // escape codes para limpar a tela (Unix)
}

func Menu() {
	fmt.Println("Aperte enter para voltar.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	Clear()
}
