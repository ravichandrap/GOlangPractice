package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("hello.html"))
}

func main() {
	fmt.Println("Templates")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	pathname := r.URL.Path
	names := strings.Split(pathname, "&")
	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
	}

	name := Name{
		names[0],
		strings.ToUpper(names[1]),
		[]string{"Shobha", "Divija", "Sarala", "Bhaskar"},
		sages,
	}
	fmt.Println(name)
	temp.Execute(w, name)
}

type Name struct {
	Fname, Lname string
	Family       []string
	Sages        map[string]string
}
