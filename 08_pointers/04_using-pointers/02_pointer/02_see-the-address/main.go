package main

import "fmt"

func zero(x *int) {
	fmt.Println(x) // address in func zero
	*x = 0
}

func main() {
	x := 5
	fmt.Println(&x) // address in main
	zero(&x)
	fmt.Println(x) // x is still 5
}
