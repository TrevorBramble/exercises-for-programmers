#!/usr/bin/env ruby

PIECES_PER_PIZZA = 8

print "How many people? "
people = gets.chomp.to_i

print "How many pizzas do you have? "
pizzas = gets.chomp.to_i

pieces = pizzas * PIECES_PER_PIZZA
piecesPerPerson = pieces / people
allocatedPieces = piecesPerPerson * people
leftoverPieces = pieces - allocatedPieces

puts "#{people} people with #{pizzas} pizzas"
puts "Each person gets #{piecesPerPerson} pieces of pizza"
puts "There are #{leftoverPieces} leftover pieces."
