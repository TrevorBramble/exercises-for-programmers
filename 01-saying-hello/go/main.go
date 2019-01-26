package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name? ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(0)
	}

	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Missing input")
		os.Exit(0)
	}

	fmt.Printf("Hello, %s, nice to meet you!\n", name)
}
