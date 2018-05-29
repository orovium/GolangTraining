package main

import "fmt"

const meters2Yards float64 = 1.09361

func main() {
	var meters float64

	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * meters2Yards
	fmt.Println(meters, " meters is", yards, " yards.")
}
