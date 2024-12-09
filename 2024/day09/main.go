// Advent of Code 2024, Day 09
//
// Simulate two different strategies for disk fragmentation, taking a list of
// digits where each pair of digit is a file size followed by space after the data
// (both in blocks). For Part 1,
// move parts of files to first available space, and calculate a "checksum".
// For Part 2, you move the whole file (all blocks) but start with the file with
// highest ID.
//
// AK, 09 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"slices"
)

type File struct {
	id, size, space int
}

func main() {

	// Read data, only one row
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)

	// Part 1: keep taking one unit at a time from the last block and
	// move that space to new blocks of length 1 in the first available spot
	files := parseData(data)
	for {

		// Find the first block with any free space after it
		insertAfter := -1
		for i := 0; i < len(files); i++ {
			if files[i].space > 0 {
				insertAfter = i
				break
			}
		}

		// If it's the last block, we are done
		if insertAfter == -1 || insertAfter == len(files)-1 {
			break
		}

		// Insert a new block after this location, with the same id as then end, length 1
		b := files[len(files)-1] // the last block
		files = slices.Insert(files, insertAfter+1, File{b.id, 1, files[insertAfter].space - 1})
		files[insertAfter].space = 0 // reduce padding of this block

		// Reduce size of last block, remove if empty
		files[len(files)-1].size -= 1
		files[len(files)-1].space += 1
		if files[len(files)-1].size == 0 {
			trailing := files[len(files)-1].space // space after last block
			files = files[:len(files)-1]          // remove the last block
			files[len(files)-1].space += trailing // add back the space
		}
	}

	fmt.Println("Part 1:", checksum(files)) // s/b 1928 / 6288599492129

	// For Part 2, keep taking the file with the highest ID, and try to move
	// the whole thing into the first available space that is big enough
	files = parseData(data)
	for fid := len(files); fid >= 0; fid-- { // files decreasing ID order

		// Find the file's current location
		this := 0
		for i := 0; i < len(files); i++ {
			if files[i].id == fid {
				this = i
				break
			}
		}

		// Is there space for it somewhere to the left?
		moveAfter := -1
		for i := 0; i < this; i++ {
			if files[i].space >= files[this].size {
				moveAfter = i
				break
			}
		}

		// If not, try the next ID
		if moveAfter == -1 {
			continue
		}

		// Remove the block from list
		b := files[this]                        // the block we are moving
		files[this-1].space += b.size + b.space // fill space removed
		files := append(files[:this], files[this+1:]...)

		// Reinsert it into new location
		newLoc := moveAfter + 1
		files = slices.Insert(files, newLoc, b)

		// Adjust the spacing
		files[newLoc].space = files[moveAfter].space - b.size
		files[moveAfter].space = 0
	}

	fmt.Println("Part 2:", checksum(files)) // s/b 2858 / 6321896265143
}

// Parse a string of digits into a list of Files, each with ID, size,
// and padding
func parseData(data []byte) []File {

	files := []File{}
	var id int // id starts at zero
	for i := 0; i < len(data); i += 2 {

		// Get size and space following block
		size := int(data[i] - '0')
		space := 0
		if i+1 < len(data) {
			space = int(data[i+1] - '0')
		}

		// Create file object
		file := File{id, size, space}
		files = append(files, file)

		// Next position and ID
		id += 1
	}
	return files
}

// Calculate checksum
func checksum(files []File) int {
	var res, pos int
	for i := 0; i < len(files); i++ {
		id := files[i].id
		for j := 0; j < files[i].size; j++ {
			res += id * pos
			pos++
		}
		pos += files[i].space // required for part 2 because of gaps
	}
	return res
}

// Print a map like in the problem statement, for debugging
func printmap(files []File) {
	for _, f := range files {
		for i := 0; i < f.size; i++ {
			fmt.Printf("%d", f.id)
		}
		for i := 0; i < f.space; i++ {
			fmt.Printf(".")
		}
	}
	fmt.Println("")
}
