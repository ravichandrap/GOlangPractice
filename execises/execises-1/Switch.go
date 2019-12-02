package main

import "fmt"

func main() {
	arr := []int{20, 30, 40, 50, 60, 3, 1, 0, 2}
	myFunction(arr)
}

func myFunction(arr []int) {

	for r, i := range arr {
		switch r {
		case 20:
			fmt.Printf("i = %d, range = %d \n ", i, r)

		case 0:
			fmt.Printf("i = %d, range = %d \n ", i, r)

		case 60:
			fmt.Printf("i = %d, range = %d \n ", i, r)

		case 50:
			fmt.Printf("i = %d, range = %d \n ", i, r)

		case 30:
			fmt.Printf("i = %d, range = %d \n ", i, r)
		default:
			fmt.Printf("default : i = %d, range = %d \n ", i, r)
		}
	}
}
