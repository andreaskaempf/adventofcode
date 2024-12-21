// Advent of Code 2024, Day 20
//
// Given a maze with only one path through it, find shortcuts of length 2 that
// cut through walls and shorten the path. For Part 1, count up the number of
// shortcuts that shorten the path by 100 or more. For Part 2, find shortcuts
// of length 2 to 20, also counting up the number that reduce the path by 100
// or more steps.
//
// This one took me a long time, because I thought the problem was more
// complex than it is. Built a Dijkstra implementation that temporarily
// ignored walls for two steps, also a recursive traversal that did the
// same. In the end, found the answer by looking at all pairs of points
// along the path, that have a combined horizontal + vertical distance
// between them of length 2 (Part 1) or between 2-20 (Part 2), and
// measured the distance saved, by using the difference between the
// sequential step numbers along the two path points, plus the length
// of the shortcut itself, to get the new distance.
//
// Old code (368 lines of it) is saved in not_used.go in case useful for other
// purposes.
//
// AK, 20-21 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Point struct {
	x, y int // coordinates
	seq  int // the sequence in the path, 0 for start
}

// Global variables
var rows [][]byte    // input data as matrix of chars
var start, end Point // start and end positions
var path []Point     // list of points on the only path

func main() {

	// Read the input file and find default (only) path
	fname := "sample.txt"
	fname = "input.txt"
	readMap(fname)
	fmt.Println("Start =", start, ", end =", end, ", path length =", len(path)-1)

	// Find "cheats" within the given length range that shorten path, and
	// show the number that save 100+ steps
	fmt.Println("Part 1 =", countCheats(2, 2, false))  // 1399
	fmt.Println("Part 2 =", countCheats(2, 20, false)) // 994807
}

// Count the number of "cheats" within a maze, that allow you to take a
// shortcut of a given length, i.e., walk through walls. For Part 1, a cheat
// must be activated when the distance from one point on the path to another
// is exactly two positions apart, vertically and/or horizontally. The "saving"
// is the difference between the sequence numbers (distance from start) between
// those two points, minus the steps you had to take for the shortcut.
// For Part 2, consider cheats of between 2 and 20 steps away.
func countCheats(from, to int, showFreq bool) int {

	freq := map[int]int{}     // for debugging
	var count int             // counter 100+ steps saved (for answer)
	for _, p1 := range path { // each point on original path
		for _, p2 := range path { // each point on the path
			dist := abs(p1.x-p2.x) + abs(p1.y-p2.y) // horiz/vertical distance
			if dist >= from && dist <= to {         // distance within range?
				saving := p2.seq - p1.seq - dist // steps saved by shortcut
				if saving >= 0 {                 // update frequency for debugging
					freq[saving]++
				}
				if saving >= 100 { // update final answer
					count++
				}
			}
		}
	}

	// Show frequency counts for debugging
	if showFreq {
		fmt.Println(freq)
	}

	// Return final answer
	return count
}

// Read the file and assemble list of sequential points along the path from S to E
func readMap(filename string) {

	// Read into rows of bytes
	data, _ := ioutil.ReadFile(filename)
	rows = bytes.Split(data, []byte("\n"))

	// Find start and end points (global)
	for ri, r := range rows {
		for ci, c := range r {
			if c == 'S' {
				start = Point{x: ci, y: ri, seq: 0}
			} else if c == 'E' {
				end = Point{x: ci, y: ri, seq: 0}
			}
		}
	}

	// Set 'S' and 'E' to periods so we don't have to keep checking for them
	rows[start.y][start.x] = '.'
	rows[end.y][end.x] = '.'

	// Walk from start to end to assemble default path
	p := start
	path = append(path, start)
	visited := map[Point]bool{}
	for p != end {

		// Mark this point as visited
		visited[p] = true

		// Look in each direction for the next unvisited point
		next := []Point{Point{p.x, p.y - 1, 0}, Point{p.x + 1, p.y, 0},
			Point{p.x, p.y + 1, 0}, Point{p.x - 1, p.y, 0}}
		for _, p1 := range next {
			if at(p1) == '.' && !visited[p1] {
				path = append(path, p1)
				p = p1
				break
			}
		}
	}

	// Assign sequence numbers last
	for i := range path {
		path[i].seq = i
	}
	end.seq = len(path) - 1
}

// What's at a location, return hash if out of range
func at(p Point) byte {
	if p.y < 0 || p.y >= len(rows) || p.x < 0 || p.x >= len(rows[p.y]) {
		return '#'
	} else {
		return rows[p.y][p.x]
	}
}

// Absolute value of a number
func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
