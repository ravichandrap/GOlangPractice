package main

import "fmt"

func main() {
	s := "Hello world"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
