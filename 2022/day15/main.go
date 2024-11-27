// Advent of Code 2022, Day 15
//
// Given a list of "sensors" and their distance to nearest "beacon",
// find positions in a row that could not possibly have a beacon (Part 1),
// and the possible location of an undetected beacon (i.e., where there is
// in coverage by known beacons) for Part 2.
//
// AK, 15 Dec 2022

package main

import (
	"fmt"
	"sort"
	"strings"
)

// A position in 2-d space
type Position struct {
	x, y int
}

// A sensor
type Sensor struct {
	at     Position // current location
	beacon *Beacon  // nearest beacon
}

// A beacon, just a position
type Beacon struct {
	at Position
}

// A range, just two numbers
type Range struct {
	lo, hi int
}

// Global list of sensors, map of beacons
var sensors []Sensor
var beacons map[Position]Beacon

func main() {

	// Read the input file, parse into list of sensors
	fname := "sample.txt"
	fname = "input.txt" // uncomment to use input
	readData(fname)

	// Part 1: use 10 for sample, 2000000 for input
	//part1(10)
	part1(2000000)

	// Part 2: use 20 for sample, 4000000 for input
	//part2(20)
	part2(4000000)
}

// Part 1: count the positions where a beacon cannot possibly be along
// just a single row.
func part1(y int) { // y is the row number

	// Create a map of all the points covered, i.e., for each sensor,
	// the points within the distance to its nearest beacon. For Part 1,
	// we are only looking at one row
	covered := map[Position]int{}
	for _, s := range sensors {
		d := dist(s.at, s.beacon.at) // Distance from this sensor to nearest beacon
		for x := s.at.x - d - 1; x <= s.at.x+d+1; x++ {
			p := Position{x, y}
			if dist(p, s.at) <= d {
				covered[p] = 1
			}
		}
	}
	//visualize(covered)

	// Count the positions where a beacon cannot possibly be along
	// just a single row
	n := 0
	for p, _ := range covered {
		if p.y == y && !beaconAt(p) {
			n++
		}
	}
	fmt.Println("Part 1 (s/b 26, 4827924):", n)
}

// Find an undetected beacon, i.e., outside the space we already identified as
// not having a beacon, within a constrained space (x & y both 0-20 for sample,
// 0 to 4000000 for input. Compute "tuning frequency" as = x * 4000000 + y.
// In sample, only 14,11 could have beacon, freq = 56000011
func part2(maxXY int) {

	// Get ranges that are covered in each row, by computing distance from
	// sensor to its beacon, and then tracing a diamond that covers the same
	// distance from the sensor in all directions.
	fmt.Println("Building coverage")
	covered := map[int][]Range{} // row =>  [Range1, ...]
	for _, s := range sensors {
		d := dist(s.at, s.beacon.at)              // distance from sensor to nearest beacon
		var w int                                 // width at tip of diamond
		for y := s.at.y - d; y <= s.at.y+d; y++ { // each row of diamond

			// Add a range for this row
			if y >= 0 && y <= maxXY {
				r := Range{s.at.x - w, s.at.x + w}
				covered[y] = append(covered[y], r)
			}

			// Adjust width of the diamond
			if y < s.at.y {
				w++
			} else if y >= s.at.y {
				w--
			}
		}
	}

	// Now search each row for a gap between coverage
	fmt.Println("Searching for gap (s/b 14,11)")
	var gapX, gapY int
	for r := 0; r <= maxXY; r++ {

		// Get the ranges for this row, and merge overlapping
		lims := merge(covered[r]) // returns sorted

		// Skip where only one range and covers entire row
		if len(lims) == 1 && lims[0].lo <= 0 && lims[0].hi >= maxXY {
			continue // no gaps possible
		}

		// Gaps to left of first range (ignored for problem)?
		if lims[0].lo > 0 {
			fmt.Println("Row", r, ": gaps up to", lims[0].lo-1)
		}

		// Gaps between ranges?
		// Note that we only consider gaps in the middle to be relevant
		// for the problem.
		for i := 1; i < len(lims); i++ {
			if lims[i].lo-lims[i-1].hi > 1 {
				fmt.Println("Row", r, ": gap from", lims[i-1].hi+1, "to", lims[i].lo-1)
				gapX = lims[i-1].hi + 1
				gapY = r
				break
			}
		}

		// Gaps after last range (ignored for problem)?
		if lims[len(lims)-1].hi < maxXY {
			fmt.Println("Row", r, ": gaps after", lims[len(lims)-1].hi)
		}
	}

	// Compute "frequency" for part 2 answer
	freq := gapX*4000000 + gapY
	fmt.Printf("Part 2 (s/b 56000011): gap at %d,%d => freq %d\n", gapX, gapY, freq)
}

// Merge overlapping ranges
func merge(ranges []Range) []Range {

	// Sort the ranges by starting x
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].lo < ranges[j].lo
	})

	// Combine ranges that overlap, or add ones that don't
	result := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		r := ranges[i]
		last := result[len(result)-1]
		if r.lo <= last.hi {
			if r.hi >= last.hi { // extend last range
				result[len(result)-1].hi = r.hi
			}
		} else {
			result = append(result, r)
		}
	}
	return result
}

// Read the input file, parse into lists of sensors and beacons
func readData(fname string) {

	beacons = map[Position]Beacon{}
	for _, l := range readLines(fname) {

		// Extract the x and y positions of the sensor and beacon
		words := strings.Split(l, " ")
		sx := words[2][2:]
		sy := words[3][2:]
		spos := Position{atoi(sx[:len(sx)-1]), atoi(sy[:len(sy)-1])}
		bx := words[8][2:]
		by := words[9][2:]
		bpos := Position{atoi(bx[:len(bx)-1]), atoi(by)}

		// Get the beacon, create if necessary
		b, ok := beacons[bpos]
		if !ok {
			b = Beacon{bpos}
			beacons[bpos] = b
		}

		// Add sensor to list
		sensors = append(sensors, Sensor{spos, &b})
	}
}

// Manhattan distance between two positions
func dist(a, b Position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

// Is a sensor located at position?
func sensorAt(p Position) bool {
	for _, s := range sensors {
		if s.at.x == p.x && s.at.y == p.y {
			return true
		}
	}
	return false
}

// Is a beacon located at position?
func beaconAt(p Position) bool {
	_, ok := beacons[p]
	return ok
}

// Visualize the covered map, used for debugging so not generalized
func visualize(covered map[Position]int) {

	x := make([]byte, 50, 50)  // adjust size of row as necessary
	for y := -2; y < 23; y++ { // adjust range as necessary

		// Blank out this row
		for i := 0; i < len(x); i++ {
			x[i] = '.'
		}

		// Mark any covered areas
		xoff := 3 // x offset, adjust as necessary to avoid negative indexes
		for p, _ := range covered {
			if p.x < 0 || p.x > 20 || p.y < 0 || p.y > 20 { // adjust as nec
				continue // restrict for part 2
			}
			if p.y == y {
				if beaconAt(p) {
					x[p.x+xoff] = 'B'
				} else if sensorAt(p) {
					x[p.x+xoff] = 'S'
				} else {
					x[p.x+xoff] = '#'
				}
			}
		}
		fmt.Printf("%2d ", y)
		fmt.Println(string(x))
	}
}
