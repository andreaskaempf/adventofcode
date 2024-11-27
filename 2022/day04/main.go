// Advent of Code 2022, Day 04
//
// Given pairs of numeric ranges, in how many pairs is one range entirely
// contained within another (Part 1), and how many pairs overlap at all
// (Part 2).
//
// AK, 4 Dec 2022

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Read the input file
	//lines := readLines("sample.txt")
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Each line consists of two ranges, split these
	var part1, part2 int
	for _, l := range lines {

		// Parse pair of ranges
		ranges := strings.Split(l, ",")
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")
		r1 := []int{atoi(range1[0]), atoi(range1[1])}
		r2 := []int{atoi(range2[0]), atoi(range2[1])}

		// // Part 1: In how many assignment pairs does one range fully contain the other?
		if (r1[0] >= r2[0] && r1[1] <= r2[1]) || (r2[0] >= r1[0] && r2[1] <= r1[1]) {
			part1++
		}

		// Part 2:how many pairs overlap at all?
		if (r1[0] >= r2[0] && r1[0] <= r2[1]) || (r1[1] >= r2[0] && r1[1] <= r2[1]) ||
			(r2[0] >= r1[0] && r2[0] <= r1[1]) || (r2[1] >= r1[0] && r2[1] <= r1[1]) {
			part2++
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
