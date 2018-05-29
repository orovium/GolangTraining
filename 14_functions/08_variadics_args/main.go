package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // data is one item
	n := average(data...)                     // but average need
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
