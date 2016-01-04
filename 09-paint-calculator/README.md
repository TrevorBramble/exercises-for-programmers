# Paint Calculator

 * [x] [Ruby](ruby/)


## Problem Statement

Calculate gallons of paint needed to paint the ceiling of a room. Prompt for
the length and width, and assume one gallon covers 350 square feet. Display the
number of gallons needed to poaint the ceiling as a whole number.


## Example Output

```
You will need to purchase 2 gallons of
paint to cover 360 square feet.
```


## Constraints

 * Use a constant to hold the conversion rate.
 * Ensure that you round *up* to the next whole number.


## Analysis


### Inputs: 

 * length
 * width


### Processes:

 * calculate area
 * calculate gallons needed


### Outputs:

 * area in feet
 * paint in gallons


## Test Plan

Input:
 * length
 * width

Expected result:
 * area in feet
 * gallons of paint


## Pseudocode

```
PaintCalculator
  Initialize gallon with 350

  Prompt for length with "What is the length of the room?"
  Prompt for width with "What is the width of the room?"

  area = length * width
  gallons = area / gallon, rounded up to the next whole number

  Display "You will need to purchase {gallons} gallons of"
  Display "paint to cover {area} square feet."
End
```
