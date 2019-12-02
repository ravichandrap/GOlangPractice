package main

import "fmt"

func main() {
	fooAppend()
}

func fooAppend() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))

}

func fooLenVsCap() {
	darr := [50]int{20, 304, 50, 60, 70, 80}
	fmt.Println(len(darr), cap(darr))

	d := darr[2:5]
	fmt.Println(len(d), cap(d))
}

func fooSlice() {
	darr := [...]int{20, 304, 50, 60, 70, 80}
	fmt.Println(darr)
	darr[2] = 30303030
	fmt.Println(darr)

	deDarr := darr[1:2]
	fmt.Println(deDarr)

	deDarr[0] = 999
	fmt.Println(darr)
	fmt.Println(deDarr)
}

func foo4() {
	arr := []int{2, 3, 4, 5, 6}
	fmt.Println(arr)
	foo5(arr)
	fmt.Println(arr)
}

func foo5(arr []int) {
	arr[2] = 303

	for v, i := range arr {
		fmt.Println(i, v)
	}
}

func foo3() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	fmt.Println(a)
	b[0] = "Andhara Pradesh"
	fmt.Println(a)
	fmt.Println(b)
}

func foo2() {
	for no, i := 2, 3; i < 20 && no < 30; i, no = i+1, no+1 {
		fmt.Println(no, i)
	}
}
func foo() {
	i := 0
	for i < 10 {
		fmt.Println("i is :", i)
		i++
	}
}
