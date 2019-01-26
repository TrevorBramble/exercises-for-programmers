package main

import (
	"fmt"
	"os"
)

func main() {
	const piecesPerPizza = 8

	people := getNumber("How many people? ")
	pizzas := getNumber("How many pizzas do you have? ")

	pieces := pizzas * piecesPerPizza
	piecesPerPerson := pieces / people
	allocatedPieces := piecesPerPerson * people
	leftoverPieces := pieces - allocatedPieces

	fmt.Printf("%d people with %d pizzas\n", people, pizzas)
	fmt.Printf("Each person gets %d pieces of pizza\n", piecesPerPerson)
	fmt.Printf("There are %d leftover pieces.\n", leftoverPieces)
}

func getNumber(prompt string) (number int) {
	fmt.Print(prompt)

	fmt.Scanln(&number)

	if number == 0 {
		fmt.Println("Let's not use zero.")
		os.Exit(0)
	}

	return
}
