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
