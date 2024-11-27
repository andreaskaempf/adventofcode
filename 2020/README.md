# Advent of Code 2020

My solutions for the Advent of Code 2020 (done mostly in 2022),
see https://adventofcode.com/2020

* **Day 1** (Python): Read through a list of numbers, and find two
  (then three) that add up to 2020 (*easy*)

* **Day 2** (Python): Read each line, and validate password against 
  character-count rules (*easy*)

* **Day 3** (Python): Count up the number of '#' encountered when traversing
  an array with a given dx/dy slope for each step. The rows wrap if x 
  coordinate is wider than the data (*easy*).

* **Day 4** (Python): Simple validation of semi-structured text "passports",
  first looking for missing fields, then checking field contents against some
  simple rules (*easy*).

* **Day 5** (Python): Apply a sequence of binary forward/back operators to find
  row and seat in an airplane, and compute a code for these. For Part 1, find
  the maximum seat ID, for Part 2 the missing one.

* **Day 6** (Python): Read a file containing a sequence of letters on each row,
  each letter being a customer's "yes" response to a survey question. Customers
  are in groups, separated by newlines.  Part 1: For each group, count the
  number of questions to which anyone answered "yes". What is the sum of those
  counts?  Part 2: For each group, count the number of questions to which
  everyone answered "yes". What is the sum of those counts? (*easy*)

* **Day 7** (Python): Recursively navigate a tree upwards (Part 1) and
  downwards (Part 2), to count how many bags contain a given type of bag, and
  how many bags a given type of bag ultimately contains. (*easy*)

* **Day 8** (Python): Simulate execute of some simple bytecode, and detect when
  an instruction is executed the second time. For Part 2, find which nop or jmp
  needs to be changed to the opposite instruction, in order to fix the infinite
  loop. (*easy*)

* **Day 9** (Julia): Given a list of numbers, and a given sliding window width, 
  find the first number that cannot be the sum of two digits in the preceding
  window. For Part 2, find any set of at least two contiguous numbers that add
  up to the number not found above, and return the sum of the minimum and
  maximum of this range (*easy*).

* **Day 10** (Python): Given a sequence of numbers, count up the number of 1-
  and 3- gaps within the list. Then, count up all the possible paths between
  the lowest and highest values, enforcing the constraint that the gap between
  any two number cannot be more than 3. Exhaustive recursive approach was too
  slow, but extremely fast if you cache interim results (7 ms in Python, *medium*).

* **Day 11** (Go): Change the state of a "seating plan" depending on 
  available seats immediately adjacent to every seat (Part 1), or visible 
  in any direction (Part 2). *Easy*, but took a long time.

* **Day 12** (Go): Simulate movement of a "ship" based on simple 
  instructions, directly for Part 1, relative to a "waypoint" for Part 2
  (should have been easy, but instructions were hard to understand, 
  so *medium*).

* **Day 13** (Go): Solve problems related to a schdule of bus times, where
  all buses leave at t = 0, but take different number of minutes to reach 
  the station. For Part 1, find the earliest bus that will arrive. 
  For Part 2, find the earliest time at which the first bus arrives at t, 
  the second at t+1, and so on. Used a brute force solution (takes 1 hr 10 
  mins), but there must be a better way. This solution was quite easy, but 
  I'm marking this problem as *hard* since I spent a lot of time trying 
  unsuccessfully to come up with an algorithm for Part 2 that would find 
  the solution directly, rather than using brute force.

* **Day 14** (Go): Read a "program" consisting of binary masks and instructions
  to set memory at given address to a value. For part 1, apply the mask to the
  value. For Part 2, first apply the mask to the address, then expand the
  address to all possible permutations where 'X' are changed to 1 and 0. For
  both parts, sum up the values in memory to get the answer. *Medium*, because
  of convoluted instructions and manipulation/explosion of bit masks.

* **Day 15** (Go): Simulate a convoluted memory game, which becomes infeasible
  using a simple list of history, when the number of iterations goes from 
  2020 (part 1) to 30 million (part 2).

* **Day 16** (Go): Read a file containing train ticket field names, and data
  for my ticket and a bunch of other tickets. In Part 1, identify and remove
  tickets that are invalid, because they do not match the allowed ranges for
  any field.  In Part 2, infer which columns relate to which fields, and report
  the value of "departure" fields for my ticket. *Hard*

* **Day 17** (Go): Input is a set of "cubes" in 2-d space, either on or off.
  For part 1, this is extended to 3-d space, for part 2 4-d space. Simulate a
  set of simple rules, depending on current state of a cube and the number of
  "on" neighbours it has. Simulation is supposed to occur "simulataneously", so
  apply changes to future state, then roll them to current state after each
  iteration. Part 2 is a trivial set of modifications to Part 1, to make it 4-d
  instead of 3-d, so only the Part 2 code is retained here. *Medium*

* **Day 18** (Go): Parse and evaluate four-function arithmetic expressions with
  parentheses, with left-to right evaluation (no operator precedence) for Part
  1, and mult/div having higher precedence for Part 2. Implemented Djikstra's
  Shunting Yard Algorithm. Part 2 was a trivial change to some precedence
  weights. *Hard*

* **Day 19** (Go): Recursively find if character pattern matches a set of
  recursive pattern rules; solved by converting rules to a large regular 
  expression, recursive for Part 2. *Hard*

* **Day 20** (Go): Assemble a set of "tiles" into an image, so adjacent edges
  match, flipping or rotating as necessary. Part 1 is the product of the IDs of
  the corner tiles.  For Part 2, strip the edges of the tiles and assemble them 
  into a combined image, then search for a 3-line "sea monster" pattern 
  (flipping and rotating the combined image as required), and report the number 
  of hash marks in the image that are not covered up by the "sea monsters" 
  found. *Very hard*

* **Day 21** (Go): Read a list of ingredients and associated allergens, and 
  determine which ingredients do not produce any allergies (Part 1), and a 
  list of ingredients which produce allergies, sorted by allergen (Part 2). 
  *Quite easy* using set operations.

* **Day 22**: TO DO

* **Day 23**: TO DO

* **Day 24** (Go): Given black/white tiles on a hexoganal grid, follow set of
  movement directions and flip over tiles, then count number of black tiles.
  For Part 2, simulate 100 days of flipping tiles based on state and number of
  adjacent black tiles. *medium*

* **Day 25**: TO DO

To compile and run a **Go** program
* Change into the directory with the program
* go build day02.go
* ./day02

To compile and run a **Rust** program
* Change into the directory with the program
* rustc day02.go
* ./day02

To run a **Julia** program
* Change into the directory with the program
* julia day02.jl

To run a **Python** program
* Change into the directory with the program
* python day06.py

AK, Dec 2021-2022
