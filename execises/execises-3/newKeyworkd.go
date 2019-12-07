package main

import "fmt"

func mainW() {

	default1 := "first default value "
	fmt.Println(changeValueWithPinter(default1))
}

func changeValueWithPinter(value string) string {
	value = " changed value here "
	return value
}
