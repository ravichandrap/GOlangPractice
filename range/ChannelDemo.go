package main

import "fmt"

func main() {
	myChannel := make(chan int)
	a := 20
	b := 30
	fmt.Println(myChannel)
	fooMyChannel(myChannel, a, b)
	fmt.Println(myChannel)
}

func fooMyChannel(myChannel chan int, a int, b int) {
	c := a + b
	myChannel <- c
}
