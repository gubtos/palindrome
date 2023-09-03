package main

import (
	"fmt"
	"os"

	"github.com/gubtos/golib"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("execute da seguinte forma:\npalindrome \"a sacada da casa\"")

		return
	}
	str := os.Args[1]
	if isPalidrome := golib.IsPalindrome(str); isPalidrome {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}
