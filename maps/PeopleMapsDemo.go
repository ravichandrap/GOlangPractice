package main

import "fmt"

type people struct {
	Name string
	Likes []string
}

func collectPeopleDetails(p *people) {

	var Likes = make([]string, 10,10)
	for i := range Likes {
		Likes[i] = fmt.Sprintf("%s:%d", p.Name,i)
	}
	fmt.Println(Likes)

	p.Name = "Ravichandra"
	p.Likes = Likes
}


func main() {
	p:=people{Name:"Ravichandra"}

	collectPeopleDetails(&p)

}


