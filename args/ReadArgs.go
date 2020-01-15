package main

import (
	"fmt"
	"os"
)

func main() {
	readArgsArray()
}

func readArgsArray() {
	args := os.Args
	for _, v := range args {
		fmt.Println(v)
	}
}

func readArgs() {
	onearg, argTwo := os.Args[1], os.Args[2]
	fmt.Println("first arg: ", onearg)
	fmt.Println("Seconds arg:", argTwo)
}
