# Pizza Party

 * [x] [Ruby](ruby/)
 * [x] [Go](go/)


## Problem Statement

Write a program to evenly divide pizzas. Prompt for the number of people, the
number of pizzas, and the number of slices per pizza. Ensure that the number of
pieces of pizza each person should get. If there are leftovers, show the number
of leftover pieces.


## Example Output

```
How many people? 8
How many pizzas do you have? 2
8 people with 2 pizzas
Each person gets 2 pieces of pizza.
There are 0 leftover pieces.
```


## Analysis


### Inputs: 

 * people
 * pizzas


### Processes:

 * divide pizzas by people


### Outputs:

 * pieces per person
 * leftover pieces


## Test Plan

Input:
 * people: 8
 * pizzas: 2

Expected result:
 * pieces per person: 2
 * leftover pieces: 0


## Pseudocode

```
PizzaParty
  Initialize piecesPerPizza to 8

  Prompt for people with "How many people?"
  Prompt for pizzas with "How many pizzas do you have?"

  pieces = pizzas * piecesPerPizza
  piecesPerPerson = pieces / people
  allocatedPieces = piecesPerPerson * people
  leftoverPieces = pieces - allocatedPieces

  Display "{people} people with {pizzas} pizzas"
  Display "Each person gets {piecesPerPerson} pieces of pizza"
  Display "There are {leftoverPieces} leftover pieces."
End
```
