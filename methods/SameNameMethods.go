package main

import "fmt"

type student struct {
	name string
	branch string
}

type teacher struct {
	language string
	marks int
}

func (s student) show() {
	fmt.Println(s)
}

func (t teacher) show() {
	fmt.Println(t)
}

func main() {
	s := student{
		name:   "Ravichadnra",
		branch: "CSE",
	}

	t := teacher{
		language: "Golang",
		marks:    100,
	}

	s.show()
	t.show()
}