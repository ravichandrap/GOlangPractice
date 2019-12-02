package main

import "fmt"

type person struct {
	first string
	last string
}

func main() {
	p1:=person{"Ravi", "Reddy"}
	p2:=person{"divija", "shobha"}
	fmt.Println(p1, p2)
	ps :=[]person{p1}
	fmt.Println(ps)
	ps = append(ps,p2)
	fmt.Println(ps)
}