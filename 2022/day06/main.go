// Advent of Code 2022, Day 06
//
// Look for first block of 4 (Part 1) or 14 (Part 2) non-repeating
// characters in a string. Increased performance 10x by using concurrency
// and simpler check for duplicate characters.
//
// AK, 6 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Read input from file
	//data, _ := ioutil.ReadFile("input.txt")
	data, _ := ioutil.ReadFile("day06_input10000")
	msg := string(data)

	// Concurrently look for markers of length 4 (Part 1) and 14 (Part 2),
	// using a "channel" to allow concurrent processes to return results
	// when ready (code will continue when both results are sent back)
	c := make(chan int)       // create the channel
	go marker(&msg, 4, c)     // run the 4-char check in background
	go marker(&msg, 14, c)    // same for 14-char check
	part1, part2 := <-c, <-c  // collect results when ready
	fmt.Println(part1, part2) // show results (may be out of order!)
}

// Find the position of the "marker", i.e., a block where n chars
// are all different, send result to channel
func marker(s *string, n int, c chan int) {
	for i := n - 1; i < len(*s); i++ {
		//substr := s[i-n+1 : i+1]  // get n-char substring
		//if !duplicates2(substr) { // no duplicate chars?
		if !duplicates3(s, i-n+1, i) { // no duplicate chars?
			c <- i + 1 // send back position of marker
			return
		}
	}
	c <- -1 // no marker found
}

// Check if a string contains any duplicate characters
// (first implementation, slow because using map)
func duplicates(s string) bool {
	chars := map[byte]int{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if chars[c] > 0 { // duplicate found
			return true
		}
		chars[c] += 1 // remember we found this character
	}
	return false
}

// Check if a string contains any duplicate characters,
// second and faster implementation using character comparisons inside
// loops instead of map
func duplicates2(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}

// Check if a string contains any duplicate characters,
// third and even faster version that takes a pointer to a string,
// rather than a new substring copy each time; i and j are the indexes
// of the characters to be searched, zero-based and inclusive
func duplicates3(s *string, i, j int) bool {
	for a := i; a <= j; a++ {
		for b := i; b < a; b++ {
			if (*s)[a] == (*s)[b] {
				return true
			}
		}
	}
	return false
}
