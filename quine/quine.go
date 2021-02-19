package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed quine.go
	src string
)

func main() {
	fmt.Println(src)
}
