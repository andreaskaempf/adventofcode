# Advent of Code 2024

My solutions for the Advent of Code 2024, 
see https://adventofcode.com/2024

* **Day 1** (Go, Rust): Given rows of pairs of numbers, extract the two 
  lists, sort them, and find the sum of the absolute difference 
  (Part 1). For Part 2, sum up the product of each number in the 
  left list, times the number of times it occurs in the right list.

* **Day 2** (Go, Rust): Given rows of numbers, check each for the following 
  conditions: all numbers are increasing/decreasing, and by 1 to 3 
  each step. For Part 1, count up how many rows meet the condition. 
  For Part 2, count up how many rows would meet the condition, if any 
  digit were removed.

* **Day 3** (Go): Given a string, find all embedded "mul(3,4)" 
  instructions, and add up the results of the multiplications. 
  For Part 2, "do()" and "don't()" turn multiplication on/off.

* **Day 4** (Go): Find all instances in XMAS in a matrix of text 
  in any direction (part 1), or MAS cross pattern (part 2).

* **Day 5** (Go): You are given a list of page number pairs, indicating 
  that the right page must be printed after the left page number. You are 
  also given a list of manual "updates", each a list of page numbers. For 
  each update, determine if the pages are in the correct sequence, and add up
  the middle page numbers.  For Part 2, rearrange the  incorrectly ordered
  updates, and add up the middle page numbers.

* **Day 6** (Go): Given a grid with starting point, initial direction, 
  and some ostacles, move in direction until you hit an obstacle (turn 90
  degrees right), or you exit grid (you're done). For Part 1,
  count up the number of cells visited. For Part 2, count up how
  many new obstacles would cause an endless loop.

* **Day 7** (Go): Given a number of target: n n n lines, add up all the 
  targets for which the n values can be computed left-to-right using some
  combination of + or * operators. For Part 2, there is third || operator,
  which string-concatenates the operands. There is no operator precedence, just
  left-to-right. Key to this was coming up with all the combinations of 2 or 3
  values across any number of columns. For Part 1, I did this with binary
  arithmetic, but had to code an arbitrary-base counting function for Part 2.

* **Day 8** (Go): Given letters on a grid, extrapolate the diagonal distance 
  between each pair of the same letter, in either direction, and count up
  the total number of cells occupied by new entries. For Part 2, extrapolate
  in a line either direction.

* **Day 9** (Go): Simulate two different strategies for disk fragmentation, 
  taking a list of digits where each pair of digit is a file size followed 
  by space after the data (both in blocks). For Part 1, move parts of files 
  to first available space, and calculate a "checksum".  For Part 2, you move 
  the whole file (all blocks) but start with the file with highest ID.

* **Day 10** (Go): Given a terrain of 0-9 digits, start at each 0 and walk 
  up any path that increments by one each step, until you reach a 9. For Part
  1, how many 9s are reached. For Part 2, what is the total number of paths
  from anywhere that reach there. Did with simple recursion.

* **Day 11** (Python, Go): Apply transformations to a list of numbers, 
  so that zeros become ones, numbers with an even number of digits get split in
  half, or the number gets multiplied by 2024 (only one test per number). Do
  all these simulataneously, i.e., changing a copy of the data and referencing
  the original without changing it. For Part 1, it was okay to naively build up
  a list over 25 iterations, but for Part 2 there were 75 iterations, and
  memory ran out after about 45 iterations. So changed formulation to use a
  dictionary, since the numbers could be processed in any order, and many of
  the numbers are repeated.

* **Day 12** (Go): Given a matrix of letters that represent different polygons, 
  find the area of each polygon, and the length of the perimiter. For Part 1,
  the sumproduct of the areas * perimiters.  For Part 2, count up the number of
  sides, and calculate the sumproduct of area * sides.  For Part 1, just
  recursively explore each shape to capture the polygon, total area is just the
  number of cells, and perimiter just requires adding up where
  left/right/up/down is something different for each cell. Part 2 was much more
  difficult, and I ended up doing each side separately. E.g., for the left
  edge, finding all cells that have nothing or something different to the left,
  grouping these by column, and counting the number of sequential blocks within
  each column group. Doing this in each of the four directions and adding up
  the number of blocks gives the right answer.

* **Day 13** (Python + GMPL): Optimize a set of "machines", by finding the number
  of presses for buttons A and B that move a "claw" to the right location to pick
  up a prize. For each machine, the button moves the claw a given x and y
  distance. For Part 1, did this easily using integer optimization with Pulp.
  For Part 2, had to add a huge number to each prize location, causing the Pulp
  optimizer to fail, presumably numeric overflow. Did Simplex in Go using
  gonum, and tried to find approximate integer answer around the optimal
  floating point result, but did not work.  Finally, tried command line GLPK,
  and found it could handle the large integers. So the final solution (which
  works for Parts 1 and 2) writes a GLPK model, runs the solver, and parses the
  result. And it's really fast!

* **Day 14** (Go): Given a list of "robots" in a ~100x100 2D space, each with an 
  x,y position and an dx,dy per-second velocity, simulate 100 iterations of the
  robots moving around, wrapping around at the edges. For Part 1, answer is the
  product of the number of robots in each of the four quadrants after 100
  iterations. For Part 2, it's the first iteration that yields ASCII art of a
  Christmas tree. First part was easy simulation. For Part 2, just drew maps of
  configuration after each iteration where there appeared to be lots of
  "robots" in one row, this was sufficient to find the answer by visual
  inspection.

* **Day 15** (Go): Given a map of walls and movable "boxes",  and a list of 
  up/down/left/right instructions, move a "robot" around, moving boxes as much
  as possible if they are in the way. For Part 2, the width of walls and boxes
  are doubled, but the robot remains width one, making movement tricker (since
  boxes can overlap). In both parts, answer is calculated from final position 
  of all the boxes. Parts 1 and 2 in separate Go files.

* **Day 16** (Go): Given a maze, find the shortest path, with the twist that 
  90 degree turns have a cost of 1000. For Part 2, count up all the squares
  that were traversed while reaching the end at the lowest cost. Did this using
  Djikstra algorithm, adapted to deal with the added cost of turns, and keeping
  track of squares traversed.

* **Day 17** (Rust): Simulate execution of an 8-bit machine language program with 
  opcodes and arcane instruction rules such as dividing, xoring, outputting
  numbers, and jumping. For Part 1, show the output of running a short
  "program", a list of 16 numbers. For Part 2, find the "register A" starting
  value that causes the program to output the same 16 values as the input
  program. Part 1 was simple implementation of virtual CPU. For Part 2,
  inspected outputs for incremental input values, and used patterns found to
  converge on likely input ranges, and used brute force from there.

* **Day 18** (Rust): Find shortest distance from one corner of a 71x71 grid to 
  the oppposite corner, avoiding obstacles from a list of coordinates. For Part
  1, just find the distance, using the first 1024 obstacles in the list. For
  Part 2, find the coordinates of the first obstacle that make it impossible to
  reach the location. Implemented Djikstra algorithm in Rust, works very well.

* **Day 19** (Go): Given a list of short patterns of letters ("stripes on a 
  towel"), check a series of long patterns ("designs") to see if they can be
  composed of any combination of the short patterns.  For Part 1, determine how
  many of the designs can be made up of patterns. For Part 2, how many
  combinations in total. Did this with simple recursion, cycling through
  patterns that match head of a string, then doing same for the tail. For the
  problem input, had to add memoization of sub-designs already encountered as
  was taking too long.

* **Day 20** (Go): Given a maze with only one path through it, find shortcuts of 
  length 2 that cut through walls and shorten the path. For Part 1, count up
  the number of shortcuts that shorten the path by 100 or more. For Part 2,
  find shortcuts of length 2 to 20, also counting up the number that reduce the
  path by 100 or more steps. This one took me a long time, because I thought
  the problem was more complex than it is. Built a Dijkstra implementation that
  temporarily ignored walls for two steps, also a recursive traversal that did
  the same. In the end, found the answer by looking at all pairs of points
  along the path, that have a combined horizontal + vertical distance between
  them of length 2 (Part 1) or between 2-20 (Part 2), and measured the distance
  saved, by using the difference between the sequential step numbers along the
  two path points, plus the length of the shortcut itself, to get the new
  distance.

* **Day 21 (Go)**: You are given a numeric keypad on which to type five 
  4-character codes. But you cannot type directly, but must mobilize a robot
  arm with a different keypad equipped with arrows and an Enter key.  That arm
  must be controlled by another similar robot, which you can control using
  another keypad equipped with arrows. So two levels of indirection, 25 for
  Part 2. You must determine the length of each sequence that you type, in
  order to ultimately cause each code to be entered on the numeric keypad.
  This was very tricky, because upstream robots will have different path
  lengths depending on the paths of downstream robots. Ultimately, the solution
  for Part 2 involved recursively expanding paths, starting with the number pad
  and moving away from that. In order to make the problem tractable in terms of
  time and memory, the solution maintains a cache of sequence lengths at each
  level of recursion.  This is possible because (A) the downstream keypads
  (after the number pad) always have to return the to A key at the end of each
  sequence, so you can treat sequences after the first as independent, and (B)
  you only need to know the final length, not the sequence itself. So the
  interim counters can be memoized, making the calculation very fast.

* **Day 22 (Go)**: Use an arcane series of calculations to calculate the next 
  2000 "secret numbers" starting with first for each of about 1600 players. For
  Part 1, add up the 2000th numbers. For Part 2, find optimal revenue that can
  be achieved, by deriving price from last digit of each secret number, the
  delta from each subsequent pair, and finding a sequence of four price changes
  common to all players, such that the revenue from that player is the price at
  the end of the first occurrence of the sequence. Used brute force, runs in
  about 2 minutes in Go.

* **Day 23** (Go): Given a list of connected node pairs, find all groups of 3 
  that are connected to each other, where at least on of the nodes starts with 
  't' (Part 1).  For Part 2, find the largest set of nodes that are all connected
  to each other, sort names and join with commas. Did this directly, no graph
  library (would not have helped).

* **Day 24** (Go): not complete

* **Day 25** (Go): Input is a series of 6x5 blocks of # and . characters, 
  representing locks (heights from top down) or keys (heights from bottom up).
  Find out how many key & lock pairs fit, i.e., heights to not overlap.
  Checking was trivial, but parsing the input into arrays of heights was a
  chore. There is no Part 2, granted automatically when you complete all the
  other puzzles.

To compile and run a **Go** program
* Change into the directory with the program
* `go mod init day01`  (*only if go.mod does not yet exist*)
* `go build`
* `./day01`  (or whatever name of the executable)

To run a **Python** program
* Change into the directory with the program
* `python day06.py`

To compile and run a **Rust** program
* Change into the directory with the program
* `rustc day01.rs`
* `./day01`  (or whatever name of the executable)
* If the program requires external dependencies ("crates"), you will 
  have to do `cargo init`, move dayXX.rs to src, add the crate to
  Cargo.toml, and then `cargo build` to compile; the executable will
  be somewhere in the ./target directory.

To compile and run a **Zig** program
* Change into the zig directory, e.g., `cd day05/zig`
* `zig build` (*debug mode, day05 runs in ~15 mins*)
* `zig build -Doptimize=ReleaseFast` (*fast mode, takes < 2 mins*)
* To run: `zig-out/bin/day05`

To compile and run a **C** program
* `gcc -O3 day05.c -o day05`
* To run: `./day05`

