package main

import (
	"fmt"
	"strings"
)

func strVariadic(s ...string) {
	fmt.Println(s)
	fmt.Println(strings.Join(s, ":"))
}

func interfaceVariadic(i ...interface{}) {
	fmt.Println("----------")
	fmt.Printf("number of arguments : %d\n", len(i))
	s := i[0].(Student)
	fmt.Println(s)
	fmt.Println("----------")
	c := i[1].(string)
	fmt.Println(c)
	fmt.Println("----------")
	col := i[2].(CollegeNew)
	fmt.Println(col)
}

type Student struct {
	name       string
	rollNumber int
}

type CollegeNew struct {
	name string
	code string
}

func main() {
	a, b, c, d, e := "one-", "two-", "three-", "four-", "five-"
	strVariadic(a, b, c, d, e)
	fmt.Println("--------------------")

	s := Student{"ravichandar", 2020}
	col := CollegeNew{}
	col.code = "1010101"
	col.name = "SVU"

	interfaceVariadic(s, c, col)
}
