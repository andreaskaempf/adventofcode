# Advent of Code 2019

My solutions for the Advent of Code 2024, completed in later years.
see https://adventofcode.com/2019

* **Day 1** (Go): Create a simple function to convert "mass" to "fuel"
  required (both integers). For Part 1, add up the fuel obtained from
  a list of mass values. For Part 2, also consider the fuel required for
  the fuel, iterating.

* **Day 2** (Go): Simulate an assembly language-like CPU, with a list of 
  numbers that represent opcodes (just add or multiply) followed by arguments
  that are input and output locations in memory. For Part 1, change two values
  in memory to fixed values, run the program, and report final value in
  location zero. For Part 2, we find the combination that results in a certain
  output value (just used brute force).

* **Day 3** (Go): Give two lists of instructions for two "wires", with each
  instruction L/R/U/D followed by a number, traverse a 2D grid along the
  specified directions, to find where the two wires cross (but don't include a
  wire crossing with itself). For Part 1, find the closest intersection
  from the starting point. For Part 2, find the intersection that had the 
  shortest combined number of steps to get there.

