// Advent of Code 2024, Day 11
//
// Apply transformations to a list of numbers, so that zeros become ones,
// numbers with an even number of digits get split in half, or the number
// gets multiplied by 2024 (only one test per number). Do all these
// simulataneously, i.e., changing a copy of the data and referencing the
// original without changing it. For Part 1, it was okay to naively build up
// a list over 25 iterations, but for Part 2 there were 75 iterations, and
// memory ran out after about 45 iterations. So changed formulation to use
// a dictionary, since the numbers could be processed in any order, and many
// of the numbers are repeated.
//
// AK, 11 Dec 2024

package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Input data: first line is sample data, uncomment the second line to
	// use input data
	data := []int{125, 17}
	data = []int{5910927, 0, 1, 47, 261223, 94788, 545, 7771}

	// Do 25 iterations for Part 1, 75 for Part 2
	iters := 25

	// Convert list of stones to a frequency dict
	stones := map[int]int{}
	for _, n := range data {
		stones[n] += 1
	}

	// Perform iterations ("blinks")
	for i := 0; i < iters; i++ { // runs out of memory after 45 iterations

		// Get list of keys, and make a deep copy of the dictionary, since
		// we can't update the dictionary while directly iterating over its keys
		nn := []int{}          // list of keys, i.e., numbers on stones
		dict2 := map[int]int{} // deep copy of stones with counts
		for k, v := range stones {
			nn = append(nn, k)
			dict2[k] = v
		}

		// Transform each stone, updating counters
		for _, n := range nn {

			nstones := stones[n]      // count for this type of stone
			s := fmt.Sprintf("%d", n) // Format to string, for second test

			// Decrement this type of stone, because they will be converted
			dict2[n] -= nstones

			// Apply first transformation that matches
			if n == 0 { // 0 becomes 1
				dict2[1] += nstones // = dict[1] + 1
			} else if len(s)%2 == 0 { // even digits -> split in half
				l := atoi(s[:len(s)/2])
				r := atoi(s[len(s)/2:])
				dict2[l] += nstones
				dict2[r] += nstones
			} else { // otherwise multiply by 2024
				dict2[n*2024] += nstones
			}
		}

		// Replace original dictionary with modified one
		stones = dict2
	}

	// For answer, add up the number of stones, just sum of dictionary values
	// Sample: 22 after 6 iterations, 55312 after 25
	// Part 1: 193607 after 25 iterations
	// Part 2: 229557103025807 after 75 iterations
	nstones := 0
	for _, v := range stones {
		nstones += v
	}
	fmt.Println(nstones)
}

// Parse an integer, return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}
