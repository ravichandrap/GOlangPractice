package main

import "fmt"

type Doctor struct {
	number int
	actorName string
	companions []string
}
func main() {
	aDoctor := Doctor {
		number:     20202,
		actorName:  "ravichandra",
		companions: []string {
			"one",
			"two",
			"three",
		},
	}
	fmt.Println(aDoctor)

	students := struct {
		name string
		rollNumber int
	}{"ravi", 303}

	fmt.Println(students)
}
