package main

import "fmt"

type tank interface {
	Tarea() float64
	Volume() float64
	Print()
}

type myValue struct {
	radius float64
	height float64
	valaue string
}

func  NewyValueInstance(r float64,h float64, v string)   *myValue {
	return &(myValue{radius: r,height:h,valaue:v})
}

func (m *myValue) Tarea() float64  {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m *myValue) Volume() float64 {
	return 3.14 * m.radius* m.radius* m.height
}

func (m *myValue) outPut()  {

	fmt.Println("Area of tank:", m.Tarea())
	fmt.Println("Volume of tank:", m.Volume())
	fmt.Println("Value of tank:", m.valaue)
}

func (m myValue)modifyValue()  {
	m.valaue = "changed "
}

func main() {
	//var t tank
	//t = myValue{10,14}
	t := myValue{10,14, "one"}
	//t.Print()

	a := make([]myValue,10,10)
	a[0] = t
	a[1] = myValue{20,24, "one"}
	a[2] = myValue{20,24, "one"}
	a[3] = myValue{20,24, "one"}

	a[2].outPut()
	a[2].modifyValue()
	a[2].outPut()
	//fmt.Print(a[2].Volume())
	vs :=  NewyValueInstance(10,20,"NEW VALUE")

	fmt.Print(vs.valaue )
}
