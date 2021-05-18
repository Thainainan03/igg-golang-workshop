package main

import "fmt"

type contact struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact
}

func main() {

	jame := person{
		firstName: "Thanainan",
		lastName:  "Khamthan",
		contact: contact{
			email: "khamthan02@gmail.com",
			zip:   55110,
		},
	}
	jame.print()
}

func (p person) print() {
	fmt.Println(p)
}
