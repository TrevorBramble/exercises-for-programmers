package main

import (
	"fmt"
	"os"
)

func main() {
	num1 := getNumber("What is the first number? ")
	num2 := getNumber("What is the second number? ")

	fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
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
