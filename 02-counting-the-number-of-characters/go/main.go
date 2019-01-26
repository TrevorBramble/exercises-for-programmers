package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the input string? ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(0)
	}

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Missing input")
		os.Exit(0)
	}

	fmt.Printf("%s has %d characters.\n", input, len(input))
}
