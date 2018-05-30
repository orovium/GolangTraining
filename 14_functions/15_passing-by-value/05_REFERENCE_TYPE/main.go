package main

import "fmt"

func main() {
	m := make([]string, 1, 25) // m is an slice
	fmt.Println(m)             // []
	changeMe(m)
	fmt.Println(m) // [TRUMAN]
}

func changeMe(z []string) {
	z[0] = "TRUMAN"
	fmt.Println(z) // [TRUMAN]
}
