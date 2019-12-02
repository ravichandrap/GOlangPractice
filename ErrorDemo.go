package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42.")
	}
	return arg + 2, nil
}

type argError struct {
	arg  int
	prob string
}

func (r *argError) Error() string {
	return fmt.Sprintf("%d - %s", r.arg, r.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with 42"}
	}
	return arg + 20202, nil
}

func main() {
	//for _, i := range []int{7,42} {
	//	if r,e:=f1(i) ;e !=nil {
	//		fmt.Println("f1 failed ", r)
	//		fmt.Println(e)
	//	} else {
	//		fmt.Println("f2 success")
	//	}
	//}

	for _, i := range []int{42} {
		_, e := f2(i);
		if r, ok := e.(*argError); ok {
			fmt.Println("f2 can't works", r)
			fmt.Println(ok)
		} else {
			fmt.Println("f2 works well ", r)
			fmt.Println(ok)
		}
	}
}
