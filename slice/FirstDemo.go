package main

import "fmt"

func main() {
	foo()
}

func foo()  {
	a := []int{1,2,3,4,5,6,7,8,9}
	b:=a[:]
	c:=a[3:]
	d:=a[:6]
	e:=a[3:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}