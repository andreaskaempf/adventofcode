// Advent of Code 2022, Day 03
//
// Sum up characters that are common in two halves of some strings (Part 1),
// and common to groups of 3 strings (Part 2)
//
// AK, 3 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Read the input file
	//lines := readLines("sample.txt")
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Part 1: sum up common characters in the two halves of each line
	tot := 0
	for _, l := range lines {
		a := l[:len(l)/2]      // left half of string
		b := l[len(l)/2:]      // right half of string
		isects := common(a, b) // common character(s)
		tot += cval(isects[0])
	}
	fmt.Println("Part 1 (s/b 157):", tot)

	// Part 2: add up characters that are common to entire line
	// in each group of 3 lines
	tot = 0
	for i := 0; i < len(lines); i += 3 {
		l1 := lines[i]
		l2 := lines[i+1]
		l3 := lines[i+2]
		c1 := common(l1, l2) // chars common to l1 and l2
		for _, c := range c1 {
			if in(c, l3) { // char that is also in l3
				tot += cval(c)
				break
			}
		}
	}
	fmt.Println("Part 2 (s/b 70):", tot)
}

// Get list of characters that are common to two strings
func common(a, b string) []byte {
	res := []byte{}
	for i := 0; i < len(a); i++ {
		if in(a[i], b) {
			res = append(res, a[i])
		}
	}
	return res
}

// Is character in a string?
func in(c byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Numeric value of a character
func cval(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	} else {
		return 0
	}
}

// Read lines from the input file
func readLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
