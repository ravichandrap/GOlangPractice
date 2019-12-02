package main

import "fmt"

type name string

func (n name)printName()  {
	fmt.Println(n)
}

func (n *name) changeName(v name)  {
	*n = v
}

func main() {
	n := name("test")
	n.printName()

	n.changeName("chagne the value ")
	n.printName()

}
