package main

import (
	_ "embed"
	"fmt"
)

//go:embed file.txt
var s string

func main() {
	//go:embed file.txt
	var s1 string
	fmt.Print(s1)
	// fmt.Print(s)
}
