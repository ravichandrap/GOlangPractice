package main

import (
	"fmt"
)

type data int

func main() {
	v1 := data(20)
	v2 := data(3)
	result := v1.mul(v2)
	fmt.Println(result)
	//------
	res := author{
		name:      "Ravichadnra",
		branch:    "cse",
		particles: 0,
	}
	fmt.Println(res)
	flag := res.setParticles(2)
	fmt.Println(flag)
	fmt.Println(res)
}

func (d1 data) mul(d2 data) data {
	return d1*d2
}

//---------------
type author struct {
	name string
	branch string
	particles int
}

func (a *author) setParticles(p int) bool {
	if p <= 0 {
		return false
	} else {
		(*a).particles = p
		return true
	}

}







