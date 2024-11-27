// Advent of Code 2020, Day 13
//
// Solve problems related to a schdule of bus times, where all
// buses leave at t = 0, but take different number of minutes to
// reach the station. For Part 1, find the earliest bus that will
// arrive. For Part 2, find the earliest time at which the first
// bus arrives at t, the second at t+1, and so on. Used a brute
// force solution (takes 1 hr 10 mins), but there must be a better
// way.
//
// AK, 24/01/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read file and split into lines
	t, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(t), "\n")

	// Line 1 has departure time, used only for Part 1 (convert
	// to minutes since midnight)
	dep, _ := strconv.ParseInt(lines[0], 10, 64)
	dep0 := int64(dep/60)*60 + dep%60
	fmt.Printf("Departure %d = %d mins since midnight\n", dep, dep0)

	// Line 2 has list of buses (numbers of minutes, or 'x' if no bus,
	// replace these with -1)
	buses_ := strings.Split(lines[1], ",")
	buses := []int64{}
	for _, b := range buses_ {
		if b == "x" {
			buses = append(buses, -1)
		} else {
			busNo, _ := strconv.ParseInt(b, 10, 64)
			buses = append(buses, busNo)
		}
	}
	fmt.Println("Buses:", buses)

	// Do Part 1
	part1(buses, dep0)

	// Do Part 2
	part2(buses)
}

// Part 1: find the earliest bus that departs after designated time
func part1(buses []int64, dep0 int64) {

	// For each bus, find the first departure at or after desired departure
	var earliestBus, earliestTime int64
	for _, b := range buses {
		if b == -1 {
			continue
		}
		var t int64 = 0
		for t < dep0 {
			t += b
		}
		fmt.Printf("Bus %d: %d\n", b, t)
		if earliestTime == 0 || t < earliestTime {
			earliestTime = t
			earliestBus = b
		}
	}

	fmt.Printf("Part 1: Earliest bus %d leaves at %d\n", earliestBus, earliestTime)
	wait := earliestTime - dep0
	fmt.Printf("Wait %d * bus# %d = %d\n", wait, earliestBus, wait*earliestBus)
}

// Part 2: figure out the earliest time a schedule "fits". This is a brute
// force solution, there must be a faster way. The only optimization here
// is to step forward not by 1 each step, but by some interval of bus
// frequency. Started using the first bus frequency, but now search for
// the longest frequency in the input, for an even bigger gain (977
// instead of 29 min jumps). So this takes about an hour 10 mins to
// find the solution on the problem input.
//
// The problem can be restated as finding a time t such that the first bus
// arrives at that time, and each subsequent bus arrives one minute later
// (skipping a minute for each 'x' bus). Thus, there is a (different)
// multiple for each bus frequency.
//
// One of the problem examples has bus times 67, 7, 59, 61.
// The solution is 754018:
//   t     bus   t/bus
// 754018   67   11254
// 754019    7  101717
// 754020   59   12780
// 754021   61   12361
//
// If the buses were all supposed to arrive at the same time, it would
// be easy to factor out common divisors and multiply the times together,
// but the offet+1 for each row makes this trick. I was not able to find
// a solution to this, and used brute force instead.

func part2(nn []int64) int64 {

	// Find the largest bus duration in the input data, to use as fastest
	// interval for stepping through
	var step, stepi, c int64
	for i := 0; i < len(nn); i++ {
		if nn[i] > 0 && (step == 0 || nn[i] > step) {
			step = nn[i]
			stepi = int64(i)
		}
	}
	fmt.Printf("\nPart2: step %d found at %d\n", step, stepi)

	// Initial value for the time, adjusted to account for the offset
	// of the longest bus duration found above
	var x int64 = step - stepi

	// Loop until we find a solution, i.e., the first time at which
	// the first bus arrives at t, the second at t+1, etc.
	for {

		// Show progress
		if c == 1000000000 {
			fmt.Printf("\r%.4E ", float64(x))
			c = 0
		}
		c++

		// Test current value: do all the buses arrive at exactly
		// at this time plus the offset, i.e., is (t + i) / bus freq
		// an exact integer for all the buses?
		foundIt := true
		for i := 0; i < len(nn); i++ {
			if nn[i] > 0 && (x+int64(i))%nn[i] != 0 {
				foundIt = false
				break
			}
		}

		// Announce and return solution if found
		if foundIt {
			fmt.Printf("\nPart 2: found solution %d\n", x)
			return x
		}

		// Next value
		x += step
	}
}
