// Advent of Code 2020, Day 15
//
// Simulate a convoluted memory game, which becomes infeasible using a simple
// list of history, when the number of iterations goes from 2020 (part 1) to
// 30 million (part 2).
//
// AK, 15/10/2022

package main

import "fmt"

func main() {

	// Test inputs, check 2020th iteration
	// Results should be 436, 1, 10, 27, 78, 438, 1836
	fmt.Println("Part 1 tests:")
	samples := [][]int{[]int{0, 3, 6}, []int{1, 3, 2}, []int{2, 1, 3},
		[]int{1, 2, 3}, []int{2, 3, 1}, []int{3, 2, 1}, []int{3, 1, 2}}
	for _, s := range samples {
		fmt.Println(s, part2(s, 2020))
	}

	// Main puzzle input, Part 1 (should be 639)
	input := []int{11, 18, 0, 20, 1, 7, 16}
	fmt.Println("Part 1 (2020 iterations):", part1(input, 2020))

	// Part 2: same, but 30M iterations (unfeasible if you use simple
	// list to keep track of history, had to change to dictionary)
	// Should be: 175594, 2578, 3544142, 261214, 6895259, 18, 362
	fmt.Println("\nPart 2 tests:")
	iters := 30000000
	for _, s := range samples {
		fmt.Println(s, part2(s, iters))
	}
	fmt.Println("Part 2 (30M iterations):", part2(input, iters))
}

// Part 1: simple memory game, too slow for Part 2
func part1(input []int, iters int) int {

	// History of numbers already recited, in one long list
	history := []int{}

	// Execute each turn
	for turn := 1; turn <= iters; turn++ {

		if turn%10000 == 0 {
			fmt.Println("Iteration", turn)
		}

		// First few turns read from list of numbers
		if turn <= len(input) {
			n := input[turn-1]
			history = append(history, n)
			//fmt.Printf("Turn %d: %d\n", turn, n)
			continue
		}

		// Find out when the last number spoken was previously spoken
		last := history[len(history)-1] // The last  number spoken
		lastSpoken := -1                // turn number when it was last spoken
		for i := len(history) - 1; i > 0; i-- {
			if history[i-1] == last {
				lastSpoken = i
				break
			}
		}

		// If that was the first time that number was spoken, say 0;
		// Otherwise, announce the number of turns since previously spoken
		if lastSpoken == -1 {
			n := 0
			//fmt.Printf("Turn %d: %d\n", turn, n)
			history = append(history, n)
		} else {
			n := turn - lastSpoken - 1
			//fmt.Printf("Turn %d: %d\n", turn, n)
			history = append(history, n)
		}
	}

	// Return the last value
	//fmt.Println("Last value:", history[len(history)-1])
	return history[len(history)-1]
}

// Part 2: same, but capable of more iterations by using dictionary instead of list
func part2(input []int, iters int) int {

	// History of numbers already recited: for each number, a list of the
	// turns in which it was recited
	history2 := map[int][]int{}
	var last int // the last number spoken

	// Execute each turn
	for turn := 1; turn <= iters; turn++ {

		// First few turns read from list of numbers
		if turn <= len(input) {
			last = input[turn-1]
			history2[last] = append(history2[last], turn)
			continue
		}

		// Find out when the last number spoken was previously spoken
		hist, ok := history2[last]
		lastSpoken := -1
		if ok && len(hist) > 1 {
			lastSpoken = hist[len(hist)-2] // not the last turn but the one before that
		}

		// If that was the first time that number was spoken, say 0;
		// Otherwise, announce the number of turns since previously spoken
		if lastSpoken == -1 {
			last = 0
		} else {
			last = turn - lastSpoken - 1
		}

		// Add this turn to the history for this number
		history2[last] = append(history2[last], turn)
	}

	// Return the last value
	return last
}
