// Advent of Code 2021, Day 19
//
// Match up 3-dimensional cubes in space, adjusting x,y,z offset and also
// rotation along any 3 axes, so that at least 12 points in the each
// pair of cubes line up exactly. Then, calculate the maximum distance
// between the cubes.
//
// AK, 19/12/2021

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A point is just a slice with x, y, z coords
type P []int64

// Read "beacons" from input file
func readData(fname string) [][]P {

	result := [][]P{}
	var curBlock []P

	// Open file
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Could not open ", fname)
		return nil
	}

	// Read each line of input file
	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {

		// --- starts a new block
		lineNo++
		t := strings.TrimSpace(scanner.Text())
		if len(t) < 3 {
			continue
		} else if string(t[:3]) == "---" {
			if curBlock != nil {
				result = append(result, curBlock)
			}
			curBlock = nil //[]P{}
		} else {
			r := make(P, 3, 3)
			parts := strings.Split(t, ",")
			if len(parts) != 3 {
				fmt.Println("Invalid line", lineNo)
				return nil
			}
			r[0], _ = strconv.ParseInt(parts[0], 10, 64)
			r[1], _ = strconv.ParseInt(parts[1], 10, 64)
			r[2], _ = strconv.ParseInt(parts[2], 10, 64)
			curBlock = append(curBlock, r)

		}
		lineNo += 1
	}
	result = append(result, curBlock)
	return result
}

// Find the x, y and z offsets that create the most overlaps, using
// a simple brute force approach of forcing each pair of points into
// alignment, then counting the overlaps implied by the resulting offsets.
// Returns the x, y, and z offsets that need to be applied, and the number
// of points that are lined up as a result.
func minimizeDistance2(s0 []P, s1 []P) (int64, int64, int64, int64) {

	var bestOffX int64 = -1
	var bestOffY int64 = -1
	var bestOffZ int64 = -1
	var bestOverlaps int64 = 0

	// Match each point in list 1 with each point in list 2,
	// and calculate and test the implied offset
	for i0, p0 := range s0 {
		for i1, p1 := range s1 {
			if i0 == i1 { // don't compare block to itself
				continue
			}
			ox := p1[0] - p0[0]
			oy := p1[1] - p0[1]
			oz := p1[2] - p0[2]
			olaps := countOverlaps(s0, s1, ox, oy, oz)
			if olaps > bestOverlaps {
				bestOverlaps = olaps
				bestOffX = ox
				bestOffY = oy
				bestOffZ = oz
			}
		}
	}
	return bestOffX, bestOffY, bestOffZ, bestOverlaps
}

// Count the number of overlaps between two lists of points, given the
// offset added to the second one
func countOverlaps(s0 []P, s1 []P, xOffset, yOffset, zOffset int64) int64 {
	var n int64
	for _, p0 := range s0 {
		for _, p1 := range s1 {
			if p0[0] == p1[0]-xOffset && p0[1] == p1[1]-yOffset && p0[2] == p1[2]-zOffset {
				n += 1
			}
		}
	}
	return n
}

// Adjust a block by applying the given offsets
func applyOffsets(s0 []P, xOffset, yOffset, zOffset int64) []P {
	s1 := make([]P, len(s0), len(s0))
	for i, p := range s0 {
		p1 := P{p[0] - xOffset, p[1] - yOffset, p[2] - zOffset}
		s1[i] = p1
	}
	return s1
}

// Compare a block with every other one (in every rotation), and find
// the first block that has at least 12 overlaps with the selected
// block, returning:
// - index of the block found
// - the block rotated to the required position
// - the x,y,z offsets applied
func findMatch(blocks [][]P, b0 []P) (int, []P, int64, int64, int64) {

	// Look at each other block, and compare  it to theselected one, looking
	// for a rotation/offset combination that results in at least 12 overlaps
	for i := 0; i < len(blocks); i++ {

		// Search the rotations for matches
		b1 := blocks[i]
		rotations := rotateAllWays(b1)
		for rot := 0; rot < len(rotations); rot++ {
			b := rotations[rot]
			xoff, yoff, zoff, olaps := minimizeDistance2(b0, b)
			if olaps >= 12 {
				return i, b, xoff, yoff, zoff
			}
		}
	}

	// Return all -1 if no solution found
	return -1, nil, -1, -1, -1
}

// Simple integer-only abs
func abs(a int64) int64 {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}

func main() {

	// Read data
	//f := "sample.txt"
	f := "input.txt"
	blocks := readData(f)
	if blocks == nil {
		return
	}
	fmt.Printf("Read %s, %d blocks\n", f, len(blocks))

	// Iterate as follows:
	// - take a block that is "frozen" (e.g., the first, or one that has
	//   already been matched to another block)
	// - find another "unfrozen" block that matches it, rotating it as required
	//   and getting the necessary offsets
	// - apply the offsets to the block, and add it to the "done" list
	//
	// Once a block is matched to another one, those blocks are "frozen" and
	// should not be adjusted
	frozen := make([]int, len(blocks), len(blocks))
	frozen[0] = 1                    // first one is frozen by definition
	scannerCoords := []P{P{0, 0, 0}} // offsets found, first is at zero
	for true {

		// Make lists of the frozen and unfrozen blocks
		frozenBlocks := [][]P{}
		unfrozenBlocks := [][]P{}
		unfrozenIndices := []int{}
		for i := 0; i < len(blocks); i++ {
			if frozen[i] == 1 {
				frozenBlocks = append(frozenBlocks, blocks[i])
			} else {
				unfrozenBlocks = append(unfrozenBlocks, blocks[i])
				unfrozenIndices = append(unfrozenIndices, i)
			}
		}

		// Stop if there are no more unfrozen blocks
		if len(unfrozenBlocks) == 0 {
			break
		}

		// Always search from frozen blocks towards unfrozen ones
		for i, fb := range frozenBlocks {

			// Try to find a match for this block among the unfrozen blocks
			bi, rotated, xoff, yoff, zoff := findMatch(unfrozenBlocks, fb)
			if bi == -1 {
				continue
			}

			// Adjust the unfrozen block with offsets, and the block in the list
			// with the rotated and offset version
			fmt.Printf("Frozen block %d matched to unfrozen block %d, offset %d,%d,%d\n", i, bi, xoff, yoff, zoff)
			adjusted := applyOffsets(rotated, xoff, yoff, zoff)
			fmt.Printf("  %d overlaps (should be 12+)\n", countOverlaps(fb, adjusted, 0, 0, 0))
			origInd := unfrozenIndices[bi] // get index in the original list
			blocks[origInd] = adjusted
			frozen[origInd] = 1

			// For part 2, add the offsets for each device
			scannerCoords = append(scannerCoords, P{-xoff, -yoff, -zoff})
		}
	}

	// Finally go through all the blocks, and count up the number of unique
	// coordinates, these are the beacons for Part 1
	points := []P{}
	for _, b := range blocks { // each block
		for _, p := range b { // each adjusted point within block
			if !containsVector(points, p) {
				points = append(points, p)
			}
		}
	}
	fmt.Printf("%d unique points found (should be 79 for sample)\n", len(points))

	// Part 2: Using the Manhattan distance, how far apart do the scanners get?
	// In the above example, scanners 2 (1105,-1205,1229) and 3 (-92,-2380,-20)
	// are the largest Manhattan distance apart. In total, they are
	// 1197 + 1175 + 1249 = 3621 units apart.
	// What is the largest Manhattan distance between any two scanners?
	var maxDist int64 = 0
	for _, s1 := range scannerCoords {
		for _, s2 := range scannerCoords {
			dist := abs(s1[0]-s2[0]) + abs(s1[1]-s2[1]) + abs(s1[2]-s2[2])
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	fmt.Printf("Max distance = %d (s/b 3621 for sample)\n", maxDist)
}
