package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)

	a, b := half(number)
	println(a, b)
}

func half(x int) (int, bool) {
	even := false
	if x%2 == 0 {
		even = true
	}
	return x / 2, even
}
