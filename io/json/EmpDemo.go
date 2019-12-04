package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name string `json:name`
	Age int `json:age`
	Salary int `json:salary`
}

func jsonMarshall(empObj employee) {


	emp, err := json.Marshal(empObj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(emp))
}

func jsonUnMarshell(empJson string)  {
	bytes := []byte(empJson)
	var emp employee

	if err := json.Unmarshal(bytes, &emp); err != nil {
		log.Fatal(err)
	}

	fmt.Println(emp)
}
func main() {
	empObj := employee{Name:"Ravichandra", Age: 38, Salary:303030}
	empJson := `{"Name":"Ravichandra","Age":38,"Salary":303030}`
	jsonMarshall(empObj)
	jsonUnMarshell(empJson)
}