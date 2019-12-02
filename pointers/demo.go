package main
import "fmt"

func main(){
	x:=0
	foo(&x)
	fmt.Println("x after:",x)
	test:="ravichandra"
	doo(&test)
	fmt.Println(test)
}

func foo(x *int)  {
	fmt.Println("=======")
	fmt.Println(x)
	*x = 45
	fmt.Println(x)
	fmt.Println(*x)
}

func doo(x *string) {
	*x = *x+"my string"
}
