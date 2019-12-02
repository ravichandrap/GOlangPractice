package main

import (
	"fmt"
	"sort"
)

func main() {
	fooPrintRandam()
}

func fooPrintRandam() {
	arr := []int{1100, 120, 100, 10, 110, 20, 99}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

	strs := []string{"c", "a", "b"}

	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)

	//for index, element := range arr {
	//
	//	fmt.Println(index, rand.Intn(element))
	//	fmt.Println(rand.Float32(), rand.Float64())
	//	fmt.Println("--------------")
	//}
}
