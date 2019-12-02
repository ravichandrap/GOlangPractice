package main

import "fmt"

func main() {
	//fooVariadic(1,2,3,4,5,6,7,8,)
	fooMapRef()
}

func fooMapRef() {
	emp := map[int]string{20: "firs value", 30: "second value", 40: " 404040404", 50: "5050505050", 60: "606060606"}
	emp2 := emp
	fmt.Println(emp)
	fmt.Println(emp2)
	delete(emp2, 40)
	fmt.Println(emp)
	fmt.Println(emp2)

}

func fooMap2() {
	emp := map[int]string{20: "firs value", 30: "second value", 40: " 404040404", 50: "5050505050", 60: "606060606"}
	fmt.Println(len(emp))

	for key, value := range emp {
		fmt.Println(key, value)
	}

}

func fooMap() {
	var students map[int]string
	fmt.Println(students == nil)
	students = make(map[int]string)

	students[2] = "Ravichandra"
	students[5] = "test"
	students[4] = "chandra"
	students[3] = "Ravic"

	fmt.Println(students)
	delete(students, 4)
	fmt.Println(students)
}

func fooVariadic(a int, b ...int) {
	fmt.Println(a, b)
	fmt.Println(b)
}
