// Advent of Code 2022, Day 01
//
// Find the maximum sum of newline-separated groups of numbers.
// For part 2, find the sum of the top 3.
//
// AK, 1 Dec 2022

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Read the input file
	//lines := readLines("sample.txt")
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Add up groups of integers, separated by blank lines
	nums := []int{} // list of totals
	var tot int
	for _, x := range lines {
		if len(x) == 0 {
			nums = append(nums, tot)
			tot = 0
		} else {
			tot += atoi(x)
		}
	}
	nums = append(nums, tot)

	// Part 1: report the largest total
	fmt.Println("Part 1:", max(nums))

	// Part 2: sum of top 3
	sort.Ints(nums)
	tot = 0
	for i := 0; i < 3; i++ {
		tot += nums[len(nums)-1-i]
	}
	fmt.Println("Part 2:", tot)
}
