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

