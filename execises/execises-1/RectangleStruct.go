package main

import "fmt"

type rectangle struct {
	length   int
	breadth  int
	color    string
	geometry struct {
		area      int
		perimeter int
	}
}

func (r rectangle) PrintDetails() {
	fmt.Println("The following are the student details")
	fmt.Println("length:", r.length)
	fmt.Println("breadth:", r.breadth)
	fmt.Println("breadth:", r.breadth)
	fmt.Println("Geometry:", r.geometry)
}

type MyRectangel struct {
	length  int
	breadth int
	color   string
}

func footMyRect() {
	var myRec MyRectangel
	//myRec = MyRectangel{length:20, color:"Yellow", breadth:30}
	//myRec = MyRectangel{length:20,  breadth:30}
	myRec = MyRectangel{20, 30, "Yellow"}
	fmt.Println(myRec)
}

func footRect() {
	var rect rectangle
	rect.breadth = 20
	rect.length = 10
	rect.color = "Red"
	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.breadth + rect.length)

	fmt.Println(rect)
	fmt.Println(rect.geometry)
	fmt.Println("Area: \t", rect.geometry.area)
	fmt.Println("Perimeter: ", rect.geometry.perimeter)

	rectNew := new(rectangle)
	rectNew.breadth = 2020
	rectNew.length = 1000

	fmt.Println(rectNew)
	fmt.Println(&rectNew)
	fmt.Println(*rectNew)
}

func main() {
	//footRect()
	//footMyRect()
	//fmt.Println(rectangle{10.5,20.5, "Red"})
	methodInStruct()
	fooTestMethod()
}

func fooTestMethod() {

}

func (r rectangle) Area() int {
	return r.breadth * r.length
}
func (r *rectangle) Area_by_reference() int {
	return r.breadth * r.length
}
func methodInStruct() {
	var s rectangle;
	s.length = 20
	s.breadth = 30
	s.geometry.perimeter = 10
	s.geometry.area = 9

	s.PrintDetails()
	area := s.Area();
	fmt.Println(area)
	fmt.Println(s.Area_by_reference())
}
