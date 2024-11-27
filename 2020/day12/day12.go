// Advent of Code 2020, Day 12
//
// Simulate movement of a "ship" based on simple instructions, directly
// for Part 1, relative to a "waypoint" for Part 2.
//
// AK, 14/01/2022 and 23/01/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read file and split into lines
	t, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(t), "\n")

	// Do parts 1 and 2
	part1(lines)
	part2(lines)
}

// Part 1: interpret instructions as simple movement of the ship N/E/S/W, or
// changing direction left/right by x degrees, or move forward in current
// direction
func part1(lines []string) {

	// Initial position is 0,0, and ship starts by facing east
	var x, y int64     // +x is east (right), +y is up (up)
	var dir int64 = 90 // 0 = north, 90 = east, 180 = south, 270 = west

	// Process each instruction
	// N/S/E/W : move north/south/east/weset by given value
	// L/R : turn left/right the given number of degrees.
	// F : move forward by the given value in the current direction
	for _, l := range lines {

		// Skip blank lines
		if len(l) == 0 {
			continue
		}

		// Parse instruction
		inst := l[0]
		amt, _ := strconv.ParseInt(l[1:], 10, 64)

		// Execute instruction
		if inst == 'N' {
			y += amt
		} else if inst == 'S' {
			y -= amt
		} else if inst == 'E' {
			x += amt
		} else if inst == 'W' {
			x -= amt
		} else if inst == 'L' {
			dir -= amt
			for dir < 0 {
				dir += 360
			}
		} else if inst == 'R' {
			dir += amt
			for dir >= 360 {
				dir -= 360
			}
		} else if inst == 'F' {
			if dir == 0 { // north (up)
				y += amt
			} else if dir == 90 { // east (right)
				x += amt
			} else if dir == 180 { // south (down)
				y -= amt
			} else if dir == 270 { // west (left)
				x -= amt
			} else {
				fmt.Printf("** Invalid direction %d\n", dir)
			}
		}

		fmt.Printf("%c %d -> %d,%d (dir %d)\n", inst, amt, x, y, dir)
	}

	dist := abs(x) + abs(y)
	fmt.Printf("Part 1: ending distance = %d\n", dist)
}

// Part 2: interpret instructions as movement of a waypoint,
// except F, which is movement of the ship towards the waypoint
// a number of times
func part2(lines []string) {

	// Initial position of the waypoint is 10 units east (right) and 1 unit
	// north (up), relative to the ship. East positions are positive X,
	// West positions are negative X, North positions are positive Y,
	// and South positions are negative Y.
	var wx int64 = 10 // Position of the waypoint (relative to ship)
	var wy int64 = 1  // (starts at 10 East, 1 North)
	var sx, sy int64  // Position of the ship (starts at 0,0)

	// Process each instruction
	// N/S/E/W : move the waypoint north/south/east/weset by given value,
	//   does not move the ship
	// L/R : rotate the waypoint around the ship left/right (counterclockwise
	//   and clockwise) the given number of degrees; moves the waypoint,
	//   but not the ship
	// F : move ship forward to the waypoint a number of times equal to the
	//   given value; each time, moves the ship the total distance between
	//   the ship and the waypoint, but does not move the waypoint (since
	//   it is always relative to the ship)
	for _, l := range lines {

		// Skip blank lines and parse instruction
		if len(l) == 0 {
			continue
		}
		inst := l[0]
		n, _ := strconv.ParseInt(l[1:], 10, 32)

		// Execute instructions
		if inst == 'N' { // Move just the waypoint
			wy += n
		} else if inst == 'S' {
			wy -= n
		} else if inst == 'E' {
			wx += n
		} else if inst == 'W' {
			wx -= n
		} else if inst == 'R' { // Rotate waypoint right around ship
			for i := n; i > 0; i -= 90 { // repeat for each 90 degrees
				wx_ := wx
				wx = wy
				wy = -wx_
			}
		} else if inst == 'L' { // Rotate waypoint left around ship
			for i := n; i > 0; i -= 90 { // repeat for each 90 degrees
				wx_ := wx
				wx = -wy
				wy = wx_
			}

		} else if inst == 'F' { // Move ship forward toward waypoint n times
			sx += n * wx
			sy += n * wy
		}

		fmt.Printf("%c %3d -> waypoint %d,%d ship %d,%d\n", inst, n, wx, wy, sx, sy)
	}

	// Manhattan distance, should be 286 for sample
	dist := abs(sx) + abs(sy)
	fmt.Printf("Part 2: ending distance = %d\n", dist)
}

// Simple absolute number
func abs(a int64) int64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
