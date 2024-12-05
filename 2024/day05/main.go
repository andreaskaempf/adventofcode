// Advent of Code 2024, Day 05
//
// You are given a list of page number pairs, indicating that the right
// page must be printed after the left page number. You are also given a list
// of manual "updates", each a list of page numbers. For each update, determine
// if the pages are in the correct sequence, and add up the middle page numbers.
// For Part 2, rearrange the  incorrectly ordered updates, and add up the middle
// page numbers.
//
// AK, 05 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// A pair of page numbers, right number must be printed AFTER left number
type Seq struct {
	l, r int
}

// List of sequence pairs (global)
var seqs []Seq

func main() {

	// Read the input file
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

	// Set up list of sequence pairs, and list of "updates"
	seqs = []Seq{}
	updates := [][]int{}
	for _, l := range lines {
		if len(l) == 0 {
			continue
		} else if l[2] == '|' {
			parts := strings.Split(l, "|")
			seq := Seq{atoi(parts[0]), atoi(parts[1])}
			seqs = append(seqs, seq)
		} else {
			updates = append(updates, parseList(l))
		}
	}

	// Part 1: process each update, checking if each page number if followed by
	// page numbers that are greater than it, add up middle page numbers of
	// correctly ordered updates.
	var part1, part2 int
	incorrect := [][]int{} // for part 2
	for _, u := range updates {

		// Check if correctly ordered, by taking each page number, and making
		// sure all the subsequent page numbers have higher sequence. This is
		// admittedly rather inefficient, but let's Go with it.
		ok := true
		for i, page := range u {
			if !allGreater(page, u[i:]) {
				ok = false
				break
			}
		}

		// If okay, add middle page number to answer, otherwise remember
		// update for part 2
		if ok {
			mid := u[len(u)/2]
			part1 += mid
		} else {
			incorrect = append(incorrect, u)
		}
	}
	fmt.Println("Part 1 (s/b 4185) =", part1)

	// Part 2: put the incorrectly sequenced updates in order, and add up
	// middle pages. This is very easy because we can re-use the isGreater
	// comparison function we used in Part 1 as an argument to sort.Slice
	for _, u := range incorrect {
		sort.Slice(u, func(i, j int) bool {
			return isGreater(u[i], u[j])
		})
		mid := u[len(u)/2]
		part2 += mid
	}
	fmt.Println("Part 2 (s/b 4480) =", part2)
}

// Check if all the page numbers in a list are greater than the given
// page number, using the sequence rules
func allGreater(p int, pages []int) bool {
	for _, p2 := range pages {
		if !isGreater(p, p2) {
			return false
		}
	}
	return true
}

// Check if page number p2 is greater than p, using the sequence rules;
// very inefficient, could be done with a map for more speed, but okay for now
func isGreater(p, p2 int) bool {
	for _, seq := range seqs {
		if seq.l == p2 && seq.r == p {
			return false
		}
	}
	return true
}

// Parse a list of integers
func parseList(s string) []int {
	nums := []int{}
	for _, n := range strings.Split(s, ",") {
		nums = append(nums, atoi(n))
	}
	return nums
}

// Parse an integer
func atoi(s string) int {
	n, _ := strconv.Atoi(s) // should check error
	return n
}
