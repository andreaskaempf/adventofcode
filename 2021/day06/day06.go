// Day 6 of Advent of Code 2021
//
// Simulate evolution of population of "lanterfish", starting at 5 but
// reproducing every 7 days. Doing this with simulation of a list as described
// in the problem is way too slow, so converted it to using a dictionary to
// keep track of the number of fish by age, much faster.
//
// AK, 6 Dec 2021

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Sum up a list of integers
func sum(nn []int) int {
	res := 0
	for _, n := range nn {
		res += n
	}
	return res
}

func main() {

	// Read one-line CSV file, which consists of a list of numbers
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Cannot open file")
		return
	}
	raw_data, _ := io.ReadAll(f)
	data := strings.TrimSpace(string(raw_data))
	fmt.Println("Raw input data:", data)

	// Convert bytes to list of comma-separated numbers
	nums := strings.Split(string(data), ",")
	fmt.Println("Input numbers:", nums)

	// Count up the number of "fish" with each counter value
	fish := make([]int, 10, 10)
	for _, x := range nums {
		n, _ := strconv.Atoi(x)
		fish[n] += 1
	}
	fmt.Println("Counts by age:", fish, " => ", sum(fish))

	// Start simulating up to 80 days (256 days in part 2). Each day, fish
	// of age 0 becomes 6 and each adds a new fish of age 8 to the population,
	// while each other fish decreases its age by 1.
	// For test data, should be 26 fish after day 18, 5934 after 80 days, and
	// 26984457539 after 256 days
	t := 1
	for t <= 256 {

		// Each fish with a counter of 0 will create a new fish with counter 8 (but
		// don't do it yet, as we don't want to decrement the age of the new fish)
		newFish8 := fish[0]

		// All fish with counter 0 will get reset to 6 (but don't do it yet, so
		// we don't decrement the age of those fish in the next step)
		newFish6 := fish[0]
		fish[0] = 0

		// All fish of ages 1-8 get decremented, by moving them down the chain:
		//   fish0 += fish1
		//   fish1 = 0
		// then
		//   fish1 += fish2
		//   fish2 = 0
		// etc.
		for age := 1; age < 9; age++ {
			fish[age-1] += fish[age]
			fish[age] = 0
		}

		// Finally, add the new fish of ages 6 and 8
		fish[6] += newFish6
		fish[8] += newFish8

		// Show current count
		if t <= 18 || t == 80 || t == 256 {
			fmt.Printf("After %d days: %d fish\n", t, sum(fish))
		}

		// Next time period
		t++
	}
}
