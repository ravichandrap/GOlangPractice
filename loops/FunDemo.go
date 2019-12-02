package main

import "fmt"

func main() {
	anonymousFuc()
	sum, bigger := namedReturnFunc(10, 20, 5)
	fmt.Println(sum, bigger)
}

func namedReturnFunc(a, b, c int) (sum int, bigger string) {
	sum = a+b+c
	if a>b && a>c {
		bigger = "a is bigger"
	} else if b>a && b>c {
		bigger = "b is bigger"
	} else {
		bigger = "c is bigger"
	}

	return
}



func anonymousFuc()  {

	func(){
	 fmt.Println("foo")
	}()

	func(s string) {
		fmt.Println(s)
	}("Divija")

	func(value, s string) {
		fmt.Println(s, value)
	} ("one", "two")
}

func init(){
	fmt.Println(" Init method call")
}