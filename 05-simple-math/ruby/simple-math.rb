#!/usr/bin/env ruby

print "What is the first number? "

first_number = gets.chomp.to_i

print "What is the second number? "

second_number = gets.chomp.to_i

sum        = first_number + second_number
difference = first_number - second_number
product    = first_number * second_number
quotient   = first_number / second_number

puts <<DOC
#{first_number} + #{second_number} = #{sum}
#{first_number} - #{second_number} = #{difference}
#{first_number} * #{second_number} = #{product}
#{first_number} / #{second_number} = #{quotient}
DOC
