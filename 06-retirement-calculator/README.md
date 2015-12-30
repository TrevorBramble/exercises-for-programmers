# Exercise Name

 * [x] [Ruby](ruby/)


## Problem Statement

Create a program that determines how many years you have left until retirement
and the year you can retire. It should prompt for your current age and the age
you want to retire and display the output as shown in the example that follows.


## Example Output

```
What is your current age? 25
At what age would you like to retire? 65
You have 40 years left until you can retire.
It's 2015, so you can retire in 2055.
```


## Constraints

 * Be sure to convert the input to numerical data before doing any math.
 * Don't hard-code the current year into your program. Get it from the system
   time via your programming language.


## Analysis


### Inputs: 

 * current age
 * retirement age


### Processes:

 * calculate retirement year
 * calculate years remaining


### Outputs:

 * years remaining
 * retirement year


## Test Plan

Input:
 * current age: 25
 * retirement age: 65

Expected result:
 * years remaining: 40
 * retirement year: 2055


## Pseudocode

```
RetirementCalculator
  Prompt for currentAge with "What is your current age?"
  Prompt for retirementAge with "At what age would you like to retire?"

  yearsRemaining = retirementAge - currentAge
  retirementYear = currentYear + yearsRemaining

  Display "You have {yearsRemaining} years left until you can retire."
  Display "It's {currentYear}, so you can retire in {retirementYear}."
End
```
