package main

import "fmt"

func main() {
	var arr = []int{100, 2, 0, 40, 999, 210}
	var f float32 = getAverage(arr)
	fmt.Println(f)
}

func getAverage(arr []int) float32 {
	fmt.Println("getAverage")
	var sum, i int
	var l int = len(arr)
	for i = 0; i < l; i++ {
		sum += arr[i]
	}

	return float32(sum / l)
}
