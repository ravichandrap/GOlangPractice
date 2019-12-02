package main

import "fmt"

func main() {
	foo()
	h := map[string]string{
		"auth": "test auth",
	}
	fmt.Println(h)

}

func foo() {
	population := make(map[string]int)
	population = map[string]int{
		"ap": 2020202,
		"tn": 30303030,
		"ka": 101010,
	}
	population["mr"] = 909090
	pop := population
	fmt.Println(population)
	delete(population, "ap")
	fmt.Println(pop)
	fmt.Println(population)
	p, ok := population["rr"]
	fmt.Println(p, ok)
	fmt.Println(len(population))
}
