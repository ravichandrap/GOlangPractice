package main

import "fmt"

func main() {
	foo()
}

func foo()  {
	x:=make([]int,10,10)
	x[0] = 20
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
