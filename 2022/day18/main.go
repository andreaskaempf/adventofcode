// Advent of Code 2022, Day 18
//
// Given a list of 1x1x1 cubes in 3-d space, count up surfaces that don't touch
// another point (Part 1).  For Part 2, only count surfaces that are outside the
// shape (may include some face inside of a "tunnel", so can't just look outward
// from surface).
//
// AK, 18 Dec 2022

package main

import (
	"fmt"
	"strings"
)

// A point in 3-d space
type Point struct {
	x, y, z int
}

// Maps of points in space, and points explored (global vars)
var points map[Point]bool
var explored map[Point]bool

func main() {

	// Read the input file into set of points
	lines := readLines("sample.txt")
	lines = readLines("input.txt") // uncomment to use problem input
	points = map[Point]bool{}      // initialize map of points in space
	explored = map[Point]bool{}    // initialize map of points explored
	for i := 0; i < len(lines); i++ {
		nums := strings.Split(lines[i], ",")
		p := Point{atoi(nums[0]), atoi(nums[1]), atoi(nums[2])}
		points[p] = true
	}

	// Part 1: for each point, count up surfaces that don't touch another point
	// Part 2: only count surfaces that are outside the shape (may include some
	// face inside of a "tunnel", so can't just look outward from surface)
	var part1, part2 int
	for p, _ := range points {
		for _, a := range getAdjacent(p) {
			if !points[a] { // Part 1: include surface of this point
				part1++           // if it touches no other
				if freePoint(a) { // Part 2: only include this point if
					part2++ // there is a route from it to "outer space"
				}
			}
		}
	}

	// Show results
	fmt.Println("Part 1 (s/b 64, 4268):", part1)
	fmt.Println("Part 2 (s/b 58, 2582):", part2)
}

// Is point free in any direction, directly or indirectly. "Free" means
// that there is a route from this point to outer space, i.e., not enclosed
// within other points
func freePoint(p Point) bool {

	// Not free if point is solid
	if points[p] {
		return false
	}

	// If this point is free in any direction, return true
	if free(p, -1, 0, 0) || free(p, 1, 0, 0) ||
		free(p, 0, -1, 0) || free(p, 0, 1, 0) ||
		free(p, 0, 0, -1) || free(p, 0, 0, 1) {
		return true
	}

	// Mark this point as explored then look at the size points around it,
	// this point is free if any of those adjacencies are "free"
	explored[p] = true
	for _, a := range getAdjacent(p) {
		if !explored[a] && freePoint(a) {
			return true
		}
	}

	// If no adjacent free point found, return false
	return false
}

// Is space free next to given point, in given direction?
// One of dx/dy/dz must be 1 or -1 to specify direction (others zero).
// Only looks 100 in any direction, adjust this if necessary.
func free(p Point, dx, dy, dz int) bool {
	for i := 1; i < 100; i++ { // adjust if necessary
		p1 := Point{p.x + dx*i, p.y + dy*i, p.z + dz*i} // look out in direction
		if points[p1] {                                 // if that point is occupied, not free in this direction
			return false
		}
	}
	return true // must be free in this direction
}

// Get coordinates of the 6 surrounding points, i.e., above/below,
// left/right, or front/back
func getAdjacent(p Point) []Point {
	return []Point{Point{p.x - 1, p.y, p.z}, Point{p.x + 1, p.y, p.z},
		Point{p.x, p.y - 1, p.z}, Point{p.x, p.y + 1, p.z},
		Point{p.x, p.y, p.z - 1}, Point{p.x, p.y, p.z + 1}}
}
