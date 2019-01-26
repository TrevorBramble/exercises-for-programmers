package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	currentAge := getNumber("What is your current age? ")
	retirementAge := getNumber("At what age would you like to retire? ")

	currentYear := time.Now().Year()

	yearsRemaining := retirementAge - currentAge
	retirementYear := currentYear + yearsRemaining

	fmt.Printf("You have %d years left until you can retire.\n", yearsRemaining)
	fmt.Printf("It's %d, so you can retire in %d.\n", currentYear, retirementYear)
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
