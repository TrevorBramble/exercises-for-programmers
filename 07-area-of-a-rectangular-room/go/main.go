package main

import "fmt"

func main() {
	const FeetToMetersFactor = 0.09290304

	var length, width int

	fmt.Print("What is the length of the room in feet? ")
	fmt.Scanln(&length)

	fmt.Print("What is the width of the room in feet? ")
	fmt.Scanln(&width)

	areaInFeet := length * width
	areaInMeters := float64(areaInFeet) * FeetToMetersFactor

	fmt.Printf("You entered dimensions of %d feet by %d feet.\n",
		length, width)
	fmt.Printf("The area is\n%d square feet\n%.3f square meters\n",
		areaInFeet, areaInMeters)
}
