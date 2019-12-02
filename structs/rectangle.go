package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t time.Time) MyTime() string {
	return t.Time.String()[0:5]
}

func main() {
	m := myTime{*time.LocalTime()}
	fmt.Println("Full time now:", m.String())      //calling existing String method on anonymous Time field
	fmt.Println("First 5 chars:", m.first5Chars()) //calling myTi
}
