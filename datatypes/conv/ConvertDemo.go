package main

import (
	"fmt"
	"strconv"
)

func main() {
	strVar := "100"
	intVar, error := strconv.Atoi(strVar)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(intVar)

	strVar1 := "-52541"
	intVar1, _ := strconv.ParseInt(strVar1, 10, 32)
	fmt.Println(intVar1)

	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
	fmt.Println(intVar2)

	// intVar := 1055
	strVar3 := strconv.FormatInt(int64(intVar), 10)
	fmt.Println(strVar3)

}
