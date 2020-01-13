package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"Wool", 17}
	temp, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}} ")
	if err != nil {
		panic(err)
	}
	err = temp.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
	fmt.Println(temp.ParseName)
}
