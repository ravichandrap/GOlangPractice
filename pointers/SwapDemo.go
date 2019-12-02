package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b:= 10,20
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println("value: ",variadicFunc(a,b, 20, 30,40,50, 60))
	fmt.Println("value: ",variadicFunc())
	str := []string{"ravi", "divija", "shoba", "sarala", "Bhaskar"}
	fmt.Println(variadicFuncStr(str...))
}

func variadicFuncStr(values...string) string {
	return strings.Join(values, "-")
}

func swap(a, b *int)  {
	temp:= *a
	*a = *b
	*b = temp
}

func variadicFunc(values...int) int  {
	sum := 0
	for i, v := range values {
		fmt.Println(i,v)
		sum += v
	}
	return sum
}
