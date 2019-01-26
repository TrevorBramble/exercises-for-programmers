package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	const gallon = 350

	length := getNumber("What is the length of the room? ")
	width := getNumber("What is the width of the room? ")

	area := length * width

	gallons := math.Ceil(float64(area) / gallon)

	fmt.Printf("You will need to purchase %d gallons of\n", int(gallons))
	fmt.Printf("paint to cover %d square feet.\n", area)
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
