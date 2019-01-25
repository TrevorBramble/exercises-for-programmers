#!/usr/bin/env ruby

print "What is your current age? "
current_age = gets.chomp.to_i

print "At what age would you like to retire? "
retirement_age = gets.chomp.to_i

years_remaining = retirement_age - current_age
current_year    = Time.now.year
retirement_year = current_year + years_remaining

puts "You have #{years_remaining} years left until you can retire."
puts "It's #{current_year}, so you can retire in #{retirement_year}."
