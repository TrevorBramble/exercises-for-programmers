package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	quote := getInput("What is the quote? ", reader)
	author := getInput("Who said it? ", reader)

	fmt.Printf("%s says, \"%s\"\n", author, quote)
}

func getInput(prompt string, reader *bufio.Reader) (input string) {
	fmt.Print(prompt)

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

	return
}
