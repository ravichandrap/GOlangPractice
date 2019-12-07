package main

import "fmt"

func main() {
	var arrNum [13]int
	arrNum[2] = 3

	i := 0
	for i < len(arrNum) {
		if i > 10 {
			continue
		}
		arrNum[i] = i + 38
		i++
	}

	for i, x := range arrNum {
		fmt.Printf("a[%d] = %d\n", i, x)
	}

	fmt.Println(arrNum)
	fmt.Println(arrNum[:])

	fmt.Println("Hello world")
}
