// Advent of Code 2024, Day 19
//
// Given a list of short patterns of letters ("stripes on a towel"), check a
// series of long patterns ("designs") to see if they can be composed of any
// combination of the short patterns.  For Part 1, determine how many of the
// designs can be made up of patterns. For Part 2, how many combinations in
// total. Did this with simple recursion, cycling through patterns that match
// head of a string, then doing same for the tail. For the problem input, had
// to add memoization of sub-designs already encountered as was taking too
// long.
//
// AK, 19 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Global list of patterns and memoization of subdesigns encountered
var patterns []string
var known map[string]int

func main() {

	// Read the input file and split into list of available stripe patterns,
	// and list of desired "designs" to be built from these
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")
	patterns = strings.Split(rows[0], ", ") // global
	designs := rows[2:]
	fmt.Printf("%d patterns, %d designs\n", len(patterns), len(designs))

	// For memoization of sub-designs encountered
	known = map[string]int{}

	// Check each design to see if it's possible to compose it from the
	// available patterns, and how many ways
	var part1, part2 int
	for _, des := range designs {
		n := canCompose(des, 0)
		if n > 0 { // part 1: count how many possible
			part1++
		}
		part2 += n // part 2: count how many ways
	}
	fmt.Println("Part 1 =", part1, ", Part 2 =", part2) // 260, 639963796864990
}

// Recursively calculate the number of ways a design can be composed
// using the available patterns?
func canCompose(des string, level int) int {

	// Return pre-computed result if already encountered (otherwise takes forever)
	k, ok := known[des]
	if ok {
		return k
	}

	// Empty string means we reached the end of a design, so a match
	if len(des) == 0 {
		return 1
	}

	// Otherwise recursively try every possible pattern
	n := 0 // number of combinations we will find
	for _, patt := range patterns {
		if strings.HasPrefix(des, patt) { // starts with pattern
			n += canCompose(des[len(patt):], level+1) // check rest
		}
	}

	// Save result for future use and return number of combinations found
	known[des] = n
	return n
}
