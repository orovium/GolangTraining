package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello Truman")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
