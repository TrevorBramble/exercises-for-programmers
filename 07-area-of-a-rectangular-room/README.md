# Area of a Rectangular Room

 * [x] [Ruby](ruby/)
 * [x] [Go](go/)


## Problem Statement

Create a program that calculates the area of a room. Prompt the user for the
length and width of the room in feet. Then display the area in both square feet
and square meters.


## Example Output

```
What is the length of the room in feet? 15
What is the width of the room in feet? 20
You entered dimensions of 15 feet by 20 feet.
The area is
300 square feet
27.871 square meters
```


## Constraints

 * Keep the calculations separate from the output.
 * Use a constant to hold the conversion factor.


## Analysis


### Inputs: 

 * length
 * width


### Processes:

 * calculate area in feet
 * convert area from feet to meters


### Outputs:

 * area in feet
 * area in meters


## Test Plan

Input:

Expected result:


## Pseudocode

```
RectangularRoomArea
  Initialize feetToMetersFactor to 0.09290304

  Prompt for length with "What is the length of the room in feet?"
  Prompt for width with "What is the width of the room in feet?"

  areaInFeet = length * width
  areaInMeters = areaInFeet * feetToMetersFactor

  Display "You entered dimensions of {length} feet by {width} feet."
  Display "The area is"
  Display "{areaInFeet} square feet"
  Display "{areaInMeters} square meters"
End
```
