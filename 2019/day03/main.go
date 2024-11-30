// Advent of Code 2019, Day 03
//
// Give two lists of instructions for two "wires", with each instruction
// L/R/U/D followed by a number, traverse a 2D grid along the specified
// directions, to find where the two wires cross (but don't include a wire
// crossing with itself). For Part 1, find the closest intersection from the
// starting point. For Part 2, find the intersection that had the shortest
// combined number of steps to get there.
//
// AK, 30 Nov 2024

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// A point in 2D space, can be used as a map key
type Point struct {
	x, y int
}

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	// Sample values for Part 1:  6, 159, 135
	// Sample values for Part 2: 30, 610, 410
	fname := "sample2.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")

	// Get all locations visited by each wire
	visited1 := traverse(rows[0]) // first "wire"
	visited2 := traverse(rows[1]) // second "wire"

	// Get the points visited by both wires
	intersections := []Point{}
	for v1, _ := range visited1 {
		_, ok := visited2[v1]
		if ok {
			intersections = append(intersections, v1)
		}
	}

	// Go through all intersections, and for Part 1 find closest intersection,
	// for Part 2 the intersection with the lowest total number of steps
	// to get there along both wires
	fmt.Println(len(intersections), "intersections found")
	part1 := -1
	part2 := -1
	for _, p := range intersections {

		// Part 1: closest manhattan distance from 0,0
		dist := abs(p.x) + abs(p.y)
		if part1 == -1 || dist < part1 {
			part1 = dist
		}

		// Part 2: lowest total number of steps to get there along both wires
		steps := visited1[p] + visited2[p]
		if part2 == -1 || steps < part2 {
			part2 = steps
		}
	}
	fmt.Println("Part 1 (s/b 258) =", part1)
	fmt.Println("Part 2 (s/b 12304) =", part2)
}

// Traverse instructions from 0,0, returning map of points visited.
// For part 1, we just need to know which points are visited. For Part 2,
// we need to know the minimum number of steps it took to get there.
func traverse(instr string) map[Point]int {

	visited := map[Point]int{} // empty map of points visited
	p := Point{0, 0}           // starting, current point
	steps := 0                 // number of steps so far, for part 2
	instructions := strings.Split(instr, ",")
	for _, i := range instructions {

		// Extract direction and number of steps from instruction
		dir := i[0]
		dist, _ := strconv.Atoi(i[1:])

		// Create dx,dy values for each step, depending on direction
		var dx, dy int // default zero
		if dir == 'L' {
			dx = -1
		} else if dir == 'R' {
			dx = 1
		} else if dir == 'U' {
			dy = -1
		} else {
			dy = 1
		}

		// Execute steps, updating position and map
		for i := 0; i < dist; i++ {
			p.x += dx
			p.y += dy
			steps += 1
			_, already_visited := visited[p]
			if !already_visited {
				visited[p] = steps
			}
		}
	}

	// Return a map of all the points visited at least once by this wire
	return visited
}

// Absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
