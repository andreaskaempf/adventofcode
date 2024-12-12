// Advent of Code 2024, Day 12
//
// Given a matrix of letters that represent different polygons, find the area
// of each polygon, and the length of the perimiter. For Part 1, the sumproduct
// of the areas * perimiters.  For Part 2, count up the number of sides, and
// calculate the sumproduct of area * sides.  For Part 1, just recursively
// explore each shape to capture the polygon, total area is just the number of
// cells, and perimiter just requires adding up where left/right/up/down is
// something different for each cell. Part 2 was much more difficult, and I
// ended up doing each side separately. E.g., for the left edge, finding all
// cells that have nothing or something different to the left, grouping these
// by column, and counting the number of sequential blocks within each column
// group. Doing this in each of the four directions and adding up the number of
// blocks gives the right answer.
//
// AK, 12 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

// A point in 2D space, so x,y coordinates can be used as key in hashmap (dictionary)
type Point struct {
	r, c int
}

// Global variables
var rows [][]byte          // rows of characters, i.e., the input data
var visited map[Point]bool // which points have been visited
var points []Point         // points contained in an area
var part1, part2 int       // answers for Parts 1 and 2

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	// Keep looking for unexplored cells and explore them
	visited = map[Point]bool{}
	for findAndExplore() { // this function returns false when nothing left to explore
	}
	fmt.Printf("Part 1 (1930 / 1363682) = %d, Part 2 (1206 / 787680) = %d\n", part1, part2)
}

// Find the first unexplored cell, and explore from there, return false
// when no more unexplored cells
func findAndExplore() bool {
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[r]); c++ {
			p := Point{r, c}
			if !visited[p] {
				points = []Point{}        // initialize list of points visited, for Part 2
				area, perim := explore(p) // find every connected cell, calc area and perimiter
				sides := countSides()     // use points slice to count up sides
				part1 += area * perim     // part 1 is area times perimiter
				part2 += area * sides     // part 2 uses number of sides
				return true               // if we processed an area, don't continue
			}
		}
	}
	return false
}

// Explore from a location: find all neighboring cells with the same letter,
// and recursively explore each
func explore(p Point) (int, int) {

	// Mark this cell as visited, and add it to the list of points for this block
	ch := at(p)
	visited[p] = true
	points = append(points, p)

	// The area will be this point plus all others connected
	area := 1
	perim := perims(p)

	// Recursively explore left/right/up/down from this point
	for _, p1 := range neighbours(p) {
		if at(p1) == ch && !visited[p1] {
			area1, perim1 := explore(p1)
			area += area1
			perim += perim1
		}
	}

	// Return the area and perimiter for this part of the block explored. The top
	// level call to this function will return the area and perimiter of the entire block
	return area, perim
}

// Count up the perimiters for a location
func perims(p Point) int {
	y := 0
	ch := at(p)
	for _, p1 := range neighbours(p) {
		if at(p1) != ch {
			y += 1
		}
	}
	return y
}

// Count up the sides, using list of points visited. A side is just a
// a group of points with the same r value, and contiguous c values.
// Or the same c value, and sequential r values
func countSides() int {

	groups := 0

	// Count left sides
	edge := getEdgePoints(0, -1)   // all points on left side
	colGroups := groupCols(edge)   // group by columns
	for _, cg := range colGroups { // within each column group
		rows := getRowNumbers(cg)              // just the row numbers
		groups += countConsecutiveGroups(rows) // number of consecutive groups
	}

	// Count right sides
	edge = getEdgePoints(0, 1)     // all points on right side
	colGroups = groupCols(edge)    // group by columns
	for _, cg := range colGroups { // within each column group
		rows := getRowNumbers(cg)              // just the row numbers
		groups += countConsecutiveGroups(rows) // number of consecutive groups
	}

	// Count top sides
	edge = getEdgePoints(-1, 0)    // all points on left side
	rowGroups := groupRows(edge)   // group by columns
	for _, rg := range rowGroups { // within each column group
		cols := getColNumbers(rg)              // just the row numbers
		groups += countConsecutiveGroups(cols) // number of consecutive groups
	}

	// Count bottom sides
	edge = getEdgePoints(1, 0)     // all points on left side
	rowGroups = groupRows(edge)    // group by columns
	for _, rg := range rowGroups { // within each column group
		cols := getColNumbers(rg)              // just the row numbers
		groups += countConsecutiveGroups(cols) // number of consecutive groups
	}

	return groups
}

// Extract edge cells from a list, i.e., just those that have something
// different (including nothing) in the direction indicated
func getEdgePoints(dr, dc int) []Point {
	var edge []Point
	for _, p := range points {
		if at(Point{p.r + dr, p.c + dc}) != at(p) {
			edge = append(edge, p)
		}
	}
	return edge
}

// Group list of points by columns, i.e., break into separate sublists
func groupCols(pp []Point) [][]Point {
	cols := [][]Point{}
	colIndex := map[int]int{} // col-> index
	for _, p := range pp {
		i, ok := colIndex[p.c]
		if ok {
			cols[i] = append(cols[i], p)
		} else {
			cols = append(cols, []Point{p})
			colIndex[p.c] = len(cols) - 1
		}
	}
	return cols
}

// Group list of points by rows, i.e., break into separate sublists
func groupRows(pp []Point) [][]Point {
	rows := [][]Point{}
	rowIndex := map[int]int{} // row-> index
	for _, p := range pp {
		i, ok := rowIndex[p.r]
		if ok {
			rows[i] = append(rows[i], p)
		} else {
			rows = append(rows, []Point{p})
			rowIndex[p.r] = len(rows) - 1
		}
	}
	return rows
}

// Get row numbers from a list of points
func getRowNumbers(pp []Point) []int {
	rr := []int{}
	for _, p := range pp {
		rr = append(rr, p.r)
	}
	return rr
}

// Get column numbers from a list of points
func getColNumbers(pp []Point) []int {
	cc := []int{}
	for _, p := range pp {
		cc = append(cc, p.c)
	}
	return cc
}

// Count the number of sequential groups in a list of numbers, by sorting
// the list and then counting how many times there is a difference of more than one
func countConsecutiveGroups(nums []int) int {
	var groups int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1]+1 {
			groups++
		}
	}
	return groups
}

// Get list of locations that are adjacent to a point, i.e., the
// coordinates for left/right/up/down
func neighbours(p Point) []Point {
	return []Point{Point{p.r - 1, p.c}, Point{p.r + 1, p.c},
		Point{p.r, p.c - 1}, Point{p.r, p.c + 1}}
}

// Char at location, '.' if out of range
func at(p Point) byte {
	if p.r < 0 || p.r >= len(rows) || p.c < 0 || p.c >= len(rows[0]) {
		return '.'
	} else {
		return rows[p.r][p.c]
	}
}
