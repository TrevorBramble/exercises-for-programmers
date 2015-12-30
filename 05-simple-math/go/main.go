package main

import "fmt"

func main() {
	var firstNumber, secondNumber int

	fmt.Print("What is the first number? ")
	fmt.Scanln(&firstNumber)

	fmt.Print("What is the second number? ")
	fmt.Scanln(&secondNumber)

	sum := firstNumber + secondNumber
	difference := firstNumber - secondNumber
	product := firstNumber * secondNumber
	quotient := firstNumber / secondNumber

	fmt.Printf("%d + %d = %d\n", firstNumber, secondNumber, sum)
	fmt.Printf("%d - %d = %d\n", firstNumber, secondNumber, difference)
	fmt.Printf("%d * %d = %d\n", firstNumber, secondNumber, product)
	fmt.Printf("%d / %d = %d\n", firstNumber, secondNumber, quotient)
}
