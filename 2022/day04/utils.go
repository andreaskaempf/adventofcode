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

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}

// Parse a float, show message and return -1 if error
func atof(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Could not parse float:", s)
		n = -1
	}
	return float64(n)
}

// Maximum of a list (of ints, floats, or strings, using generics)
func max[T int | float64 | string](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] > y {
			y = l[i]
		}
	}
	return y
}

// Minimum of a list (of ints, floats, or strings, using generics)
func min[T int | float64 | string](l []T) T {
	var y T
	for i := 0; i < len(l); i++ {
		if i == 0 || l[i] < y {
			y = l[i]
		}
	}
	return y
}
