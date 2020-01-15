package main

import (
	"os"
	"strings"
	"text/template"
)

var temp *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	val := strings.TrimSpace(s)
	val = val[:3]
	return val
}

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseFiles("hello.html"))
}

func main() {
	n := Name{"Ravichandra", "Reddy"}
	temp.ExecuteTemplate(os.Stdout, "hello.html", n)
}

type Name struct {
	Fname string
	Lname string
}
