package main

import (
	"bytes"
	"fmt"
)

func main() {
	compareTwoSlice()
}

func compareTwoSlice()  {
	slice1 := []byte{'g','s','w','r','y'}
	slice2 := []byte{'m','p','v','q','e'}

	result := bytes.Compare(slice1, slice2)
	fmt.Println(result)
}
