package main

import "fmt"

func main() {
	fooPanic()
}

func fooPanic() {
	fmt.Println("Please enter your value!")
	var myValue int
	fmt.Scanln(&myValue)

	switch myValue {
	case 1:
		fmt.Println("this is 1 action ")
	case 2:
		fmt.Println("thi is second action ")
	default:
		panic("Can not take action for your option!")
	}
	fooPanic()
}
