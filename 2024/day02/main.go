// Advent of Code 2024, Day 02
//
// Given rows of numbers, check each for the following conditions:
// all numbers are increasing/decreasing, and by 1 to 3 each step.
// For Part 1, count up how many rows meet the condition. For Part 2,
// count up how many rows would meet the condition, if any digit were
// removed.
//
// AK, 02 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read the input file into rows
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")

	// Check each row, and count up passes
	var part1, part2 int // initialized to zero
	for _, r := range rows {

		// Parse list of numbers
		nums := []int{}
		parts := strings.Split(r, " ")
		for _, n := range parts {
			nums = append(nums, parseInt(n))
		}

		// Part 1: count up how many rows meet checks as-is
		part1 += check(nums)

		// Part 2: check if the row would pass, if any item were removed
		part2 += check2(nums)
	}

	fmt.Println("Part 1 (s/b 606):", part1)
	fmt.Println("Part 2 (s/b 644):", part2)
}

// For Part 2, try removing each number in the list, and check if passes checks
func check2(row []int) int {

	// Consider removing each element of list
	for i := 0; i < len(row); i++ {

		// Create a copy of the list with this element removed
		row2 := []int{}
		for j := 0; j < len(row); j++ {
			if j != i {
				row2 = append(row2, row[j])
			}
		}

		// Check this list, return success if ok
		if check(row2) == 1 {
			return 1
		}
	}

	// Failed with all attempts
	return 0
}

// Check a row: numbers must be all increasing or all decreasing,
// by between 1, 2 or 3 steps each time
func check(row []int) int {

	// Assume a rows with less than 2 elements is not valid, since it can't
	// be increasing/decreasing
	if len(row) < 2 {
		return 0
	}

	// Check the "direction" of the first to elements of this row, must be
	// incrasing or decreasing
	direction := sign(row[0], row[1])
	if direction == 0 {
		return 0
	}

	// Check this row: must be either all increasing or all decreasing,
	// always in the same direction
	for i := 1; i < len(row); i++ {
		delta := abs(row[i] - row[i-1])
		if sign(row[i-1], row[i]) != direction || delta < 1 || delta > 3 {
			return 0
		}
	}

	// Passed all checks
	return 1
}

// Parse an integer
func parseInt(s string) int {
	n, _ := strconv.Atoi(s) // TODO: handle error
	return n
}

// Absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

// Return sign of difference, i.e., 1 if b>a, -1 if b<a, 0 if the same
func sign(a, b int) int {
	if b > a {
		return 1
	} else if b < a {
		return -1
	} else {
		return 0
	}
}
