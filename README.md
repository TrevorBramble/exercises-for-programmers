# Exercises for Programmers
Solutions for [Exercises for
Programmers](https://pragprog.com/book/bhwb/exercises-for-programmers), by
[Brian Hogan](//github.com/napcs)

## Plan

I'm going to implement the solution to each exercise in a handful of languages.
Largely I hope to cover the exercises with Ruby and Go, and later Haskell,
Python, Elm, Rust, or C.

Each exercise has a few "challenges" appended, elaborating on the original
problem. I'm going to ignore this on the first run through, then revisit them
subsequently.

## Organization

Each exercise will have its own self-contained sub-directory within this
repository including a README to document the problem and related details.

Within each exercise directory will be sub-directories enclosing solutions, one
per language.

## Running the Solutions

How to execute the implemented solutions will depend on each language, among
other things.

### Ruby

All solutions were written with Ruby 2.2.3 (MRI/CRuby). With a compatible Ruby
installed they should all be executable in the usual fashion, by passing the
file name to the interpreter.

```
$ ruby solution.rb
```
Additionally, all of the Ruby solutions have a "shebang" (`#!`) specified, so
you may be able to execute them directly, depending.

```
$ ./solution.rb
```
### Go

All solutions in Go are written against 1.5.2. They can be built or run as-is
in the given directory, if you have the Go platform installed.

```
$ go build main.go
$ ./main
```
```
$ go run main.go
```


## The Exercises

 1. [Saying Hello](01-saying-hello/)
 2. [Counting the Number of Characters](02-counting-the-number-of-characters/)
 3. [Printing Quotes](03-printing-quotes/)
 4. [Mad Lib](04-mad-lib/)
 5. [Simple Math](05-simple-math/)
 6. [Retirement Calculator](06-retirement-calculator/)
 7. Area of a Rectangular Room
 8. Pizza Party
 9. Paint Calculator
 10. Self-Checkout
 11. Currency Conversion
 12. Computing Simple Interest
 13. Determining Compound Interest
 14. Tax Calculator
 15. Password Validation
 16. Legal Driving Age
 17. Blood Alcohol Calculator
 18. Temperature Converter
 19. BMI Calculator
 20. Multistate Sales Tax Calculator
 21. Numbers to Names
 22. Comparing Numbers
 23. Troubleshooting Car Issues
 24. Anagram Checker
 25. Password Strength Indicator
 26. Months to Pay Off a Credit Card
 27. Validating Inputs
 28. Adding Numbers
 29. Handling Bad Input
 30. Multiplication Table
 31. Karvonen Heart Rate
 32. Guess the Number Game
 33. Magic 8 Ball
 34. Employee List Removal
 35. Picking a Winner
 36. Computing Statistics
 37. Password Generator
 38. Filtering Values
 39. Sorting Records
 40. Filtering Records
 41. Name Sorter
 42. Parsing a Data File
 43. Website Generator
 44. Product Search
 45. Word Finder
 46. Word Frequency Finder
 47. Who’s in Space?
 48. Grabbing the Weather
 49. Flickr Photo Search
 50. Movie Recommendations
 51. Pushing Notes to Firebase
 52. Creating Your Own Time Service
 53. Todo List
 54. URL Shortener
 55. Text Sharing
 56. Tracking Inventory
 57. Trivia App
