package main

import (
	"fmt"
	"net/url"
)

func main() {
	fooParseURL()
}

func fooParseURL() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err:= url.Parse(s)
	fmt.Println(u.User)
	fmt.Println(err)
}
