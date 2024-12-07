// Advent of Code 2024, Day 07
//
// Given a number of target: n n n lines, add up all the targets for which
// the n values can be computed left-to-right using some combination of + or *
// operators. For Part 2, there is third || operator, which string-concatenates
// the operands. There is no operator precedence, just left-to-right. Key to
// this was coming up with all the combinations of 2 or 3 values across any
// number of columns. For Part 1, I did this with binary arithmetic, but had
// to code an arbitrary-base counting function for Part 2.
//
// AK, 07 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Set this to 0 for Part 1, or 1 for Part 2
	part2 := 1

	// Read the input file
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")

	// Process each row
	ans := 0
	for _, l := range rows {

		// Extract target value and operands
		parts := strings.Split(l, " ")
		target := atoi(strings.TrimSuffix(parts[0], ":"))
		nums := parts[1:] // leave as strings

		// Get all possible combinations of operators, two ops for Part 1
		// and three ops for Part 2
		ncols := len(nums) - 1                // number of operators
		combs := combinations(ncols, 2+part2) // 2 for Part 1, 3 for Part 2

		// Check if any combination of operations yield target
		for _, ops := range combs {

			// Create equation
			eq := strings.TrimSpace(nums[0])
			for i := 1; i < len(nums); i++ {
				op := ops[i-1]
				if op == 0 {
					eq += " + "
				} else if op == 1 {
					eq += " * "
				} else {
					eq += " || "
				}
				eq += strings.TrimSpace(nums[i])
			}

			// Evaluate it, add to answer and stop if okay
			y := evaluate(eq)
			if y == target {
				ans += y
				break

			}
		}
	}

	// Sample 3749 / 11387", data 20665830408335 / 354060705047464
	fmt.Println(ans)
}

// Evaluate left-to-right an equation of numbers interspersed with +, * or ||
// operators, where || is string concatenation (not addition)
func evaluate(eq string) int {

	// The numbers and operators must be surrounded by spaces
	parts := strings.Split(eq, " ")
	y := atoi(parts[0]) // the first number, parse string to number
	i := 2              // position in equation, skip to second number
	for i < len(parts) {
		op := parts[i-1]    // operator is before the next number
		n := atoi(parts[i]) // the next number
		if op == "+" {      // add
			y += n
		} else if op == "*" { // multiply
			y *= n
		} else if op == "||" { // string concatenate
			y = atoi(fmt.Sprintf("%d%d", y, n))
		} else {
			panic("Invalid operator " + op)
		}
		i += 2
	}

	// Return result
	return y
}

// Generate all possible lists of length 'cols', containing numbers up to
// the base. For example, if base is 2, will generate lists like [0, 1, 0].
// This is used to generate all possible combinations of operators between
// the numbers.
func combinations(cols int, base int) [][]int {

	// Initialize list of combinations
	res := [][]int{}

	// Initialize a list of numbers
	counters := make([]int, cols, cols)

	// Count up and store each value as a list in binary digits
	for {
		// Save this combination
		res = append(res, copynums(counters))

		// Increase the last digit
		col := len(counters) - 1
		counters[col] += 1

		// If overflow, carry left
		for col > 0 && counters[col] >= base {
			counters[col-1] += 1
			counters[col] = 0
			col -= 1
		}

		// Finished
		if counters[0] >= base {
			break
		}
	}

	return res
}

// Make a copy of a list of numbers
func copynums(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	return res
}

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return n
}
