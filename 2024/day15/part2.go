// Advent of Code 2024, Day 15, part 2
//
// Given a map of walls and movable "boxes",  and a list of up/down/left/right
// instructions, move a "robot" around, moving boxes as much as possible if
// they are in the way. For Part 2, the width of walls and boxes are doubled,
// but the robot remains width one, making movement tricker (since boxes can
// overlap). In both parts, answer is calculated from final position of all
// the boxes.
//
// AK, 15 & 17 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"maps"
)

type Point struct {
	x, y int
}

var obstacles map[Point]byte

func main() {

	// Read the input file into a list of byte vectors
	fname := "sample3.txt"
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

	// Turn floor into what's at each location, doubling width of
	// walls, spaces and obstacles for Part 2
	obstacles = map[Point]byte{}
	var pos Point // current position of '@'
	for y, r := range floor {
		x := 0
		for _, b := range r {
			if b == '@' { // current position, one blank afterward
				pos = Point{x, y}
				x++ // skip over position after @?
			} else if b == '#' || b == '.' { // double width of # and .
				obstacles[Point{x, y}] = b
				x++
				obstacles[Point{x, y}] = b
			} else if b == 'O' { // obstacles as []
				obstacles[Point{x, y}] = '['
				x++
				obstacles[Point{x, y}] = ']'
			}
			x++
		}
	}

	fmt.Println("Robot starts at", pos, string(at(pos)))
	printFloor(pos)

	// Start moving the robot
	for iter, instr := range instructions {

		// Set target location, skip move if hit a wall
		fmt.Println(iter+1, "Move", string(instr))

		// Move one position in direction where possible, shoving boxes out of the way
		if instr == '<' { // move left
			p1 := Point{pos.x - 1, pos.y} // target location is one left
			if at(p1) == ']' {            // if hit box(es), shove them left
				moveBoxesLeft(p1)
			}
			if at(p1) == '.' { // if free space (i.e., not box or wall), move
				pos = p1
			}
		} else if instr == '>' { // move right
			p1 := Point{pos.x + 1, pos.y} // target location is one right
			if at(p1) == '[' {            // if hit box(es), shove them right
				moveBoxesRight(p1)
			}
			if at(p1) == '.' { // if free space (i.e., not box or wall), move
				pos = p1
			}
		} else if instr == '^' { // move up
			p1 := Point{pos.x, pos.y - 1}       // target location is one up
			if at(p1) == '[' || at(p1) == ']' { // if hit box(es), shove them up
				moveBoxesUp(p1)
			}
			if at(p1) == '.' { // if free space (i.e., not box or wall), move
				pos = p1
			}
		} else if instr == 'v' { // move down
			p1 := Point{pos.x, pos.y + 1}       // target location is one down
			if at(p1) == '[' || at(p1) == ']' { // if hit box(es), shove them down
				moveBoxesDown(p1)
			}
			if at(p1) == '.' { // if free space (i.e., not box or wall), move
				pos = p1
			}
		}
	}

	// Answer is the sum of all "GPS coordinates"
	var ans int
	for p, v := range obstacles {
		if v == '[' { // look at left edges of all boxes
			ans += p.x + 100*p.y
		}
	}
	fmt.Println(ans) // Part 2: 1543141
}

// Move boxes one position to left
func moveBoxesLeft(p Point) {

	// You must be on right side of a box
	if at(p) != ']' {
		fmt.Printf("Error: moveBoxesLeft not on right side of box: %c\n", at(p))
		return
	}

	// Find the first space to the left, fail if none
	sp := Point{p.x, p.y}
	for at(sp) != '.' {
		if at(sp) == '#' { // hit wall, fail
			return
		}
		sp.x--
	}
	assert(at(sp) == '.', "Not on a space")

	// Now move everything to the right of the space one position left
	for sp.x <= p.x {
		obstacles[sp] = obstacles[Point{sp.x + 1, sp.y}]
		sp.x++
	}

	// Clear the original location
	delete(obstacles, p)
}

// Move boxes one position to right
func moveBoxesRight(p Point) {

	// You must be on left side of a box
	if at(p) != '[' {
		fmt.Printf("Error: moveBoxesRight not on left side of box: %c\n", at(p))
		return
	}

	// Find the first space to the right, fail if none
	sp := Point{p.x, p.y}
	for at(sp) != '.' {
		if at(sp) == '#' { // hit wall, fail
			return
		}
		sp.x++
	}
	assert(at(sp) == '.', "Not on a space")

	// Now move everything to the left of the space one position right
	for sp.x >= p.x {
		obstacles[sp] = obstacles[Point{sp.x - 1, sp.y}]
		sp.x--
	}

	// Clear the original location
	delete(obstacles, p)
}

// Move boxes up one position, including any that touch entirely or corners
func moveBoxesUp(p Point) {

	// You must be on a box
	if at(p) == ']' { // move to left side of box
		p.x--
	}
	if at(p) != '[' {
		fmt.Printf("Error: moveBoxesUp not on a box: %c\n", at(p))
		return
	}

	// If a wall above, fail
	upLeft := Point{p.x, p.y - 1}
	upRight := Point{p.x + 1, p.y - 1}
	if at(upLeft) == '#' || at(upRight) == '#' {
		return
	}

	// Preserve state, to restore it in case one of the moves fails
	before := maps.Clone(obstacles)

	// If there is a box above, try to move it
	if at(upLeft) == '[' || at(upLeft) == ']' { // box directly above or above-left
		moveBoxesUp(upLeft)
	}
	if at(upRight) == '[' { // box above-right
		moveBoxesUp(upRight)
	}

	// If free space above, move this box to there, otherwise restore state
	if at(upLeft) == '.' && at(upRight) == '.' {
		rside := Point{p.x + 1, p.y} // right side of the box being moved
		obstacles[upLeft] = obstacles[p]
		obstacles[upRight] = obstacles[rside]
		delete(obstacles, p)
		delete(obstacles, rside)
	} else {
		obstacles = before
	}
}

// Move boxes down one position, including any that touch entirely or corners
func moveBoxesDown(p Point) {

	// You must be on a box
	if at(p) == ']' { // move to left side of box
		p.x--
	}
	if at(p) != '[' {
		fmt.Printf("Error: moveBoxesDown not on a box: %c\n", at(p))
		return
	}

	// If moving box down stopped by wall, fail
	downLeft := Point{p.x, p.y + 1}
	downRight := Point{p.x + 1, p.y + 1}
	if at(downLeft) == '#' || at(downRight) == '#' {
		return
	}

	// Preserve state, to restore it in case one of the moves fails
	before := maps.Clone(obstacles)

	// If there is a box below, try to move it
	if at(downLeft) == '[' || at(downLeft) == ']' { // box directly below or below-left
		moveBoxesDown(downLeft)
	}
	if at(downRight) == '[' { // box below-right
		moveBoxesDown(downRight)
	}

	// If free space below, move this box to there, otherwise restore state
	if at(downLeft) == '.' && at(downRight) == '.' {
		rside := Point{p.x + 1, p.y} // right side of the box being moved
		obstacles[downLeft] = obstacles[p]
		obstacles[downRight] = obstacles[rside]
		delete(obstacles, p)
		delete(obstacles, rside)
	} else {
		obstacles = before
	}
}

// What's at position, period if nothing
func at(p Point) byte {
	b, ok := obstacles[p]
	if ok {
		return b
	} else {
		return '.'
	}
}

// Print a map of the floor, for debugging (set number of rows & cols accordingly)
func printFloor(pos Point) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 20; x++ {
			if x == pos.x && y == pos.y {
				fmt.Print("@")
			} else {
				p := Point{x, y}
				fmt.Print(string(at(p)))
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

// For debugging
func assert(cond bool, msg string) {
	if !cond {
		panic("Assert failed:" + msg)
	}
}
