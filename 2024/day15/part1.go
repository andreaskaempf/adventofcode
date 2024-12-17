// Advent of Code 2024, Day 15, Part 1
//
// Given a map of walls and movable "boxes", and a list of up/down/left/right
// instructions, move a "robot" around, moving boxes as much as possible if
// they are in the way. For Part 1, the walls, boxes, and robot are all of
// width 1. Answer is calculated from final position of all the boxes.
//
// AK, 15 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Point struct {
	x, y int
}

var obstacles map[Point]byte

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Split floor layout and instructions
	readingFloor := true
	floor := [][]byte{}
	instructions := []byte{}
	for _, r := range rows {
		if len(r) == 0 {
			readingFloor = false
		} else if readingFloor {
			floor = append(floor, r)
		} else {
			instructions = append(instructions, r...)
		}
	}

	// Turn floor into what's at each location
	obstacles = map[Point]byte{}
	var pos Point
	for y, r := range floor {
		for x, b := range r {
			p := Point{x, y}
			if b == '@' {
				pos = p
			} else if b != '.' {
				obstacles[p] = b
			}
		}
	}

	fmt.Println("Robot starts at", pos)

	// Start moving the robot
	for _, instr := range instructions {

		// Get new candidate location
		var dx, dy int
		if instr == '<' { // move left
			dx = -1
		} else if instr == '>' { // move right
			dx = 1
		} else if instr == '^' { // move up
			dy = -1
		} else if instr == 'v' { // move down
			dy = 1
		} else {
			fmt.Println("Bad instruction")
		}
		p1 := Point{pos.x + dx, pos.y + dy}

		// If would hit a wall, don't do anything
		if at(p1) == '#' {
			continue
		}

		// If next location is an obstacle, try to move it and any others
		if at(p1) == 'O' {
			clearObstacles(p1, dx, dy)
		}

		// If now free, move to new position
		if at(p1) == ' ' {
			pos = p1
		}
	}

	// For Part 1, calculate the sum of all GPS coordinates
	var part1 int
	for p, v := range obstacles {
		if v == 'O' { // ignore spaces, which we assigned to delete boxes when moving them
			part1 += p.x + 100*p.y
		}
	}
	fmt.Println("Part 1 =", part1) // 1505963

}

// Recursively move any obstacles from this location, in the designated direction
func clearObstacles(p Point, dx, dy int) {

	// To clear this obstacle, need to move it in given direction
	p1 := Point{p.x + dx, p.y + dy}

	// If that location is a wall, can't do it
	if at(p1) == '#' {
		return
	}

	// If that location is another obstacle, first try to move it
	if at(p1) == 'O' {
		clearObstacles(p1, dx, dy)
	}

	// If obstacle(s) cleared, move the point
	if at(p1) == ' ' {
		obstacles[p] = ' '
		obstacles[p1] = 'O'
	}
}

// What's at position, space if nothing
func at(p Point) byte {
	b, ok := obstacles[p]
	if ok {
		return b
	} else {
		return ' '
	}
}

// Print a map of the floor
func printFloor(pos Point) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 12; x++ {
			p := Point{x, y}
			if p == pos {
				fmt.Print("@")
			} else {
				fmt.Print(string(at(p)))
			}
		}
		fmt.Println("")
	}
}
