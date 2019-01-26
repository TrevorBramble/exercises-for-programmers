package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf(
		"Do you %s your %s %s %s? That's hilarious!\n",
		getInput("Enter a noun: "),
		getInput("Enter a verb: "),
		getInput("Enter an adjective: "),
		getInput("Enter an adverb: "),
	)
}

func getInput(prompt string) (input string) {
	fmt.Print(prompt)

	fmt.Scanln(&input)

	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Missing input")
		os.Exit(0)
	}

	return
}
