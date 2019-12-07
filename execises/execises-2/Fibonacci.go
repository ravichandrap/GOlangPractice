package main

import "fmt"

func main() {
	num := 3
	f := fib(num)
	fmt.Println(f)
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-1) + fib(num-2)
}
