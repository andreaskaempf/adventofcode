// Advent of Code 2020, Day 11
//
// Change the state of a "seating plan" depending on available seats
// immediately adjacent to every seat (Part 1), or visible in any direction
// (Part 2).
//
// AK, 14/01/2022

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Set to false for part 1, true for part 2
const part2 = true

func main() {

	// Read each line of input file
	lines := [][]byte{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		l := []byte{}
		for i := 0; i < len(t); i++ {
			l = append(l, t[i])
		}
		lines = append(lines, l)
	}

	// Show  initial state
	fmt.Println("Starting:")
	printBoard(lines)

	// Iterate until no more changes
	iter := 0
	for {

		// Make a copy: always look at the current state, but make changes
		// to a copy, so the changes can be "simulataneous"
		lines1 := [][]byte{}
		for _, l := range lines {
			l1 := make([]byte, len(l), len(l))
			copy(l1, l)
			lines1 = append(lines1, l1)
		}

		// The following rules are applied to every seat simultaneously:
		// 1. If a seat is empty (L) and there are no occupied seats adjacent
		//    to it, the seat becomes occupied.
		// 2. If a seat is occupied (#) and four or more seats adjacent to it
		//    are also occupied, the seat becomes empty.
		// Otherwise, the seat's state does not change.
		changed := false
		for r := 0; r < len(lines); r++ {
			for c := 0; c < len(lines[r]); c++ {
				var nOccup, thresh int
				if part2 {
					nOccup = adjacentOccupied2(lines, r, c)
					thresh = 5
				} else {
					nOccup = adjacentOccupied1(lines, r, c)
					thresh = 4
				}
				if lines[r][c] == 'L' && nOccup == 0 {
					lines1[r][c] = '#'
					changed = true
				}
				if lines[r][c] == '#' && nOccup >= thresh {
					lines1[r][c] = 'L'
					changed = true
				}
			}
		}

		// Show result of this iteration
		iter += 1
		fmt.Printf("\nIteration %d:\n", iter)
		printBoard(lines1)

		// Stop if no more changes, otherwise prepare for next iteration
		lines = lines1
		if !changed {
			break
		}
	}

}

// Part 1: count the  number of adjacent seats around a given seat that are
// occupied, just immediate adjacenies up/down/left/right/diagonal
func adjacentOccupied1(lines [][]byte, r int, c int) int {
	result := 0
	for ri := r - 1; ri <= r+1; ri++ {
		for ci := c - 1; ci <= c+1; ci++ {
			if ri == r && ci == c {
				continue
			}
			if ri >= 0 && ri < len(lines) && ci >= 0 && ci < len(lines[ri]) && lines[ri][ci] == '#' {
				result += 1
			}
		}
	}
	return result
}

// Part 2: count the  number of visible seats around a given seat that are
// occupied, in any direction up/down/left/right/diagonal
func adjacentOccupied2(lines [][]byte, r int, c int) int {
	result := 0
	diffs := []int{-1, 0, 1}
	for _, dr := range diffs {
		for _, dc := range diffs {
			if !(dr == 0 && dc == 0) && look(lines, r, c, dr, dc) == '#' {
				result++
			}
		}
	}
	return result
}

// From a given coordinate, look in any direction, expressed as +/- dx
// and dy, so dx = 1 and dy = 1 to look diagonally downward, or dx = -1 and
// dy = 0 to look left. Returns the first non '.' found, or '.' if that's
// all that can be seen.
func look(lines [][]byte, r, c, dr, dc int) byte {
	ri := r + dr
	ci := c + dc
	for ri >= 0 && ri < len(lines) && ci >= 0 && ci < len(lines[ri]) {
		if lines[ri][ci] != '.' {
			return lines[ri][ci]
		}
		ri += dr
		ci += dc
	}
	return '.'
}

// Count the  number of seats on a board that are occupied
func occupied(lines [][]byte) int {
	result := 0
	for ri := 0; ri < len(lines); ri++ {
		for ci := 0; ci < len(lines[ri]); ci++ {
			if lines[ri][ci] == '#' {
				result += 1
			}
		}
	}
	return result
}

// Print a graphic representation of the board
func printBoard(lines [][]byte) {
	for _, l := range lines {
		fmt.Println(string(l))
	}
	fmt.Printf("%d seats occupied\n", occupied(lines))
}
