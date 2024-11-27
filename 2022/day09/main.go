// Advent of Code 2022, Day 09
//
// Simulate movement of "knots" along a rope, in response to the first knot
// being moved. For part 1, there are only two knots (head and tail), for part
// 2 there are 10. After the simulation, report the number of positions the
// tail has covered.
//
// AK, 9 Dec 2022

package main

import (
	"fmt"
)

// A position in 2D space
type Position struct {
	x, y int
}

func main() {

	// Read the input file, a list of "Dir n" instructions
	//lines := readLines("sample2.txt")  // use sample2.txt for part 2
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Part 1: number of tail positions visited, with just 2 knots
	fmt.Println("Part 1 (s/b 13, 5735):", simulate(lines, 2))

	// Part 2: same, with 10 knots
	fmt.Println("Part 2 (s/b 36, 2478):", simulate(lines, 10))

}

// Simulate movement of "knots" along a rope, return the number of positions
// visited by the last "knot"
func simulate(lines []string, knots int) int {

	// Process each line
	tVisited := map[Position]int{}              // places the tail has visited
	positions := make([]Position, knots, knots) // current positions of 10 knots
	tVisited[Position{0, 0}] = 1                // tail has already visited 0,0
	for _, l := range lines {

		// Parse instruction
		dir := l[0]      // R/U/L/D
		n := atoi(l[2:]) // steps to move

		// Do each step of the instruction
		for i := 0; i < n; i++ {

			// Move the head one step
			h := &positions[0]
			if dir == 'R' {
				h.x += 1
			} else if dir == 'L' {
				h.x -= 1
			} else if dir == 'U' {
				h.y -= 1
			} else { // down
				h.y += 1
			}

			// Process each subsequent "knot"
			for k := 1; k < knots; k++ {

				// The current and previous "knot"
				h = &positions[k-1] // previous knot
				t := &positions[k]  // this one

				// If knot is "touching" the previous one, no need to adjust
				if abs(h.x-t.x) <= 1 && abs(h.y-t.y) <= 1 {
					continue
				}

				// If the knot is two steps up, down, left, or right from the
				// previous one, move previous one step in that direction
				if t.y == h.y && abs(t.x-h.x) == 2 { // left/right
					t.x += ifElse(t.x > h.x, -1, 1)
				} else if t.x == h.x && abs(t.y-h.y) == 2 { // up/down
					t.y += ifElse(t.y > h.y, -1, 1)
				} else {
					// Otherwise, move tail one step diagonally to keep up
					t.y += ifElse(t.y > h.y, -1, 1)
					t.x += ifElse(t.x > h.x, -1, 1)
				}
			}

			// Remember the new position of the last knot
			tVisited[positions[knots-1]] = 1
		}
	}

	// Return the number of positions visited by the last knot
	return len(tVisited)
}
