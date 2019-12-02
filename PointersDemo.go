package main

import "fmt"

func main() {
	workWithPointers();
}
func workWithPointers() {
	var a int = 20

	ip := &a;
	fmt.Println(ip)
	fmt.Println(*ip)
}
