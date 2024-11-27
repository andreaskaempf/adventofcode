// Utility functions

package main

import (
	"strconv"
)

// Convert a big integer to a list of digits
func intToDigits(n int64) []int {
	result := []int{}
	s := strconv.FormatInt(n, 10)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			panic("Invalid digit")
		}
		result = append(result, int(c-'0'))
	}
	return result
}

// Represent a list of digits as a number
func digitsToString(digits []int) string {
	chars := []byte{}
	for _, d := range digits {
		chars = append(chars, byte(d+'0'))
	}
	return string(chars)
}

// Check if two integer lists are the same
func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Check if a string has any zeros in it
func hasZeros_(s string) bool {
	for _, c := range s {
		if c == '0' {
			return true
		}
	}
	return false
}

// Increment a number, represented as an array of digits;
// skip 0s, as required by problem
func increment(nn []int) []int {
	nn[len(nn)-1] += 1
	for i := len(nn) - 1; i >= 0 && nn[i] > 9; i-- {
		nn[i] = 1 // skip zero
		nn[i-1] += 1
	}
	return nn
}

// Decrement a number, represented as an array of digits;
// skip 0s, as required by problem
func decrement(nn []int) []int {
	nn[len(nn)-1] -= 1
	if nn[len(nn)-1] == 0 { // skip zero
		nn[len(nn)-1] -= 1
	}
	for i := len(nn) - 1; i >= 0 && nn[i] < 0; i-- {
		nn[i] = 9
		nn[i-1] -= 1
		if nn[i-1] == 0 {
			nn[i-1] -= 1
		}
	}
	return nn
}
