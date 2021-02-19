package main

import (
	"embed"
	"fmt"
)

//go:embed message.txt
var f embed.FS

func main() {
	message, _ := f.ReadFile("message.txt")
	fmt.Println(string(message))
}
