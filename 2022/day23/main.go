// Advent of Code 2022, Day 23
//
// Simulate movement of "elves" on a map, with proposed moves rejected if they
// clash with any of the other "elves". For Part 1, report the number of free
// spaces in the rectangle that encloses all the elves at the end of round 10.
// For Part 2, report the first round in which there is no more movement.
//
// AK, 23 Dec 2022

package main

import (
	"fmt"
)

// Info about an elf
type Elf struct {
	number      int   // elf ID (not used)
	now, consid Point // current position, and proposed next position
	canMove     bool  // flag to indicate whether move is feasible
}

// A point in 2D space
type Point struct {
	x, y int
}

// The list of elves (pointers, because we change the elves)
var elves []*Elf

func main() {

	// Read the input file into a map of positions of the elves
	fname := "sample.txt"
	//fname = "input.txt"  // uncomment to run on real input
	var y int
	for i, l := range readLines(fname) {
		for x := 0; x < len(l); x++ {
			if l[x] == '#' {
				p := Point{x + 1, y + 1}
				elves = append(elves, &Elf{number: i + 1, now: p})
			}
		}
		y++
	}

	// Begin the simulation of rounds
	dxy := []int{-1, 0, 1}
	directions := []byte{'N', 'S', 'W', 'E'}  // gets rotated each iteration
	for round := 1; round <= 10000; round++ { // will stop when no more movement

		// For Part 2, find the number of the first round where no Elf moves?
		moved := false

		// Go through elves, check if has any neighbours (skip if not), then
		// find the first feasible direction she could move
		for _, e := range elves {

			// Initialize state for this round
			e.consid = Point{-999, -999}
			e.canMove = false

			// If nobody next to you in any direction, do nothing
			hasNeighbour := false
			x := e.now.x
			y := e.now.y
			for _, dx := range dxy {
				for _, dy := range dxy {
					if dx == 0 && dy == 0 {
						continue // don't consider current location
					}
					if !free(x+dx, y+dy) {
						hasNeighbour = true
					}
				}
			}
			if !hasNeighbour { // next elf if this one has no neighbours
				continue
			}

			// Consider four directions in current order, choose the first one
			// that meets criteria, i.e.,"If there is no Elf in the N, NE, or
			// NW adjacent positions, the Elf proposes moving north one step."
			// Note that I understood the free conditions to be ORs, but they
			// need to be ANDs.
			// Also note that it is possible for an elf not to find a feasible
			// movement, in which case consid.x and consid.y will remain -999.
			for _, dir := range directions {
				if dir == 'N' {
					if free(x-1, y-1) && free(x, y-1) && free(x+1, y-1) {
						e.consid = Point{x, y - 1}
						break
					}
				} else if dir == 'S' {
					if free(x-1, y+1) && free(x, y+1) && free(x+1, y+1) {
						e.consid = Point{x, y + 1}
						break
					}
				} else if dir == 'E' { // right
					if free(x+1, y-1) && free(x+1, y) && free(x+1, y+1) {
						e.consid = Point{x + 1, y}
						break
					}
				} else if dir == 'W' { // left
					if free(x-1, y-1) && free(x-1, y) && free(x-1, y+1) {
						e.consid = Point{x - 1, y}
						break
					}
				}
			}
		}

		// Simultaneously, each Elf moves to their proposed destination
		// tile if they were the only Elf to propose moving to that position.
		// If two or more Elves propose moving to the same position, none of
		// those Elves move.
		for _, e := range elves {

			// Skip this elf if could not find a place to move
			c := e.consid
			if c.x == -999 && c.y == -999 {
				e.canMove = false
				continue
			}

			// Otherwise reject move if anyone else is considering moving to
			// the same place
			e.canMove = true
			for _, e1 := range elves {
				c1 := e1.consid
				if e1 != e && c1.x == c.x && c1.y == c.y {
					e.canMove = false
					break
				}
			}
		}

		// Move any elves that can move, and record that a move has
		// happened for Part 2
		for _, e := range elves {
			if e.canMove {
				e.now.x = e.consid.x
				e.now.y = e.consid.y
				moved = true
			}
		}

		// Rotate directions left by one
		directions = append(directions[1:], directions[0])

		// For Part 1, after round 10 find the smallest rectangle that contains
		// the Elves, and report how many empty ground tiles does that
		// rectangle contain
		if round == 10 {
			min, max := minMax()
			space := (max.x - min.x + 1) * (max.y - min.y + 1)
			fmt.Println("Part 1 (s/b 110):", space-len(elves))
		}

		// For Part 2, report if there were no moves
		if !moved {
			fmt.Println("Part 2 (s/b 20): no more moves in round", round)
			break
		}
	}
}

// Get min/max coords
func minMax() (Point, Point) {
	var min, max Point
	for i, e := range elves {
		if i == 0 || e.now.x < min.x {
			min.x = e.now.x
		}
		if i == 0 || e.now.x > max.x {
			max.x = e.now.x
		}
		if i == 0 || e.now.y < min.y {
			min.y = e.now.y
		}
		if i == 0 || e.now.y > max.y {
			max.y = e.now.y
		}
	}
	return min, max
}

// Is a point currently free?
func free(x, y int) bool {
	for _, e := range elves {
		if e.now.x == x && e.now.y == y {
			return false
		}
	}
	return true
}
