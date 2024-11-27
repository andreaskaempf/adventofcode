# Advent of Code 2021

My solutions for the Advent of Code 2021.

* **Day 1** (C, Python, Rust): Count number of increases in a text file of
    lines (*easy*); Rust version is my first experiment in this language
* **Day 2** (Go, Python, Rust): Process the state of a submarine by interpreting 
    instructions "forward", "up", "down", etc. (*easy*)
* **Day 3** (Julia, Python, Rust): Read list of binary numbers to establish 
    codes based on column-wise bit frequency (*easy*)
* **Day 4** (Go, Rust): Simulate game of bingo, by applying list of called numbers
    until first board wins (for part 2, also find the last board
    that wins -- *easy*)
* **Day 5** (Go, Rust): Trace lines on a grid, given starting and ending 
    coordinates of each line, and count up where 2 or more lines cross 
    (part 1 only includes horizontal and vertical lines, part 2 diagonal 
    as well -- *easy*).
* **Day 6** (Python, Go, Rust): Simulate growth of a population of "lanternfish",
    tricky because performance crawls if you try to simulate individual
    creatures (used dictionary/array of fish by age for speed -- *medium*).
* **Day 7** (Julia): Find target position that minimizes fuel use for 
    a population of crabs that need to move there, under different 
    assumptions about fuel use per each crab's distance moved 
    (optimization problem, but used brute force -- *easy*).
* **Day 8** (Julia): Decipher scrambled digits on a digital display 
    (failed using logic so solved using brute force, *very hard*)
* **Day 9** (Julia): Find lowest points in a 2-dimensional space, and the 
    areas around these low points, bounded by edges or 9s (*medium*).
* **Day 10** (Python): Identify and resolve different types of mis-matched
    brackets in a file of strings (*easy*).
* **Day 11** (Julia): Simulate energy growth of a field of "octopusses", with each
    exploding when they reach a level, infecting those around them; for part 2,
    identify when they all explode in the same iteration (*easy*).
* **Day 12** (Julia): Find the number of paths from start to end in a graph, allowing repeat
    visits for "large caves" (capitalized), but either no (part 1) or at
    most 2 (part 2) visits for "small" (lower case) ones (*hard*).
* **Day 13** (Julia): Repeatedly "fold" a matrix of 1/0 values along 
    horizontal or vertical lines, and (part 1) count up the number of 
    points remaining after one fold, and (part 2) visualize the matrix
    after all folds to yield a dot matrix alphabetic code (*easy*).
* **Day 14** (Python): Apply character insertion rules to a sequence of characters,
    by managing list of possible transitions rather than building up string
    of characters (otherwise runs out of memory in Part 2, *hard*).
* **Day 15** (Go, Julia): Find the lowest cost path through a graph, starting
    at the top left, and ending up at the bottom right, adding up any cells 
    you enter along the path that minimizes total cost; Go using graph library,
    Julia with Dijstra shortest-path algorithm (*medium*).
* **Day 16** (Python): Parse hierarchy of binary codes into numeric literals
    and operations. For part 2, recursively evaluate the resulting expression
    using a LISP-like interpreter approach (*medium*).
* **Day 17** (Python): Simple projectile simulation, with search for
    initial "velocities" that achieve highest position, and total
    number of possible velocity values that reach target area; just 
    used a simple grid search (*easy*).
* **Day 18** (Go): Calculate operations on nested lists of numbers, first
    "adding" (concatenating them), then reducing them according
    to arcane rules. Part 1: add up each line in file, reducing
    after each line, and report the final "magnitude" (calculated
    by recursively collapsing inner pairs of numbers). For Part 2,
    find the pair of input lines which, when added together and
    "reduced", yield the highest magnitude (*hard*).
* **Day 19** (Go): Match up 3-dimensional cubes in space, adjusting x,y,z
    offset and also rotation along any 3 axes, so that at least 12 points in
    the each pair of cubes line up exactly. Then, calculate the maximum
    distance between the cubes (*hard*).
* **Day 20** (Julia): Transform an image by successively replacing pixels with
    values looked up from a translation table, the index being the value of the
    9 cells surrounding each pixel, converted from binary to decimal. Much
    complicated by the fact that the input data has a 1 in the first position
    of the translation table, meaning that empty areas are filled with 1s,
    which muck up the pixel count (*hard*).
* **Day 21** (Go): Simulate a game of rolling dice and moving round a board,
    trivial in part 1. In part 2, fork a set of parallel "universes" with
    identical state after every throw of a 3-sided die, and get
    the number of universes in which the winner won (*hard*).
* **Day 22** (Go): Turn on/off points in space, defined by 3-d ranges (like
    rectangular cubes). Part 1 quite *easy* (basically used brute force), but
    for Part 2 used recursive evaluation of volumes solution subtracting
    intersections (*hard*).
* **Day 23** (Go): Find the most economical solution to a board game, involving
    4 (later 8) pieces from random starting tunnels to ordered destination
    tunnels via a corridor, sort of like the Towers of Hanoi. Solved the first
    part (8 pieces) on paper, second part using recursive depth-first search,
    eliminating branches that exceeded best solution found so far (*hard*).
* **Day 24** (Go): Check whether a 14-digit code is valid based on
    assembly-language like series of computations, and find the highest (Part 1) 
    or lowest (Part 2) codes that are valid, i.e, leave a "z" register at
    zero after the computations (*very hard*, used brute force with parallelization
    and early pruning of infeasible solutions, takes about 30 mins per part).
* **Day 25**: Simulation of "sea cucumber" particle movement, iterating until
    the system reaches a stable state (i.e., no more changes). Part 1 done
    (*easy*), and Part 2 granted after completing all previous problems.

To compile and run a Go program
* Change into the directory with the program
* go build day02.go
* ./day02

To run a Julia program
* Change into the directory with the program
* julia day02.jl

To compile and run a C program
* Change into the directory with the program (e.g., cd day01)
* gcc day01.c -o day01
* ./day01

And of course, to run a Python program
* Change into the directory with the program
* python day06.py

AK, Dec 2021
