// Advent of Code 2022, Day 08
//
// Given a topographical map (matrix) of tree heights, count the number of
// trees that have visibility all the way to the edge (Part 1), and the
// highest "visibility" score of any tree, where that score is the product
// of the numbers of trees less than the current tree in each direction.
//
// AK, 8 Dec 2022

package main

import (
	"fmt"
)

func main() {

	// Read the input file
	//lines := readLines("sample.txt")
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Part 1: count how many trees are "visible"
	fmt.Println("Part 1 (s/b 21 or 1763):", part1(lines))

	// Part 2: maximum scenic score
	// 560 and 14400 both too low
	fmt.Println("Part 2 (s/b 8 or 671160):", part2(lines))
}

// Part 1: how many trees are visible?
func part1(lines []string) int {
	nr := len(lines)
	nc := len(lines[0])
	nvis := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if isVisible(r, c, lines) {
				nvis++
			}
		}
	}
	return nvis
}

// For Part 1, is tree in the given position "visible" from any direction?
func isVisible(r, c int, forest []string) bool {

	// Trees on the edges are always visible
	nr := len(forest)
	nc := len(forest[0])
	if r == 0 || r == nr-1 || c == 0 || c == nc-1 {
		return true
	}

	// Otherwise, look up & down left & right, to see
	// if there is an "opening" in any direction, i.e.,
	// all trees in that direction are lower than this one
	x := forest[r][c] // this "tree"

	visLeft := true
	for i := 0; i < c; i++ { // left
		if forest[r][i] >= x {
			visLeft = false
		}
	}

	visRight := true
	for i := c + 1; i < nc; i++ { // right
		if forest[r][i] >= x {
			visRight = false
		}
	}

	visAbove := true
	for i := 0; i < r; i++ {
		if forest[i][c] >= x { // above
			visAbove = false
		}
	}

	visBelow := true
	for i := r + 1; i < nr; i++ {
		if forest[i][c] >= x { // below
			visBelow = false
		}
	}

	// A tree is visible if not blocked from all directions
	return visLeft || visRight || visAbove || visBelow
}

// Part 2: maximum "score" of tree visibility
func part2(lines []string) int {
	score := 0
	nr := len(lines)
	nc := len(lines[0])
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			ss := scenicScore(r, c, lines)
			if ss > score {
				score = ss
			}
		}
	}
	return score
}

// For Part 1, is tree in the given position "visible"
func scenicScore(r, c int, forest []string) int {

	// If on the edge, score is zero
	if r == 0 || r == nr-1 || c == 0 || c == nc-1 {
		return 0
	}

	// Count how many trees are visible in each direction, by extracting the
	// trees from the current tree in each direction, then checking that sequence
	nr := len(forest)
	nc := len(forest[0])

	seq := []byte{}
	for i := c; i >= 0; i-- { // left
		seq = append(seq, forest[r][i])
	}
	visLeft := nVisible(seq)

	seq = []byte{}
	for i := c; i < nc; i++ { // right
		seq = append(seq, forest[r][i])
	}
	visRight := nVisible(seq)

	seq = []byte{}
	for i := r; i >= 0; i-- { // up
		seq = append(seq, forest[i][c])
	}
	visUp := nVisible(seq)

	seq = []byte{}
	for i := r; i < nr; i++ { // down
		seq = append(seq, forest[i][c])
	}
	visDown := nVisible(seq)

	// Return product to get score
	return visLeft * visRight * visUp * visDown
}

// Count the number of trees visible from the first tree in given sequence,
// i.e., not higher than the starting tree
func nVisible(trees []byte) int {
	n := 0
	for i := 1; i < len(trees); i++ {
		n++
		if trees[i] >= trees[0] {
			break
		}
	}
	return n
}
