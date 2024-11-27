// Advent of Code 2021, Day 24
//
// Check whether a 14-digit code is valid based on assembly-language like
// series of computations, and find the highest (Part 1) or lowest (Part 2)
// codes that are valid, i.e, leave a "z" register at zero after the
// computations.
//
// Not a good solution, used brute force with Go parallelization and some
// pruning of solutions, came up with the right answers but not efficient
// or clever. Took 24 mins for Part 1, 33 mins for Part 2.
//
// AK, 29 December 2021

package main

import (
	"fmt"
	"os"
)

// The magic numbers
var A = [14]int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}          // 7 pairs!
var B = [14]int{12, 11, 13, 11, 14, -10, 11, -9, -3, 13, -5, -10, -4, -5} // 7 negs!
var C = [14]int{4, 11, 5, 11, 14, 7, 11, 4, 6, 5, 9, 12, 14, 14}

// Set this for the part you want to run
var Part1 bool = true

// Simplified implementation of the original algorithm. Works with any length
// of input, including less than 14, for validation so can cut short
// unfeasible search space
func monad3(input []int, verbose bool) int {

	// The registers (all start at zero)
	var w, z, prevZ, increases int

	// Simplified transcription of the problem, entirely in terms of w and z
	// (x,y no longer needed). Quick fail if any digit causes Branch 2 to
	// be chosen when A == 26, or vice versa, or if z keeps increasing.
	for i := 0; i < len(input); i++ {

		// Get the next digit, validate it
		w = input[i]
		if w < 1 || w > 9 {
			panic("Invalid digit")
		}

		if w == (z%26)+B[i] { // x = 0, y = 0
			if A[i] == 1 {
				return -1 // can't be valid, fail here
			}
			z = div(z, A[i]) // Branch 1
		} else { // x = 1, y = w + C[i]
			if A[i] == 26 {
				return -1 // can't be valid, fail here
			}
			z = div(z, A[i])*26 + w + C[i] // Branch 2
		}

		// z needs to start decreasing at some point, stop if this is not
		// happening by the midway point
		if z > prevZ {
			increases++
		}
		prevZ = z
		if increases > 7 {
			return -1
		}

		// Verbose debugging info
		if verbose {
			fmt.Printf("%2d: w = %d, z = %8d, A = %3d, B = %3d, C = %3d\n",
				i, w, z, A[i], B[i], C[i])
		}
	}

	// Input is considered valid if z is zero at the end
	return z
}

// One worker for concurrent processing, processes any available
// items groups and sends results to output channel
func worker(jobs <-chan int64, results chan<- int) {
	for j := range jobs {
		results <- processBatch(j)
	}
}

// Process one range of numbers, but check for feasibility first,
// so we don't waste time on batches that won't work
func processBatch(prefix int64) int {

	fmt.Printf("Processing %d\n", prefix)

	// Range is prefix to a 14-digit number, so convert it to
	// a list of digits
	digits := intToDigits(prefix)

	// Do a preliminary check on whether numbers starting with this prefix
	// are likely to be valid
	for _, d := range digits {
		if d == 0 {
			fmt.Println("  skipping prefix with zero:", digitsToString(digits))
			return -1
		}
	}
	if monad3(digits, false) == -1 {
		fmt.Println("  prefix failed, skipping batch", digitsToString(digits))
		return 0
	}

	// Pad the prefix to full 14 digits, and create ending number as well
	until := make([]int, len(digits))
	copy(until, digits)
	if !same(until, digits) {
		fmt.Println("Copy failed:", digits, until)
		panic("Copy failed!")
	}
	for len(digits) < 14 {
		digits = append(digits, 1)
		until = append(until, 9)
	}

	// Start searching for solutions within this range
	solutions := 0
	for true {

		// Calculate monad and if a solution is found, write to
		// a file so we don't lose the info, then exit immediately
		z := monad3(digits, false)
		if z == 0 {
			fmt.Println("*** Solution found:", digitsToString(digits))
			os.WriteFile(digitsToString(digits), []byte("z == 0\n"), 0644)
			solutions += 1
			os.Exit(1)
		}

		// If this is the end of the batch, quit loop
		if same(digits, until) {
			break
		}

		// Next number
		if Part1 {
			digits = increment(digits)
		} else {
			digits = decrement(digits)
		}
	}

	fmt.Printf("  finished prefix %d: %d solutions found\n", prefix, solutions)
	return solutions
}

// Main execution splits all numbers into batches and run them concurrently
func main() {

	// Create a list of numbers, each the beginning of a batch
	// The prefixes can be any length, so shorter for bigger
	// batches
	var n, until, step int64 // For Part 2, just the leading digits
	if Part1 {
		n = 999999     // just the leading digits
		until = 111111 // same number of digits
		step = -1
	} else {
		n = 111111     // just the leading digits
		until = 999999 // same number of digits
		step = 1
	}
	batches := []int64{}
	for n != until {
		batches = append(batches, n)
		n += step
	}

	// Input and output queues
	// Concurrent worker pattern borrowed from: https://gobyexample.com/worker-pools
	numJobs := len(batches)
	jobs := make(chan int64, numJobs)
	results := make(chan int, numJobs)

	// Start workers, one per CPU
	for i := 0; i < 8; i++ {
		go worker(jobs, results)
	}

	// Add each batch to the jobs queue
	for _, b := range batches {
		jobs <- b
	}
	close(jobs)

	// Aggregate the results from the ouput queue
	total := 0
	for i := 0; i < numJobs; i++ {
		res := <-results
		total += res
	}
}
