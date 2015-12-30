package main

import "fmt"

func main() {
	var str string

	fmt.Print("What is the input string? ")

	fmt.Scanln(&str)

	fmt.Printf("%s has %d characters.\n", str, len(str))
}
