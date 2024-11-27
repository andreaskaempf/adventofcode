# Advent of Code 2022

My solutions for the Advent of Code 2022, 
see https://adventofcode.com/2022
Line counts exclude blank lines and comments, and utility functions in utils.go

* **Day 1** (Go, 27 lines): Find the maximum sum of newline-separated 
  groups of numbers. For part 2, find the sum of the top 3 (*easy*,
  requires reading a list of numbers from a file and sorting).

* **Day 2** (Go, 72 lines): Rock/Paper/Scissors simulation (Part 1), with 
  identifying strategies that result in predefined outcome for Part 2 (*easy*,
  requires reading data, simulating simple decision rules)

* **Day 3** (Go 65, Python 23 lines): Sum up characters that are common in two
  halves of some strings (Part 1), and common to groups of 3 strings (Part 2)
  (*easy*, requires converting between numbers and characters, and simple set
  and search operations)

* **Day 4** (Go, 26 lines): Given pairs of numeric ranges, in how many pairs is
  one range entirely contained within another (Part 1), and how many pairs
  overlap at all (Part 2) (*easy*)

* **Day 5** (Go, 50 lines): Simulate execution of instructions to move crates
  from one tower to another, one at a time (Part 1), or in groups (Part 2)
  (*easy*, save time by pre-parsing and hard-coding problem input)

* **Day 6** (Go, 54 lines): Look for first block of 4 (Part 1) or 14 (Part 2)
  non-repeating characters in a string.  (*easy*)

* **Day 7** (Go, 95 lines): Given a list of Unix shell commands and output
  (just ls and cd), parse these and find the sum of space (calculated
  recursively, i.e., including subdirectories) used by all directories <= 100k
  (Part 1), and the size of the smallest directory that would free up a
  required amount of space. (*medium*, required parsing and executing commands
  and saving state)

* **Day 8** (Go, 108 lines): Given a topographical map (matrix) of tree
  heights, count the number of trees that have visibility all the way to the
  edge (Part 1), and the highest "visibility" score of any tree, where that
  score is the product of the numbers of trees less than the current tree in
  each direction. (*medium*)

* **Day 9** (Go, 51 lines): Simulate movement of "knots" along a rope, in
  response to the first knot being moved. For part 1, there are only two knots
  (head and tail), for part 2 there are 10. After the simulation, report the
  number of positions the tail has covered. (*easy*)

* **Day 10** (Go, 42 lines): Simulate accumulator register during a series of
  given ADD or NOOP instructions, and report the accumulator values during 
  selected clock cycles (Part 1). For part 2, simulate drawing of pixels 
  on a screen, using the sequence of acculator values, and report the eight 
  letters that appear. (*medium*)

* **Day 11** (Go, 100 lines): Simulate transfer of objects between a bunch of
  monkeys, with "worry levels" assigned to each object. Each monkey modifies
  the worry level according to some rules, then passes it to one of two
  monkeys, depending on whether the worry level is divisible by that monkey's
  "test" number.  Count up the number of inspections each monkey makes during
  the simulation. The answer is the product of the two highest inspection
  counts.  Trivial (if tedious) for 20 iterations in Part 1, but integer values 
  overflow for 10,000 iterations in Part 2, unless you apply an adjustment 
  that preserves the decision outcomes while keeping the numbers fom getting 
  too large (*hard* for Part 2).

* **Day 12** (Go, 70 lines): Find the lowest cost path (i.e., shortest number
  of steps) through a terrain of letters, from point S to E, allowing
  'increase' (e.g., next letter) of maximum 1. For Part 2, find the shortest
  path from any 'a' cell to 'E'. (*medium*, using yourbasic/graph library)

* **Day 13** (Python, 49 lines): Given pairs of nested lists of numbers, count
  up how many are in the right order according to an arcane comparison function
  (Part 1), then combine all the pair elements into one big list, add a couple
  of marker elements, and sort the list according to the comparison function. 
  For Part 2, report the product of the indices of the two marker elements.
  (*medium*, did in Python because of mixed types)

* **Day 14** (Go, 110 lines): Simulate grains of sand dropping from a hole into
  2-d space.  For Part 1, count how many grains of sand before they start
  dropping of edges of existing rock. For Part 2, add a "foor" below  bottom
  layer of rock, and count how many grains of sand before a pyramid is built,
  and the hole at the top becomes blocked. (*medium*)

* **Day 15** (Go, 166 lines): Given a list of "sensors" and their distance to
  nearest "beacon", find positions in a row that could not possibly have a
  beacon (Part 1), and the possible location of an undetected beacon (i.e.,
  where there is in coverage by known beacons) for Part 2. (*hard*)

* **Day 16** (Go, also 166 lines): Given a network (graph) of closed "valves",
  each with a certain flow rate, connected by "tunnels", find the sequence of
  opening the valves (takes one minute, plus one minute per step to get there)
  that yields the highest possible total flow during a 30-minute period. For
  Part 2, same but try two decisions (one for you and one for the "elephant")
  each time step, over 26 minutes. Used simple depth-first dynamic programming
  solution, recursively tries each feasible candidate unopened valve, excluding
  those for which we wouldn't have enough time to get any flow. Same for Part
  2, but tried all possible pairs of remaining valves, one for each actor (slow
  but works, *very hard*).

* **Day 17** (Go 123 lines + Python 87 lines): Simulate simple geometric shapes
  falling down a shaft, getting moved left and right by gusts of "gas", and
  falling on top of each other. For Part 1, determine the total height of the
  shapes after 2022 have fallen. For Part 2, do the same for 1 000 000 000 000
  shapes (infeasible to simulate, so looked for repeating pattern in height
  deltas, and applied simple math in separate Python script, *hard*)

* **Day 18** (Go, 65 lines): Given a list of 1x1x1 cubes in 3-d space, count up
  surfaces that don't touch another point (Part 1).  For Part 2, only count
  surfaces that are outside the shape (may include some face inside of a
  "tunnel", so can't just look outward from surface). (*medium*)

* **Day 19** (Go, 152 lines): Basically a set of optimizations, to find the
  maximum number of "geodes" that can be produced over 24 periods from a set of
  four types of "robots".  Each robot can produce one mineral of its own kind
  each time period. There are 30 "blueprints" (cost schedules), each of which
  lists the number of each type of ingredient required to build a robot. So
  it's a production plan optimization.  Part 1 asks you to optimize all 30
  schedules, Part 2 only the first 3 blueprints, but for 32 periods instead of
  24 (*hard*, used dynamic programming but linear programming would have been
  possible).

* **Day 20** (Go, 95 lines): Given a list of numbers (7 in sample, but 5000 in
  input), simulate moving each number forward or backward in the (circular)
  list, forward if positive or backward if negative. For Part 1, do this once,
  and report the sum of the values 1000, 2000, and 3000 after zero. For Part 2,
  multiply each number by a huge value, and do it 10 times, report same sum.
  Complicated by *duplicate values* in the  main input, so you can't just look
  for position of a value. Also, iterations in Part 2 are infeasible with large
  multiplier as well as 10 iterations (*hard*).

* **Day 21** (Go, 72 lines): Given a list of variable names, each with either a
  numbers or a simple formula, recursively evaluate the root node (Part 1), and
  find the value for one cell that makes the two sides of the root node equal
  (solved using gradient descent, *medium*).

* **Day 22** (Go, 374 lines): Simulate movement on a 2D map, according to a
  list of instructions, which can either be to move n steps, or to turn 90
  degrees left or right. There are obstacles to avoid, and one wraps around to
  the other side when walking off and edge. For Part 1, the map is in 2D. For 
  Part 3, the map gets folded into a cube (*hard*, created very verbose
  solution that mapped each possible transition from one surface to another).

* **Day 23** (Go, 132 lines): Simulate movement of "elves" on a map, with
  proposed moves rejected if they clash with any of the other "elves". For Part
  1, report the number of free spaces in the rectangle that encloses all the
  elves at the end of round 10.  For Part 2, report the first round in which
  there is no more movement. (*medium*)

* **Day 24** (Go, 175 lines): Find shortest path from entry to exit of a
  rectangular field, avoiding "blizzards" that move every time step. For Part
  2, also move back to entry then back to exit, and add up all the steps. Used
  dynamic programming, depth-first search with memoization of previously found
  best values for each position+time combination (*medium*).

* **Day 25** (Go, 102 lines): Decode a series of numbers which are in base 5
  with some special characters representing negative numbers. Then, add up the
  decimal equivalents, and return the result encoded back into the special
  representation (*hard*). Part 2 was granted for free after after completing
  the other days.

To compile and run a **Go** program
* Change into the directory with the program
* `go mod init day01`  (*only if go.mod does not yet exist*)
* `go build`
* `./day01`  (or whatever name of the executable)
* Days 12 and 16 require installing a graph library: `go get github.com/yourbasic/graph`

To run a **Python** program
* Change into the directory with the program
* `python day06.py`

AK, Dec 2022
