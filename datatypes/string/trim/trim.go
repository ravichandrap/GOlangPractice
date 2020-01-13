package main

import (
	"fmt"
	"strings"
)

func main() {
	trimSpace()
	fmt.Println("------------------")
	trimNewLine()
	fmt.Println("------------------")
}

func trimSpace() {
	s := "some thing to trim "
	fmt.Println(s, "<")

	trimS := strings.Trim(s, " ")
	fmt.Println(trimS, "<")
}

func trimNewLine() {
	s := "some thing to trim new line\n"
	fmt.Println(s)

	trimS := strings.Trim(s, "\n")
	fmt.Println(trimS)
}
