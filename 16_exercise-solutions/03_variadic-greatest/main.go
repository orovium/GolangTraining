package main

import "fmt"

const theBest int = 42

func main() {
	n, err := greatest(51, 43, 44, 41, 65)
	if err {
		println("Best number not found")
	} else {
		fmt.Println("Greatest number is: ", n)
	}
}

func greatest(numbers ...int) (int, bool) {
	var bestFound bool
	for _, v := range numbers {
		if v == theBest {
			bestFound = true
			break
		}
	}
	return theBest, !bestFound
}
