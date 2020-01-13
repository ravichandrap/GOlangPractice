package main

import (
	"fmt"
	"strings"
)

func main() {
	split()
}

func split() {
	myString := "one two three"

	myNewSlice := strings.Split(myString, " ")
	for i, v := range myNewSlice {
		fmt.Println(i, strings.ToUpper(v))
	}

}
