package main

import (
	"fmt"

	"github.com/orovium/GolangTraining/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse(("!oG ,olleH")))
	fmt.Println(stringutil.Reverse(stringutil.MyName))
}
