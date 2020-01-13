package main

import (
	"fmt"
	"runtime"
	"time"
)

func doSomething() {
	fmt.Println("first do some thing")
}

func doSomething2() {
	fmt.Println("second do some thing")
}

type toSchedule struct {
	myFunc   func()
	duration time.Duration
}

func main() {
	runtime.GOMAXPROCS(1)
	schdulePlan := []toSchedule{
		toSchedule{myFunc: doSomething, duration: time.Duration(1) * time.Second},
		toSchedule{myFunc: doSomething2, duration: time.Duration(2) * time.Second},
	}

	for _, s := range schdulePlan {
		go func(s toSchedule) {
			for {
				<-time.After(s.duration)
				s.myFunc()
			}
		}(s)
	}

	<-time.After(time.Duration(10) * time.Second)
}
