package main

import "fmt"

func main() {
	//arr:= []int {20,43,12,31,54,10,50,90}
	//foo_range_array(arr)

	//str := "divija reddy"
	//range_string(str)

	rangeWithMap()
}

func rangeWithMap() {
	myMap := make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	myMap["five"] = 5
	myMap["six"] = 6

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println(myMap)
	delete(myMap, "two")
	myMap["test"] = 202
	fmt.Println(myMap)
}

func rangeString(str string) {
	fmt.Println(str)
	for index, element := range str {
		fmt.Println(str[index], element)
	}
}

func fooRangeArray(arr []int) {
	fmt.Println(arr)
	for index, element := range arr {
		fmt.Println(index, element)
	}
}
