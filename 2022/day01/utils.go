// Utility functions for Advent of Code

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Read lines from the input file
func readLines(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

// Parse number
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse, assuming zero:", s)
		n = 0
	}
	return int(n)
}

// Maximum of a list
func max(l []int) int {
	var y int
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] > y {
			y = l[i]
		}
	}
	return y
}
