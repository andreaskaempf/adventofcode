// Advent of Code 2024, Day 06
//
// Given a grid with starting point, initial direction, and some
// ostacles, move in direction until you hit an obstacle (turn
// 90 degrees right), or you exit grid (you're done). For Part 1,
// count up the number of cells visited. For Part 2, count up how
// many new obstacles would cause an endless loop.
//
// AK, 06 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// A 2D point
type Point struct {
	x, y int
}

// Global variables
var rows [][]byte      // rows of characters
var obs map[Point]bool // locations of obstacles
var start Point        // starting point

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	// Find the starting position, i.e., caret, and mark positions of all
	// obstacles
	obs = map[Point]bool{} // locations of obstacles
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[r]); c++ {
			ch := rows[r][c]
			p := Point{c, r}
			if ch == '^' {
				start = p
			} else if ch == '#' {
				obs[p] = true
			}
		}
	}
	fmt.Println("Start:", start)

	// Part 1: number of points visited
	fmt.Println("Part 1 (4663) =", walk())

	// Part 2: count how many new obstacles would cause endless loop
	// (using brute force)
	part2 := 0
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[r]); c++ {

			// Skip location if start or an obstacle
			p := Point{c, r}
			if p == start || obs[p] {
				continue
			}

			// Temporarily mark this location as an obstacle and see if
			// it causes endless loop
			obs[p] = true
			if walk() == -1 {
				part2 += 1
			}
			obs[p] = false
		}
	}
	fmt.Println("Part 2 (1530) = ", part2)
}

// Walk from starting position, return number of steps or -1 if loop
func walk() int {

	visited := map[Point]bool{} // locations visited
	dy := -1                    // start in direction up
	dx := 0                     // horizontal to move up
	pos := start                // always same starting position
	iters := 0                  // to detect endless loop
	for {

		// For part 2, if too many iterations, stop here
		iters += 1
		if iters > 10000 { // obtained through trial and error
			return -1
		}

		// Mark this point as visited and choose next candidate location
		visited[pos] = true
		next := Point{pos.x + dx, pos.y + dy}
		if next.x < 0 || next.x >= len(rows[0]) || next.y < 0 || next.y >= len(rows) {
			break // out of bounds, we're done
		}

		// If hit obstacle, change direction, otherwise move
		if obs[next] {
			dx, dy = chgDir(dx, dy)
		} else {
			pos = next
		}
	}

	// Return number of points visited
	return len(visited)
}

// Change direction, always 90 degrees right
func chgDir(dx, dy int) (int, int) {
	if dy == -1 && dx == 0 { // up -> right
		dy = 0
		dx = 1
	} else if dy == 0 && dx == 1 { // right -> down
		dy = 1
		dx = 0
	} else if dy == 1 && dx == 0 { // down -> left
		dy = 0
		dx = -1
	} else if dy == 0 && dx == -1 { // left -> up
		dy = -1
		dx = 0
	}
	return dx, dy
}
