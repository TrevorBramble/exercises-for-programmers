# Simple Math

 * [x] [Ruby](ruby/)
 * [x] [Go](go/)


## Problem Statement

Write a program that prompts for two numbers. Print the sum, difference,
product, and quotient of those numbers as shown in the example output:


## Example Output

```
What is the first number? 10
What is the second number? 5
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2
```


## Constraints

 * Values coming from users will be strings. Ensure that you convert these
   values to numbers before doing the math.
 * Keep the inputs and outputs separate from the numerical conversions and
   other processing.
 * Generate a single output statement with line breaks in the appropriate
   spots.


## Analysis


### Inputs: 

 * first number
 * second number


### Processes:

 * prompt for two numbers
 * find their sum
 * find their difference
 * find their product
 * find their quotient
 * print results


### Outputs:

 * sum
 * difference
 * product
 * quotient


## Test Plan

Input:
 * first number: 10
 * second number: 5

Expected result:
 * sum: 15
 * difference: 5
 * product: 50
 * quotient: 2


## Pseudocode

```
SimpleMath
  Prompt for firstNumber with "What is the first number?"
  Prompt for secondNumber with "What is the second number?"

  sum = firstNumber + secondNumber
  difference = firstNumber - secondNumber
  product = firstNumber * secondNumber
  quotient = firstNumber / secondNumber

  Display "{firstNumber} + {secondNumber} = {sum}
           {firstNumber} - {secondNumber} = {difference}
           {firstNumber} * {secondNumber} = {product}
           {firstNumber} / {secondNumber} = {quotient}"
End
```
