package main

import "fmt"

func main() {
	fooType()
}

func fooType() {
	Println("string")
	Println(2.33)
	Println(true)
	Println(202020)
}

func Println(i interface{}) {

	switch i.(type) {
	case string:
		fmt.Println(" type  is string ")
	case int:
		fmt.Println("type is int")
	case bool:
		fmt.Println("type is bool")
	case float64:
		fmt.Println("type is float64")
	}
}
