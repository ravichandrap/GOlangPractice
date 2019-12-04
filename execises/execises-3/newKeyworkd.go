package main

import "fmt"

func main() {

	default1 := "first default value "
	default2 := new("first default value ")
	fmt.Println(changeValueWithPinter(*default1))
	fmt.Println(changeValueWithPinter(default2))
}

func changeValueWithPinter(value *string) {
	*value = "changed value here "
}
