package main

import "fmt"

func main() {
	arr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven"}
	next := next(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(next())
	}

}

func next(arr []string) func() string {
	i := 0
	return func() (value string) {
		if i < len(arr) {
			value = arr[i]
		} else {
			value = "Unknown"
		}
		i++
		return
	}

}
