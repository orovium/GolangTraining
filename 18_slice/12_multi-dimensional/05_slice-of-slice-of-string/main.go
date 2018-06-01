package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Truman"
	student1[1] = "Capote"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// strore the record
	records = append(records, student1)
	//student 2
	student2 := make([]string, 4)
	student2[0] = "Sandalio"
	student2[1] = "Bocachancla"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// strore the record
	records = append(records, student2)
	// print
	fmt.Println(records)
}
