package main

import (
	"fmt"
	"sync"
	"time"
)


var wg sync.WaitGroup;
func printMe(s string) {
	for i:=0;i<3;i++ {
		fmt.Println( s,i )
		time.Sleep(time.Millisecond*100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go printMe("ravi")
	wg.Add(1)
	go printMe("Shobha")
	wg.Wait()
}

