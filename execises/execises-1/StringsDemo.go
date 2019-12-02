package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fooStringDemo1()
	//fooStringDemo2()
	//fooStringDemo3()
	//fooStringDemo4()

	//a,b := 20,30
	//aR, bR := fooStringDemo5(a, b)
	//fmt.Println(aR,bR)

	//result:=foo_Variadic(a,b,aR,bR)
	//fmt.Println(result)

	//intSeq := intSeq()
	//fmt.Println(intSeq)
	//fmt.Println(intSeq())
	//fmt.Println(intSeq())
	//fmt.Println(intSeq())
	fact := fact(10)
	fmt.Println(fact)
}

func fact(i int) int {
	if i == 0 {
		return 1
	}

	return i * fact(i-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func foo_Variadic(i ...int) int {
	sum := 0
	for _, num := range i {
		sum += num
	}
	return sum
}

func fooStringDemo5(a int, b int) (int, int) {
	return b, a
}

func fooStringDemo4() {
	str := "14"
	output, error := strconv.Atoi(str)
	fmt.Println("output", output)
	fmt.Println("Error:", error)
}

func fooStringDemo3() {
	s1 := "ravi,divija,shobha,sothith,jyothish,bhaskar,sarala"
	s1 = strings.ToUpper(s1)
	output := strings.Split(s1, ",")
	fmt.Print(output)

}

func fooStringDemo2() {
	var array = []string{"one", "two", "Three", "four"}
	output := strings.Join(array, "-000-")
	fmt.Println(output)
}
func fooStringDemo1() {
	var str1 = "golang"
	var str2 = "Tutorial"

	var output = strings.Join([]string{str1, str2}, "|")
	fmt.Println(output)
}
