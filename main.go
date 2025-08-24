package main

import (
	"fmt"
	"learn/just-go-fundamental/helpers"
)

func Hello(name string) string {
	return name
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(helpers.TextCapitalize("hello world"))
}
