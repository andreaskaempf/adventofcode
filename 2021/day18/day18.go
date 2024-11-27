// Advent of Code 2021, Day 18
//
// Calculate operations on nested lists of numbers, first
// "adding" (concatenating them), then reducing them according
// to arcane rules. Part 1: add up each line in file, reducing
// after each line, and report the final "magnitude" (calculated
// by recursively collapsing inner pairs of numbers). For Part 2,
// find the pair of input lines which, when added together and
// "reduced", yield the highest magnitude.
//
// AK, 18/12/2021

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// An expression is just an array of numbers, with ASCII values for
// brackets, or numbers. Need to use int rather than just bytes
// because we reduce them and add up at the end to calculate
// "magnitude" and bytes aren't big enough.
type Expression []int

const LB = 91
const RB = 93

// Parse a string into a flat list of tokens. Assumes numbers
// are one or two digit (only single digit for input data, but double
// digits needed for tests), and commas are ignored.
func parse(s string) Expression {

	res := Expression{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '[' {
			res = append(res, LB)
		} else if c == ']' {
			res = append(res, RB)
		} else if isdigit(s[i]) {
			n := int(s[i] - '0') // convert this digit to a number
			if isdigit(s[i+1]) { // two digits?
				n = n*10 + int(s[i+1]-'0')
				i++ // skip the second digit
			}
			res = append(res, n)
		}
	}
	return res
}

// Print a tokenized expression
func printExpr(expr Expression) {
	for _, b := range expr {
		if b == LB {
			fmt.Print("[ ")
		} else if b == RB {
			fmt.Print("] ")
		} else {
			fmt.Print(b, " ")
		}
	}
	fmt.Print("\n")
}

// "Add" two tokenized expressions, but putting both inside a list
func add(a Expression, b Expression) Expression {
	res := Expression{}
	res = append(res, LB)
	for _, c := range a {
		res = append(res, c)
	}
	for _, c := range b {
		res = append(res, c)
	}
	res = append(res, RB)
	return res
}

// "Explode" a pair:
// 1. Find the first pair of numbers that is nested 4 (or more?) deep
// 2. Add the left value of that pair the the next number to its left (if any)
// 3. Add the right value of that pair to the next number on its right (if any)
// 4. Replace the pair (including brackets) with the number zero
// Returns the (possibly changed) expression, and true/false indicating
// whether a change was made.
func explode(expr Expression) (Expression, bool) {

	// Find the first pair of numbers that is within 4 outer pairs
	level := 0
	pair := -1 // location of explodable pair if found
	for i, b := range expr {
		if level == 4 && b == LB && isnumber(expr[i+1]) && isnumber(expr[i+2]) && expr[i+3] == RB {
			pair = i
			break
		} else if b == LB {
			level += 1
		} else if b == RB {
			level -= 1
		}
	}

	// If no pair found, stop here
	if pair == -1 {
		return expr, false
	}

	// Find the first number to the left and add left side of pair to it
	for j := pair - 1; j >= 0; j-- {
		if isnumber(expr[j]) {
			expr[j] += expr[pair+1]
			break
		}

	}

	// Find the first number to the right of the pair, and add the right
	// number of the pair to it
	for j := pair + 4; j < len(expr); j++ { // first token after the pair found
		if isnumber(expr[j]) {
			expr[j] += expr[pair+2] // add the second number of the pair
			break
		}
	}

	// Now replace the four characters of the original pair with a zero
	expr[pair] = 0
	res := Expression{}
	for i := 0; i < len(expr); i++ {
		if i <= pair || i > pair+3 {
			res = append(res, expr[i])
		}
	}
	return res, true
}

// Go through expression, find any number >= 10, and split it. Returns
// possibly change expression, and true/false if changed or not.
func splitFirst(expr Expression) (Expression, bool) {

	// Find first number >= 10
	ten := -1
	for i := 0; i < len(expr); i++ {
		if isnumber(expr[i]) && expr[i] >= 10 {
			ten = i
			break
		}
	}

	// Return unchanged if none found
	if ten == -1 {
		//fmt.Println("No >= 10 found")
		return expr, false
	}

	// Split the number into [a b]
	sexpr := splitNum(expr[ten])

	// Insert the expression in place of the original number
	res := Expression{}
	res = append(res, expr[:ten]...)
	res = append(res, sexpr...)
	res = append(res, expr[ten+1:]...)

	// Return the changed expression
	return res, true
}

// Reduce a list, by successively exploding and splitting
func reduce(expr Expression) Expression {

	// Repeat until no more changes
	var res Expression
	res = append(res, expr...)
	done := false
	for !done {

		// Try explode
		res1, exploded := explode(res)
		if exploded {
			//fmt.Print("After explode: ")
			//printExpr(res1)
			res = nil
			res = append(res, res1...)
			continue
		}

		// Otherwise, try split
		res2, splitted := splitFirst(res)
		if splitted {
			//fmt.Print("After split: ")
			//printExpr(res2)
			res = nil
			res = append(res, res2...)
			continue
		}

		// Otherwise done
		done = true
	}

	// Return result of reduction
	//fmt.Print("Final result: ")
	//printExpr(res)
	return res
}

// "Split" a number: replace it with a pair; the left element of the pair
// should be the regular number divided by two and rounded down, while the right
// element of the pair should be the regular number divided by two and rounded up.
// For example, 10 becomes [5,5], 11 becomes [5,6], 12 becomes [6,6], and so on.
func splitNum(n int) Expression {
	left := int(math.Floor(float64(n) / 2))
	right := int(math.Ceil(float64(n) / 2))
	return Expression{LB, left, right, RB}
}

// If a token is not a bracket, it's a number
func isnumber_(b byte) bool {
	return b != LB && b != RB
}
func isnumber(b int) bool {
	return b != LB && b != RB
}

// Determine character is a digit
func isdigit(c byte) bool {
	//return c >= 48 && c <= 57
	return c >= '0' && c <= '9'
}

// Add up lines of a file and return final expression, nil if error
func addUpFile(fname string) Expression {

	// Open file
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Could not open ", fname)
		return nil
	}

	// The current result
	var last Expression

	// Read each line of input file
	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {

		// Read line
		t := scanner.Text()
		lineNo += 1
		//fmt.Printf("\nLine %d: %s\n", lineNo, t)

		// Parse the line (does not need to be reduced)
		expr := parse(t)
		//fmt.Print("Parsed: ")
		//printExpr(expr)

		// If not first line, add it to previous and reduce
		if lineNo > 1 {
			last = add(last, expr)
			last = reduce(last)
			//fmt.Print("Added and reduced: ")
			//printExpr(last)
		} else {
			last = expr
		}
	}

	// Return final expression
	return last
}

// To check whether it's the right answer, the snailfish teacher only checks
// the magnitude of the final sum. The magnitude of a pair is 3 times the
// magnitude of its left element plus 2 times the magnitude of its right
// element. The magnitude of a regular number is just that number.
//
// For example, the magnitude of [9,1] is 3*9 + 2*1 = 29;
// the magnitude of [1,9] is 3*1 + 2*9 = 21.
// Magnitude calculations are recursive:
// the magnitude of [[9,1],[1,9]] is 3*29 + 2*21 = 129.
//
// Here are a few more magnitude examples:
//
// [[1,2],[[3,4],5]] becomes 143.
// [[[[0,7],4],[[7,8],[6,0]]],[8,1]] becomes 1384.
// [[[[1,1],[2,2]],[3,3]],[4,4]] becomes 445.
// [[[[3,0],[5,3]],[4,4]],[5,5]] becomes 791.
// [[[[5,0],[7,4]],[5,5]],[6,6]] becomes 1137.
// [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]] becomes 3488.

func magnitude(expr Expression) int {

	// The magnitude of a pair is 3 times the magnitude of its left element
	// plus 2 times the magnitude of its right element. The magnitude of a
	// regular number is just that number.

	// Find inner pairs, and reduce them to a number, repeating
	// until no more pairs
	for true {

		// Find the first pair of numbers
		pair := -1 // location of pair if found
		for i := 0; i < len(expr); i++ {
			b := expr[i]
			if len(expr)-i >= 4 && b == LB && isnumber(expr[i+1]) && isnumber(expr[i+2]) && expr[i+3] == RB {
				pair = i
				break
			}
		}

		// If a pair found, replace it with magnitude calculation
		if pair < 0 { // if no pair found, stop here
			break
		} else {

			// Calculate the magnitude of this pair
			left := expr[pair+1]
			right := expr[pair+2]
			mag := left*3 + right*2

			// Put it in the first position of the original pair
			expr[pair] = mag

			// Remove the rest of the original pair from the expression
			res := Expression{}
			for i := 0; i < len(expr); i++ {
				if i <= pair || i > pair+3 {
					res = append(res, expr[i])
				}
			}
			expr = res
		}
	}

	//fmt.Print("After calculating magnitude: ")
	//printExpr(expr)

	return int(expr[0])
}

// For part 2, find the maximum magnitude from adding any two inputs in the file
func part2(fname string) {

	// Open file
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Could not open ", fname)
		return
	}

	// Read and parse lines into an array
	lines := []Expression{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		expr := parse(t)
		lines = append(lines, expr)
	}

	// Calculate all possibilities and find highest magnitude
	best := 0
	for _, x1 := range lines {
		for _, x2 := range lines {
			sum := add(x1, x2)
			sum = reduce(sum)
			mag := magnitude(sum)
			if mag > best {
				best = mag
			}
		}
	}

	// Show best found
	fmt.Println("Part 2: Best pairwise magnitude found =", best)
}

func main() {

	// Part 1: Calculate the magnitude of the input file after "adding" and
	// reducing all lines
	fname := "input.txt" // "sample2.txt"
	expr := addUpFile(fname)
	mag := magnitude(expr)
	fmt.Println("Part 1: Magnitude =", mag)

	// Part 2: add up all pairs, find highest magnitude
	part2(fname)
}
