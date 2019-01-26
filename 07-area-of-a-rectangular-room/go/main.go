package main

import (
	"fmt"
	"os"
)

func main() {
	const FeetToMetersFactor = 0.09290304

	length := getNumber("What is the length of the room in feet? ")
	width := getNumber("What is the width of the room in feet? ")

	areaInFeet := length * width
	areaInMeters := float64(areaInFeet) * FeetToMetersFactor

	fmt.Printf("You entered dimensions of %d feet by %d feet.\n",
		length, width)
	fmt.Printf("The area is\n%d square feet\n%.3f square meters\n",
		areaInFeet, areaInMeters)
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
