package main

import "fmt"

func main() {
	b := read()
	fmt.Println(b)

	b1 := make([]byte, 32)
	read2(b1)
	fmt.Println(b1)
}

func read() []byte {
	b := make([]byte, 32)
	b[1] = '1'
	return b
}

func read2(b []byte) {
	b[1] = '2'
}
