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

