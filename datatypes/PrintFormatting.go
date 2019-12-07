package main

import "fmt"

func main() {
	var (
		a = 659
		b = false
		c = 2.309
		d = 4 + 1i
		e = "India"
		f = 15.2 * 4249.409
	)
	fmt.Printf("d for Integer: %d\n", a)
	fmt.Printf("6d for Integer: %6d\n", a)
	fmt.Printf("t for Boolean: %t\n", b)
	fmt.Printf("g for Float: %g\n", c)
	fmt.Printf("E for Float: %E\n", d)
	fmt.Printf("s for String: %s\n", e)
	fmt.Printf("g for Float: %g\n", f)

}
