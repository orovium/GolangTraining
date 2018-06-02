package main

import "fmt"

func main() {

	// var myGreeting map[string]string // This will not work
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
