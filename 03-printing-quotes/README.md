# Exercise Name

 * [x] [Ruby](ruby/)
 * [x] [Go](go/) (w/ caveat)


## Problem Statement

Create a program that prompts for a quote and an author.
Display the quotation and author as shown in the example
output.

## Example Output

```
What is the quote? These aren't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids
you're looking for."
```


## Constraints

 * Use a single output statement to produce this output,
using appropriate string-escaping techniques for quotes.

 * If your language supports string interpolation or string
substitution, don’t use it for this exercise. Use string
concatenation instead.

## Analysis


### Inputs: 

 * quote
 * author


### Processes:

 * prompt for quote and author
 * display quotation with author


### Outputs:

 * quote
 * author


## Test Plan

Input:
 * quote: These aren't the droids you're looking for.
 * author: Obi-Wan Kenobi

Expected result:
 * Obi-Wan Kenobi says, "These aren't the droids you're looking for." 


## Pseudocode

```
PrintQuote
  Prompt for quote with "What is the quote?"
  Prompt for author with "Who said it?"

  Display author + " says, "" + quote + """
End
```
