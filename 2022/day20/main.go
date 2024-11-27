// Advent of Code 2022, Day 20
//
// Given a list of numbers (7 in sample, but 5000 in input), simulate moving
// each number forward or backward in the (circular) list, forward if positive
// or backward if negative. For Part 1, do this once, and report the sum of the
// values 1000, 2000, and 3000 after zero. For Part 2, multiply each number
// by a huge value, and do it 10 times, report same sum. Complicated by
// duplicate values in the  main input, so you can't just look for position
// of a value. Also, iterations in Part 2 are infeasible with large multiplier
// as well as 10 iterations.
//
// AK, 20 Dec 2022

package main

import (
	"fmt"
)

// Number, with its original position (to get around duplicates)
type Number struct {
	pos   int
	value int64
}

func main() {

	var part2 bool
	part2 = true // uncomment this line to run part 2

	// Read list of numbers from input file
	fname := "sample.txt"
	fname = "input.txt"    // uncomment this line to use problem input
	var seq, nums []Number //  original and reordered lists
	for i, l := range readLines(fname) {
		n := Number{i, atoi64(l)}
		if part2 {
			n.value *= 811589153
		}
		nums = append(nums, n)
		seq = append(seq, n)
	}
	fmt.Println(len(seq), "numbers")

	// Process each number in original sequence, 10 times in part 2
	iterations := ifElse(part2, 10, 1)
	for iter := 0; iter < iterations; iter++ {
		fmt.Println("Iteration", iter)
		for _, n := range nums {
			seq = move(seq, n)
		}
	}

	// Get answer: sum of numbers at positions 1000, 2000, 3000
	zero := findValue(Number{-1, 0}, seq)
	n1000 := seq[(zero+1000)%len(seq)].value
	n2000 := seq[(zero+2000)%len(seq)].value
	n3000 := seq[(zero+3000)%len(seq)].value
	ans := n1000 + n2000 + n3000
	fmt.Printf("Answer (s/b 3, 11703 for Part 1, 1623178306 for Part 2): %d\n", ans)
}

// Move given number x by same number of positions (negative for left)
func move(seq []Number, x Number) []Number {

	// Number of moves to make: take remainder to avoid too many iterations
	// (not sure why you need -1, but it works)
	n := abs(x.value) % int64(len(seq)-1)

	// Move the number forward or backward in the list n times
	var i int64
	for i = 0; i < n; i++ {
		if x.value > 0 {
			seq = moveRight1(seq, x)
		} else if x.value < 0 {
			seq = moveLeft1(seq, x)
		}
	}
	return seq
}

// Move given element one position to the right
func moveRight1(seq []Number, x Number) []Number {

	// Move element to beginning of list if at end
	i := findKey(x, seq) // position of the element we want to move
	if i == len(seq)-1 {
		seq = shiftRight(seq)
		i = 0
	}

	// Swap the element with the one to its right
	tmp := seq[i+1]
	seq[i+1] = seq[i] // = x
	seq[i] = tmp
	return seq
}

// Move given element one position to the left
func moveLeft1(seq []Number, x Number) []Number {

	// Move element to end if at beginning of list
	i := findKey(x, seq) // position of the element we want to move
	if i == 0 {
		seq = shiftLeft(seq)
		i = len(seq) - 1
	}

	// Swap the element with the one to its left
	tmp := seq[i-1]
	seq[i-1] = seq[i] // = x
	seq[i] = tmp
	return seq
}

// Shift a list one position left, move first element to end
func shiftLeft(l []Number) []Number {
	return append(l[1:], l[0])
}

// Shift a list one position right, last element goes to beginning
func shiftRight(l []Number) []Number {
	return append([]Number{l[len(l)-1]}, l[:len(l)-1]...)
}

// Index of given element in list, search by position (key), not value
func findKey(n Number, l []Number) int {
	for i := 0; i < len(l); i++ {
		if l[i].pos == n.pos {
			return i
		}
	}
	fmt.Println("Warning: key", n, "not found!")
	return -1
}

// Find by value
func findValue(n Number, l []Number) int {
	for i := 0; i < len(l); i++ {
		if l[i].value == n.value {
			return i
		}
	}
	fmt.Println("Warning: value", n, "not found!")
	return -1
}
