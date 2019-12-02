package main

import "fmt"

func main() {
	i := uint(2)
	fact := factorial(i)
	fmt.Println(fact)
}

func factorial(f uint) uint {
	if f == 0 {
		return 1
	}
	return f * factorial(f-1)
}
