package main

import "fmt"

var g string = "cowboy"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", g)
}
