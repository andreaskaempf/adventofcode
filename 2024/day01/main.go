// Advent of Code 2024, Day 01
//
// Given rows of pairs of numbers, extract the two lists, sort them, and
// find the sum of the absolute difference (Part 1). For Part 2, sum up
// the product of each number in the left list, times the number of times
// it occurs in the right list.
//
// AK, 01 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Read the input file into two lists of numbers
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")
	list1 := []int{}
	list2 := []int{}
	for _, r := range rows {
		parts := strings.Split(r, "   ")
		if len(parts) == 2 {
			list1 = append(list1, parseInt(parts[0]))
			list2 = append(list2, parseInt(parts[1]))
		}
	}

	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)

	// For Part 1, add up the absolute differences
	part1 := 0
	for i := 0; i < len(list1); i++ {
		part1 += abs(list1[i] - list2[i])
	}
	fmt.Println("Part 1 (s/b 1938424):", part1)

	// For Part 2, calculate a "similarity score" by adding up the product
	// of each number in the left list, times the number of times it appears
	// in the right list
	part2 := 0
	for _, i := range list1 {
		part2 += i * occurs(i, list2)
	}
	fmt.Println("Part 2 (s/b 22014209):", part2)
}

// Count the number of times a number occurs in a list
func occurs(n int, l []int) int {
	c := 0
	for _, i := range l {
		if i == n {
			c++
		}
	}
	return c
}

// Parse an integer
func parseInt(s string) int {
	n, _ := strconv.Atoi(s) // TODO: handle error
	return n
}

// Absolute value of an integer
func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}
