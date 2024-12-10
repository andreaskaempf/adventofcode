// Advent of Code 2024, Day 10
//
// Given a terrain of 0-9 digits, start at each 0 and walk up any path that
// increments by one each step, until you reach a 9. For Part 1, how many
// 9s are reached. For Part 2, what is the total number of paths from anywhere
// that reach there. Did with simple recursion.
//
// AK, 10 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Matrix of x,y numbers (the data)
var locs [][]int

// A point in 2D, only used for keeping track of which 9s were reached
type Point struct {
	x, y int
}

// For part 1, keep track if which points were reached (they may have been
// reached multiple times, i.e., from each step along the path, but only
// count it once).
var reached map[Point]bool

func main() {

	// Read the input file and convert into matrix
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))
	locs = [][]int{}
	for _, r := range rows {
		l := make([]int, len(r))
		for i := 0; i < len(r); i++ {
			l[i] = int(r[i] - '0')
		}
		locs = append(locs, l)
	}

	// Find each zero and walk from there
	var part1, part2 int
	for y := 0; y < len(locs); y++ {
		for x := 0; x < len(locs[y]); x++ {
			if at(x, y) == 0 {
				reached = map[Point]bool{} // initialize which 9s reached
				part2 += trails(x, y)      // total number of paths to get there
				part1 += len(reached)      // number of 9s reached from this zero
			}
		}
	}

	fmt.Printf("Part 1 = %d, Part 2 = %d\n", part1, part2)
}

// Using recursion, find the number of trails that start from this cell,
// go up by one elevation each step, and end up at 9. This will count
// each step along the way as a separate path (required for Part 2). For
// Part 1, mark all 9s reached.
func trails(x, y int) int {

	// If on a 9, we found a trail
	elev := at(x, y)
	if elev == 9 {
		reached[Point{x, y}] = true // for Part 1
		return 1                    // for Part 2
	}

	// Otherwise look around
	elev++                  // := at(x, y) + 1 // looking for this elevation
	heads := 0              // number of trails from next step that reach a 9
	if at(x, y-1) == elev { // up
		heads += trails(x, y-1)
	}
	if at(x, y+1) == elev { // down
		heads += trails(x, y+1)
	}
	if at(x-1, y) == elev { // left
		heads += trails(x-1, y)
	}
	if at(x+1, y) == elev { // left
		heads += trails(x+1, y)
	}
	return heads
}

// The value of a cell, -1 if invalid
func at(x, y int) int {
	if x >= 0 && x < len(locs[0]) && y >= 0 && y < len(locs) {
		return locs[y][x]
	} else {
		return -1
	}
}
