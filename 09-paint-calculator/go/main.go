package main

import "fmt"
import "math"

func main() {
	const gallon = 350

	var length, width int

	fmt.Print("What is the length of the room? ")
	fmt.Scanln(&length)

	fmt.Print("What is the width of the room? ")
	fmt.Scanln(&width)

	area := length * width

	gallons := math.Ceil(float64(area) / gallon)

	fmt.Printf("You will need to purchase %d gallons of\n", int(gallons))
	fmt.Printf("paint to cover %d square feet.", area)
}
