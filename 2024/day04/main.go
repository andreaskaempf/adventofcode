// Advent of Code 2024, Day 04
//
// Find all instances in XMAS in a matrix of text in any direction (part 1)
// or MAS cross pattern (part 2).
//
// AK, 04 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"

	"bytes"
)

// Global variable with matrix of characters
var rows [][]byte

func main() {

	// Read the input file into rows of bytes
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte{'\n'})

	// Get answers
	fmt.Printf("Part 1 = %d, part 2 = %d\n", findall1(), findall2())
}

// Part 1: find and count all instances of XMAS in any direction
func findall1() int {

	var n int
	xmas := []byte{'X', 'M', 'A', 'S'}
	dirs := []int{-1, 0, 1}

	// Search in each direction from each x/y position, by setting
	// dx and dy to -1, 0 or +1, and count up finds
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[r]); c++ {
			for _, dx := range dirs {
				for _, dy := range dirs {
					n += match(xmas, r, c, dx, dy)
				}
			}
		}
	}
	return n
}

// Part 2: find and count all M-A-S crosses
func findall2() int {

	var n int

	// Search in each direction from each x/y position, count up finds
	// w x
	//  A
	// y z
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(rows[r]); c++ {
			a := at(r, c)
			w := at(r-1, c-1)
			x := at(r-1, c+1)
			y := at(r+1, c-1)
			z := at(r+1, c+1)
			if a == 'A' && ((w == 'M' && z == 'S') || (w == 'S' && z == 'M')) && ((x == 'M' && y == 'S') || (x == 'S' && y == 'M')) {
				n += 1
			}
		}
	}

	return n
}

// Check if sequence found at position, with dx/dx indicating direction,
// returns 1 or 0 for easy adding
func match(bb []byte, r, c, dx, dy int) int {

	// Can't find if no direction to search in
	if dx == 0 && dy == 0 {
		return 0
	}

	// Search from the starting position, moving in the direction indicated
	// by dx/dy (can be horizonal, vertical, or diagonal, either forwards
	// or backwards)
	for _, b := range bb {
		if at(r, c) != b {
			return 0
		}
		r += dy
		c += dx
	}
	return 1
}

// Get char at position, ' ' if out of bounds
func at(r, c int) byte {
	if r >= 0 && r < len(rows) && c >= 0 && c < len(rows[r]) {
		return rows[r][c]
	} else {
		return ' '
	}
}
