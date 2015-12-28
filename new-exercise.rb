#!/usr/bin/env ruby

#TODO: modify new README to reflect exercise name
#TODO: modify project README to link new exercise

puts "Exercise number?"
number = gets.chomp

number = "0" + number if number.length == 1

puts "Exercise name?"
name = gets.chomp

name.downcase.gsub!(' ', '-')

directory = "./" + number + "-" + name

Dir.mkdir directory

IO.copy_stream "./WORKSHEET.md", directory + "/README.md"

Dir.mkdir directory + "/ruby"
