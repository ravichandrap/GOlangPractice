package main

import (
	"fmt"
	"sort"
)

func mapInterface() {
	myMap := make(map[string]interface{})
	myMap["one"] = 1
	myMap["two"] = "two"

	fmt.Println(myMap)
}

func mapSort() {
	unSortedMap := map[string]int{"india": 91, "canada": 70, "Germany": 15}
	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		fmt.Println(k)
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}
}

func merge() {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4, "f": 30}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)
}

func printDefault() {
	base := Base{}
	fmt.Print(base.APIVersion)
}

func variadicFunc(i interface{}) {

}

// The Base struct aggregates attributes that are part of all CoRE requests.
type Base struct {
	APIVersion       string `json:"apiVersion"`
	ApplicationID    string `json:"applicationId"`
	DataStatus       string `json:"dataStatus"`
	Destination      string `json:"destination"`
	CarLineID        string `json:"carLineId"`
	ISet             string `json:"iSet"`
	XSet             string `json:"xSet"`
	Importer         bool   `json:"importer"`
	ConfiguratorDate int    `json:"configuratorDate"`
	Language         string `json:"language"`
	PriceType        string `json:"priceType"`
	PriceGroup       string `json:"priceGroup"`
}

func main() {
	//mapInterface()
	//mapSort()
	//merge()
	printDefault()

}
