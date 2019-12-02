package main

import "fmt"

type Student struct {
	name string
	rollNumber int
}

type School struct {
	schoolName string
	schoolNumber int
}

type city struct {
	Student
	School
	cityName string
}

func (c city)printStudent()  {
	fmt.Print(c)
}

func (s Student)printStudentDetails()  {
	fmt.Print(s)
}

func main() {
	c := city{Student{name:"ravi", rollNumber:10101}, School{schoolName:"SVCET", schoolNumber:909090}, "hyd"}
	fmt.Println(c)
	c.printStudent()
}