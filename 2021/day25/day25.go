// Advent of Code 2021, Day 25
//
// Simulation of "sea cucumber" particle movement, iterating until
// the system reaches a stable state (i.e., no more changes). Part 1
// no problem, but Part 2 only granted when have completed all previous
// problems.
//
// AK, 25/12/2021

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Read each line of input file into a list of byte arrays
func readData(filename string) [][]byte {

	data := [][]byte{}
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		r := make([]byte, len(t), len(t))
		for i := 0; i < len(t); i++ {
			r[i] = t[i]
		}
		data = append(data, r)
	}
	return data
}

// Make a deep copy of the state
func copyData(d1 [][]byte) [][]byte {
	d2 := make([][]byte, len(d1), len(d1))
	for i, _ := range d1 {
		d2[i] = make([]byte, len(d1[i]))
		copy(d2[i], d1[i])
	}
	return d2
}

// Print the grid, for debugging
func printGrid(g [][]byte) {
	for _, r := range g {
		fmt.Println(string(r))
	}
}

func main() {

	// Read data
	data := readData("input.txt")
	fmt.Println("Starting data:")
	printGrid(data)

	// Iterate: move each cell in its designated direction if
	// there is room, working on a copy
	iter := 0
	for true {

		iter++
		fmt.Printf("\nStarting iteration %d\n", iter)
		printGrid(data)

		// Sideways move first
		d2 := copyData(data)
		changed := false
		for r := 0; r < len(data); r++ {
			for c := 0; c < len(data[r]); c++ {
				if data[r][c] == '>' {
					c1 := c + 1
					if c1 >= len(data[r]) {
						c1 = 0
					}
					if data[r][c1] == '.' {
						d2[r][c1] = '>'
						d2[r][c] = '.'
						changed = true
					}
				}
			}
		}

		// Downward move next,
		data = d2 // use the changed data and make another copy
		d2 = copyData(data)
		for r := 0; r < len(data); r++ {
			for c := 0; c < len(data[r]); c++ {
				if data[r][c] == 'v' {
					r1 := r + 1
					if r1 >= len(data) {
						r1 = 0
					}
					if data[r1][c] == '.' {
						d2[r1][c] = 'v'
						d2[r][c] = '.'
						changed = true
					}
				}
			}
		}

		// All done if no changes
		if !changed {
			fmt.Printf("No changes after iteration %d\n", iter)
			break
		}

		// Otherwise start next iteration with changed data
		data = d2
	}
}
