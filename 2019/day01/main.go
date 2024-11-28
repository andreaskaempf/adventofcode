// Advent of Code 2019, Day 01
//
// Create a simple function to convert "mass" to "fuel" required (both
// integers). For Part 1, add up the fuel obtained from a list of mass values.
// For Part 2, also consider the fuel required for the fuel, iterating.
//
// AK, 28 Nov 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Process each row
	var part1, part2 int64
	for _, l := range rows {

		// Part 1: add up the fuel requirements for each mass in the input
		mass, _ := strconv.ParseInt(string(l), 10, 64)
		fuel := mass_to_fuel(mass)
		part1 += fuel

		// Part 2: also consider the fuel required to carry the required fuel,
		// iteratively
		part2 += fuel
		fuel = mass_to_fuel(fuel)
		for fuel > 0 {
			part2 += fuel
			fuel = mass_to_fuel(fuel)
		}
	}

	fmt.Println("Part 1 (s/b 3 390 596) =", part1)
	fmt.Println("Part 2 (s/b 5 083 024) =", part2)
}

// To find the fuel required for a module, take its mass, divide by three,
// round down, and subtract 2.
func mass_to_fuel(n int64) int64 {
	return n/3 - 2
}
