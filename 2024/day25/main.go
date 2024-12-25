// Advent of Code 2024, Day 25
//
// Input is a series of 6x5 blocks of # and . characters, representing
// locks (heights from top down) or keys (heights from bottom up).
// Find out how many key & lock pairs fit, i.e., heights to not
// overlap. Checking was trivial, but parsing the input into arrays
// of heights was a chore. There is no Part 2, granted automatically
// when you complete all the other puzzles.
//
// AK, 25 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// Global lists of column heights for locks and keys
var locks [][]int
var keys [][]int

func main() {

	// Read rows of data from input file
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Convert blocks to rows & key heights
	// NOTE: there must be a blank line at the bottom of input file
	item := [][]byte{}
	for _, r := range rows {
		if len(r) > 0 {
			item = append(item, r)
		} else {
			parseLockKey(item)
			item = [][]byte{}
		}
	}
	fmt.Println("Locks:", locks)
	fmt.Println("Keys:", keys)

	// Part 1: how many unique lock/key pairs fit together
	// without overlapping in any column? Just look at
	// each pair and make sure each column does not add
	// up to more than 5.
	ans := 0
	for _, lock := range locks {
		for _, key := range keys {
			fits := true
			for c := 0; c < 5; c++ {
				if lock[c]+key[c] > 5 {
					fits = false
				}
			}
			fmt.Println(lock, key, fits)
			if fits {
				ans++
			}
		}
	}
	fmt.Println("Part 1 =", ans)

	// There is no Part 2, granted automatically when you get
	// all 50 stars (have 2 more to do)
}

// Parse a 6x5 matrix of characters, either a lock (top row solid #####)
// or a key (bottom row solid #####). Convert to list of heights and add
// to global lists.
func parseLockKey(item [][]byte) {

	if string(item[0]) == "#####" { // a lock, heights from top down
		heights := make([]int, 5, 5)
		for c := 0; c < 5; c++ {
			h := 0
			for r := 1; r < 7; r++ { // measure height from second row down
				if item[r][c] == '.' {
					break
				}
				h++
			}
			heights[c] = h
		}
		locks = append(locks, heights)
	} else { // a key, heights from bottom up
		heights := make([]int, 5, 5)
		for c := 0; c < 5; c++ {
			h := 0
			for r := 5; r > 0; r-- { // measure height from bottom up
				if item[r][c] == '.' {
					break
				}
				h++
			}
			heights[c] = h
		}
		keys = append(keys, heights)
	}
}
