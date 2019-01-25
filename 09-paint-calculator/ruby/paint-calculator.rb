#!/usr/bin/env ruby

GALLON = 350

print "What is the length of the room? "
length = gets.chomp.to_i

print "What is the width of the room? "
width = gets.chomp.to_i

area    = length * width
gallons = area < GALLON ? 1 : (area.to_f / GALLON).ceil

puts "You will need to purchase #{gallons} gallons of"
puts "paint to cover #{area} square feet."
