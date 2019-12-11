package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	setEnv()
}

func setEnv() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if t := scanner.Text(); strings.Contains(t, "export") {
			env := strings.Split(strings.Split(t, " ")[1], "=")
			os.Setenv(env[0], env[1])
			fmt.Println(env[0], env[1])
		}

	}
}
