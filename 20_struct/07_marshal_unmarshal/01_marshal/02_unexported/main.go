package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string // ** start with lowercase == not exported
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
	// ** So for this (not exported) Println print a empty struct
}
