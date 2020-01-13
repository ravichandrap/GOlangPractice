package main

import "fmt"

type MyRunner struct {
	run (string)
}

type animal struct {
	name string
}

func (a *animal) run(n string) {
	a.name = n
}

func (a animal) myName() {
	fmt.Println(a.name)
}

func printMyName(i interface{}) {
	i.(animal).myName()
}

func main() {
	a1 := animal{}
	a1.run("horse")

	a2 := animal{}
	a2.run("cow")

	printMyName(a1)
	printMyName(a2)
}
