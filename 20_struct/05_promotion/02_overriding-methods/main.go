package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm Just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("I'm Truman Capote")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  42,
	}

	p2 := doubleZero{
		person: person{
			Name: "Truman Capote",
			Age:  32,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
