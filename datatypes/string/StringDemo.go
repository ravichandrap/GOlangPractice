package main

import "fmt"

func main() {
	s := "Hello world"
	stringPrintlnChar(s)
	stringPrintlnV(s)
	stringPrintlnT(s)
}

func stringPrintln(s string) {

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	fmt.Println()
}

func stringPrintlnChar(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}

func stringPrintlnV(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v", s[i])
	}
	fmt.Println()
}

func stringPrintlnT(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T", s[i])
	}
	fmt.Println()
}
