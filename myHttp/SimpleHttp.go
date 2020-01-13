package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		resp.Body.Close()
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
