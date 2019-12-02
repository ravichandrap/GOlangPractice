package main

import "fmt"

func main() {
	//ifDemo()
	fooForLoop()
}

func ifDemo()  {

	population := make(map[string]int)
	population = map[string]int {
		"ap":2020202,
		"tn":30303030,
		"ka":101010,
	}
	fmt.Println(population)
	if pop, ok := population["ap"]; ok {
		fmt.Println(pop)
	}

}

func fooForLoop(){
	mymap := map[int]string {
		22: "one",
		33: "some value with 33",
		44:"another value with 44",
	}

	for key, value := range mymap {
		fmt.Println(key, value)
	}
}