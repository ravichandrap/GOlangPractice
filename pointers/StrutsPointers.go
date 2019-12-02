package main

import "fmt"

type myStruct struct {
	name int
}

func main() {
	var ms *myStruct
	ms = new(myStruct)
	fmt.Println(ms)
	fmt.Println(&ms)
	//foo2(ms)

	m :=  myStruct{name:34}
	m.foo3()
	//fmt.Println()
}
func foo2(ms myStruct)  {
	ms.name = 231
}

func (ms *myStruct) foo3() {
	fmt.Println(ms.name)
}