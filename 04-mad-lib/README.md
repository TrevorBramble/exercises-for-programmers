# Mad Lib

 * [x] [Ruby](ruby/)
 * [x] [Go](go/)


## Problem Statement

Create a simple mad-lib program that prompts for a noun, a verb, an adverb, and
an adjective and injects those into a story that you create.

## Example Output

```
Enter a noun: dog
Enter a verb: walk
Enter an adjective: blue
Enter an adverb: quickly
Do you walk your blue dog quickly? That's hilarious!
```


## Constraints

 * Use a single output statement for this program.
 * If your language supports string interpolation or string substitution, use
   it to build up the output.


## Analysis


### Inputs: 

 * noun
 * verb
 * adjective
 * adverb


### Processes:

 * prompt for noun, verb, adjective, and adverb
 * inject inputs into story


### Outputs:

 * story


## Test Plan

Input:
 * noun: dog
 * verb: walk
 * adjective: blue
 * adverb: quickly

Expected result:
 * story: Do you walk your blue dog quickly? That's hilarious!


## Pseudocode

```
MadLib
  Prompt for noun with "Enter a noun:"
  Prompt for verb with "Enter a verb:"
  Prompt for adjective with "Enter an adjective:"
  Prompt for adverb with "Enter an adverb:"

  Display "Do you {verb} your {adjective} {noun} {adverb}? That's hilarious!"
End
```
