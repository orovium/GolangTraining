package main

import (
	"fmt"
	"sort"
)

func main() {

	// 1
	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)

	// 2
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	// sort.SrtingSlice(s).Sort()
	fmt.Println(s)

	// 3
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)

	// Reverse
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println(studyGroup)
	fmt.Println(s)
	fmt.Println(n)
}
