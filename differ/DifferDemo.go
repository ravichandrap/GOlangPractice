package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.example.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		res.Body.Close()
	}()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
