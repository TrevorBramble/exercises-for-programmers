#!/usr/bin/env ruby

FEET_TO_METERS_FACTOR = 0.09290304

print "What is the length of the room in feet? "
length = gets.chomp.to_i

print "What is the width of the room in feet? "
width = gets.chomp.to_i

area_in_feet = length * width
area_in_meters = area_in_feet * FEET_TO_METERS_FACTOR

puts <<-DOC
You entered dimensions of #{length} feet by #{width} feet.
The area is
#{area_in_feet} square feet
#{area_in_meters.round(3)} square meters
  DOC
