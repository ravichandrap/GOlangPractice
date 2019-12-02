package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	 Name string `json:name`
}

func animalMarshell(animals []Animal) (string)  {
	anim, err := json.Marshal(animals)
	if err != nil {
		log.Fatal(err)
	}
	return string(anim)
}

func animalUnmarshel(animalJson string) ([]Animal)  {
	var animals []Animal;
	if err := json.Unmarshal([]byte(animalJson), & animals); err != nil {
		log.Fatal(err)
	}
	return animals
}
func main() {
	zoo := make([]Animal, 3, 5)
	var zooAnimals []Animal

	zoo[0] = Animal{Name:"Gopher"}
	zoo[1] = Animal{Name:"Zebra"}
	zoo[2] = Animal{Name:"Lion"}
	animals := animalMarshell(zoo)
	fmt.Println(animals)

	zooAnimals = animalUnmarshel(animals);
	fmt.Println(zooAnimals)

}