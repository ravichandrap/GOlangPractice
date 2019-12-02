package main

import "fmt"

type address struct {
	name string
	houseNumber int
	street string
	city string
	country string
}

func (a *address)show()  {
	fmt.Println((*a).name)
	fmt.Println((*a).houseNumber)
	fmt.Println((*a).street)
	fmt.Println((*a).city)
	fmt.Println((*a).country)
}

func main() {
	a := &address{
		name:        "Ravichandra",
		houseNumber: 7,
		street:      "street one",
		city:        "city BW",
		country:     "Germany",
	}

	fmt.Println(a.country)
	a.show()
}
