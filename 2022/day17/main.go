// Advent of Code 2022, Day 17
//
// Simulate simple geometric shapes falling down a shaft, getting moved left
// and right by gusts of "gas", and falling on top of each other. For part 1,
// determine the total height of the shapes after 2022 have fallen. For Part 2,
// do the same for 1 000 000 000 000 shapes (infeasible to simulate, need to
// find a better way, perhaps by finding when the pattern starts to repeat).
//
// AK, 17 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
)

// A position in 2-d space
type Point struct {
	x, y int64
}

// Rocks in the chamber (sparse matrix)
var chamber map[Point]byte

func main() {

	// Read the input file, just one line, but be sure to
	// remove trailing newlines with editor
	fname := "sample.txt"
	fname = "input.txt"
	patt, _ := ioutil.ReadFile(fname) // returns array of bytes
	fmt.Println("Pattern length is", len(patt))

	// Five rock shapes, expressed as pseudo-matrices
	minus := [][]int{[]int{1, 1, 1, 1}}
	plus := [][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 0}}
	L := [][]int{[]int{0, 0, 1}, []int{0, 0, 1}, []int{1, 1, 1}}
	I := [][]int{[]int{1}, []int{1}, []int{1}, []int{1}}
	square := [][]int{[]int{1, 1}, []int{1, 1}}
	shapes := [][][]int{minus, plus, L, I, square}

	// Simulate the falling of rocks to the bottom of a chamber 7-wide
	chamber = map[Point]byte{} // initialize map used as sparse matrix
	nextShape := 0             // type of the next rock
	var height int64 = 0       // current height of the hightest rock
	pi := 0                    // start in position 0 of pattern
	var rocks int64 = 20000    // 2022 for Part 1, longer for Part 2
	var rock int64             // the current rock
	var prevHeight int64       // previous height of chamber, so we can calculate deltas for part 2
	deltas := []int64{}        // height added by each rock during simulation
	for rock = 1; rock <= rocks; rock++ {

		// Get the shape of this rock
		shape := shapes[nextShape]
		nextShape++
		if nextShape >= len(shapes) {
			nextShape = 0
		}

		// Simulate appearance of the new rock: each rock appears so that its
		// left edge is two units away from the left wall and its *bottom* edge
		// is three units above the highest rock in the room (or the floor, if
		// there isn't one)
		var x int64 = 3
		var y int64 = height + int64(3+len(shape))

		// Simulate movement/falling of rock
		for {

			// Get the direction of next gas burst
			gas := patt[pi]
			pi++
			if pi >= len(patt) {
				pi = 0
			}

			// Move left/right according to gas burst, if possible
			assert(gas == '<' || gas == '>', "Bad pattern symbol!")
			if gas == '<' {
				if x > 1 && !occupied(x-1, y, shape) {
					x--
				}
			} else if gas == '>' {
				if x+int64(len(shape[0]))-1 < 7 && !occupied(x+1, y, shape) {
					x++
				}
			}

			// Fall if possible, stop this rock if not
			y-- // adjust y down
			if y-int64(len(shape)) < 0 || occupied(x, y, shape) {
				y++ // move back up
				placeShape(x, y, shape)
				if y > height {
					height = y
				}
				break
			}
		}

		// Part 1 is the answer at 2022 rocks, but continue simulation for part 2
		if rock == 2022 {
			fmt.Println("Part 1 (s/b 3068, 3114):", height)
		}

		// Show difference in height from previous iteration
		deltas = append(deltas, height-prevHeight)
		prevHeight = height
	}

	// Part 2: find repeating patterns, create output that be copied
	// to separate Python script
	// part2(deltas)  //  uncomment to run this (lots of noisy output)
}

// Part 2 is computed using a separate Python script, but this function
// finds the repeating blocks in sequence of height deltas created during
// the simulation. You need to copy the relevant numbers from the top
// of the output into the Python script to get the Part 2 answer.
// TODO: Find pattern and run calculations for Part 2 here, rather
// than in Python script.
func part2(deltas []int64) {

	// l is the length of the pattern, start small and adjust to
	// the value of the gaps between blocks found
	l := 1740 // we found 1740 for input data

	// Start looking for repeating patterns from the beginning of the deltas
	for from := 0; from < len(deltas); from++ { // pattern starts here
		seq := deltas[from : from+l] // sequence to look for
		var prevFind, finds, firstFind int
		gaps := []int{}                      // for gaps between sequences found
		for i := 0; i < len(deltas)-l; i++ { // start searching from the beginning
			this := deltas[i : i+l] // this extract
			if same(seq, this) {    // if it matches,
				if firstFind == 0 { // remember location of first match
					firstFind = i
				}
				if prevFind > 0 { // add distance from previous find to list of gaps
					gap := i - prevFind
					gaps = append(gaps, gap)
				}
				prevFind = i // remember location of last find
				finds++      // add up the number of finds
			}
		}

		// If more than one find, probably the pattern, so output info so
		//  it can be copied to the Python script
		if finds > 2 {
			fmt.Println("Sequence found", finds, "times:")
			fmt.Println("  starts at", from, ", first found at", firstFind)
			fmt.Println("  gaps =", gaps)
			fmt.Printf("  sequence is %d long, sum = %d\n", len(seq), sum(seq))
			fmt.Printf("  prefix is %d long, sum = %d\n", firstFind, sum(deltas[:firstFind]))
			fmt.Println("Sequence:", seq) // for Python script
		}
	}
}

// Check if position is occupied by a rock of given shape,
// i.e., if any point touches
func occupied(x, y int64, shape [][]int) bool {
	for sy := 0; sy < len(shape); sy++ {
		for sx := 0; sx < len(shape[0]); sx++ {
			if shape[sy][sx] == 1 && filled(x+int64(sx), y-int64(sy)) {
				return true
			}
		}
	}
	return false
}

// Place shape at position
func placeShape(x, y int64, shape [][]int) {
	for sy := 0; sy < len(shape); sy++ {
		for sx := 0; sx < len(shape[0]); sx++ {
			if shape[sy][sx] == 1 {
				p := Point{x + int64(sx), y - int64(sy)}
				chamber[p] = 1
			}
		}
	}
}

// Check if position is occupied
func filled(x, y int64) bool {
	_, ok := chamber[Point{int64(x), int64(y)}]
	return ok
}
