// Advent of Code 2019, Day 02
//
// Simulate an assembly language-like CPU, with a list of numbers that
// represent opcodes (just add or multiply) followed by arguments that are
// input and output locations in memory. For Part 1, change two values in
// memory to fixed values, run the program, and report final value in location
// zero. For Part 2, we find the combination that results in a certain output
// value (just used brute force).
//
// AK, 29 Nov 2024

package main

import (
	"fmt"
)

func main() {

	// Input is just a list of numbers, hard coded rather than input files
	data := []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19,
		1, 9, 19, 23, 2, 23, 13, 27, 1, 27, 9, 31, 2, 31, 6, 35, 1, 5, 35, 39,
		1, 10, 39, 43, 2, 43, 6, 47, 1, 10, 47, 51, 2, 6, 51, 55, 1, 5, 55, 59,
		1, 59, 9, 63, 1, 13, 63, 67, 2, 6, 67, 71, 1, 5, 71, 75, 2, 6, 75, 79,
		2, 79, 6, 83, 1, 13, 83, 87, 1, 9, 87, 91, 1, 9, 91, 95, 1, 5, 95, 99,
		1, 5, 99, 103, 2, 13, 103, 107, 1, 6, 107, 111, 1, 9, 111, 115, 2, 6,
		115, 119, 1, 13, 119, 123, 1, 123, 6, 127, 1, 127, 5, 131, 2, 10, 131,
		135, 2, 135, 10, 139, 1, 13, 139, 143, 1, 10, 143, 147, 1, 2, 147, 151,
		1, 6, 151, 0, 99, 2, 14, 0, 0} // the input values
	//data = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50} // uncomment for sample values

	// Part 1: get simulation result with two initial input values
	fmt.Println("Part 1 (s/b 3306701) =", simulate(data, 12, 2))

	// Part 2: find the combination of input values that yields a target value
	done := false
	for init1 := 0; init1 < 100 && !done; init1++ {
		for init2 := 0; init2 < 100 && !done; init2++ {
			res := simulate(data, init1, init2)
			//fmt.Printf("%d, %d => %d\n", init1, init2, res)
			if res == 19690720 { // looking for this output value
				fmt.Println("Part 2 (s/b/ 7621) =", init1*100+init2)
				done = true
			}
		}
	}
}

// Simulate the execution of a program, return final value in position 0
func simulate(program []int, init1, init2 int) int {

	// Make a memory map of values at each location
	mem := map[int]int{}
	for i, n := range program {
		mem[i] = n
	}

	// Initialize state with input values at locations 1 & 2. For Part 1,
	// these are fixed. For Part 2, we need to find the combination that
	// results in a certain output value
	mem[1] = init1
	mem[2] = init2

	// Process instructions, each 4 numbers:
	// - opcode (1 = add, 2 = multiply, 99 = halt)
	// - input 1
	// - input 2
	// - output location
	ip := 0             // instruction pointer
	op := mem[ip]       // next operation
	inLoc1 := mem[ip+1] // will be zero if nothing at that location
	inLoc2 := mem[ip+2]
	outLoc := mem[ip+3]
	for op != 99 {

		// Execute instruction
		if op == 1 {
			mem[outLoc] = mem[inLoc1] + mem[inLoc2]
		} else if op == 2 {
			mem[outLoc] = mem[inLoc1] * mem[inLoc2]
		}

		// Get next instruction
		ip += 4            // instruction pointer to next instruction
		op = mem[ip]       // next operation
		inLoc1 = mem[ip+1] // will be zero if nothing at that location
		inLoc2 = mem[ip+2]
		outLoc = mem[ip+3]
	}

	// Answer is the value at location 0
	return mem[0]
}
