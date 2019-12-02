package main

import (
	"fmt"
	"reflect"
)

func main() {
	fooSwitch()
	fooSwitchExpresion()
	fooSwitchExpresionMulit()

	i:=10
	fooTypeSwithch(reflect.TypeOf(i).Kind())
	fooTypeSwithchStruct()
}

type User struct {
	name string
	id int
}

type Animal struct {
	name string
	canFly bool
}

func fooTypeSwithchStruct()  {
	var in User
	switch in {
	case Animal{}:
		fmt.Println("Animal ")
	case User{}:
		fmt.Println("User")
	}

}

func fooTypeSwithch(i interface{})  {
	fmt.Println(i)
	switch i {
	case reflect.Int:
		fmt.Println("int type")
	case i == "bool":
		fmt.Println("bool type")
	case i == "string":
		fmt.Println("string type")
	}
}

func fooSwitchExpresionMulit()  {
	value:="jayasri"

	switch value {
	case "jagan", "jayasri", "jyothis", "sohith":
		fmt.Println("Jagan family")
	case "ravi", "divija", "shobha":
		fmt.Println("Shobha family")

	}
}

func fooSwitch()  {
	switch day :=2; day {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		case 5:
			fmt.Println("5")

	}
}

func fooSwitchExpresion()  {
	value := 20
	switch  {
	case value == 10:
		fmt.Println("10")
	case value == 20:
		fmt.Println("20")
	}
}