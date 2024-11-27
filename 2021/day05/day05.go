// Day 5 of Advent of Code 2021
//
// Trace lines on a grid, given starting and ending coordinates of
// each line, and count up where 2 or more lines cross (part 1 only
// includes horizontal and vertical lines, part 2 diagonal as well).
//
// AK, 5 December 2021

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Widith & height of the ocean floor
const SIZE = 1000

// Set this to false for Part 2 (looks at diagonal lines)
const Part1 = false

func main() {

	// Represent the ocean floor as a matrix of 1000x1000 integers,
	// automatically initialized to zero
	floor := [][]int{}
	for i := 0; i < SIZE; i++ {
		floor = append(floor, make([]int, SIZE, SIZE))
	}

	// Read each row of the CSV file, which consists of x1,y1,x2,y2
	// (arrows changed to commas using an editor), and set matrix
	// cells to represent number of lines that cross
	//f, err := os.Open("test_input.txt")
	f, err := os.Open("input_commas.txt")
	if err != nil {
		fmt.Println("Cannot open file")
		return
	}
	r := csv.NewReader(f)
	for {

		// Read and parse one data row
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		x1, _ := strconv.Atoi(row[0])
		y1, _ := strconv.Atoi(row[1])
		x2, _ := strconv.Atoi(row[2])
		y2, _ := strconv.Atoi(row[3])
		//fmt.Println(row, "=>", x1, y1, x2, y2)

		// For Part 1, ignore diagonals
		if Part1 && !(x1 == x2 || y1 == y2) {
			//fmt.Println("  Skipping diagonal")
			continue
		}

		// Determine the direction of travel for x and y
		var dx, dy int // initialized to zero, for horiz/vertical
		if x1 < x2 {
			dx = 1
		} else if x1 > x2 {
			dx = -1
		}
		if y1 < y2 {
			dy = 1
		} else if y1 > y2 {
			dy = -1
		}
		//fmt.Printf("  dx = %d, dy = %d\n", dx, dy)

		// Increment counters for all points covered by this line
		x := x1
		y := y1
		for {
			floor[y][x] += 1        // first index is row, second is column
			if x == x2 && y == y2 { // finish loop
				break
			}
			x += dx
			y += dy
		}
	}

	// Now count up the number of cells with 2 or more crossings
	twos := 0
	for _, r := range floor { // each row
		for _, c := range r { // each column within row
			if c >= 2 {
				twos++
			}
		}
	}
	fmt.Println("Count of 2+ overlaps =", twos)
}
