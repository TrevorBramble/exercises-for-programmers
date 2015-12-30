package main

import "fmt"

func main() {
	const piecesPerPizza = 8

	var people, pizzas int

	fmt.Print("How many people? ")
	fmt.Scanln(&people)

	fmt.Print("How many pizzas do you have? ")
	fmt.Scanln(&pizzas)

	pieces := pizzas * piecesPerPizza
	piecesPerPerson := pieces / people
	allocatedPieces := piecesPerPerson * people
	leftoverPieces := pieces - allocatedPieces

	fmt.Printf("%d people with %d pizzas\n", people, pizzas)
	fmt.Printf("Each person gets %d pieces of pizza\n", piecesPerPerson)
	fmt.Printf("There are %d leftover pieces.\n", leftoverPieces)
}
