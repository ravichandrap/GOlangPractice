package main

import "fmt"

func main() {
	x := []int{
		48, 96, 40, 86,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest int

	smallest = x[0]

	for _, value := range x {
		if value < smallest {
			smallest = value
		} else {
			continue
		}
	}

	fmt.Printf("Smallest %d", smallest)
}
