// Advent of Code 2024, Day 03
//
// Given a string, find all embedded "mul(3,4)" instructions, and add up the
// results of the multiplications. For Part 2, "do()" and "don't()"
// instructions turn multiplication on/off.
//
// AK, 03 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// Read the input file into lines
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

	// Pattern to look for: mul(a,b) in Part 1, also do() and don't() in Part 2
	patt := `mul\([0-9]+,[0-9]+\)`
	patt = patt + `|do\(\)|don't\(\)`
	r := regexp.MustCompile(patt)

	// Process each row, finding the valid instructions, executing them,
	// and adding up values
	multOn := 1          // flag for whether multiplication is on
	var part1, part2 int // answers, initialized to zero
	for _, l := range lines {
		instr := r.FindAllString(l, -1) // find all instructions
		for _, i := range instr {       // each instruction
			if i == "do()" { // turn multiplication ON (Part 2)
				multOn = 1
			} else if i == "don't()" { // turn multiplication OFF
				multOn = 0
			} else { // do multiplication
				val := execute(i)     // get value, e.g., multiplication
				part1 += val          // always add for Part 1
				part2 += val * multOn // only when mult turned on for Part 2
			}
		}
	}

	// Show answers
	fmt.Println("Part 1 (s/b 182780583):", part1)
	fmt.Println("Part 2 (s/b 90772405):", part2)
}

// Execute an instruction, e.g., "mul(2,3)" -> 6
func execute(s string) int {

	// Pattern: instr(a,b)
	patt := `([a-z]+)\(([0-9]+),([0-9]+)\)`
	r := regexp.MustCompile(patt)

	// Extract the numbers to multiply and return result
	parts := r.FindStringSubmatch(s)
	arg1 := parseInt(parts[2]) // ignore parts[1], always "mul"
	arg2 := parseInt(parts[3])
	return arg1 * arg2 // do the multiplication
}

// Parse an integer
func parseInt(s string) int {
	n, _ := strconv.Atoi(s) // TODO: handle error
	return n
}
