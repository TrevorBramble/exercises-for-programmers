package main

import "fmt"

/*
 * This code is currently broken due to a bug or bugs in Go's fmt package.
 * For more details: https://github.com/golang/go/commit/dbaf5010b33e5819050ebb0a387eb0bff2cfb8bf
 *
 * There are other ways to do this and I might come back to do them, but I'm
 * too irked at the moment to bother.
 */
func main() {
	var quote, author string

	fmt.Print("What is the quote? ")

	fmt.Scanln(&quote)

	fmt.Print("Who said it? ")

	fmt.Scanln(&author)

	fmt.Sprintf("%s says, \"%s\"\n", author, quote)
}
