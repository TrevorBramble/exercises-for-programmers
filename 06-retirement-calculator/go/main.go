package main

import (
	"fmt"
	"time"
)

func main() {
	var currentAge, retirementAge int

	fmt.Print("What is your current age? ")
	fmt.Scanln(&currentAge)

	fmt.Print("At what age would you like to retire? ")
	fmt.Scanln(&retirementAge)

	currentYear := time.Now().Year()

	yearsRemaining := retirementAge - currentAge
	retirementYear := currentYear + yearsRemaining

	fmt.Printf("You have %d years left until you can retire.\n", yearsRemaining)
	fmt.Printf("It's %d, so you can retire in %d.\n", currentYear, retirementYear)
}
