// Advent of Code 2020, Day 20
//
// Assemble a set of "tiles" into an image, so adjacent edges match, flipping or
// rotating as necessary. Part 1 is the product of the IDs of the corner tiles.
// For Part 2, strip the edges of the tiles and assemble them into a combined
// image, then search for a 3-line "sea monster" pattern (flipping and rotating
// the combined image as required), and report the number of hash marks in the
// image that are not covered up by the "sea monsters" found.
//
// AK, 23/11/2022

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Tile struct {
	number int64    // the ID of this tile
	rows   []string // the original rows of data for this tile
	// For part 1
	edges    []string // the top, right, bottom, left edges
	reversed []string // the same edges, reversed
	// For part 2
	top, right, bottom, left string   // the top, right, bottom, left edges
	placed                   bool     // true when placed in position
	position                 Position // position when placed
}

type Position struct {
	row, col int
}

func main() {

	//// Read tiles, extract the edges
	//tiles := readTiles("sample.txt")
	tiles := readTiles("input.txt")

	// Part 1: find matching edges, report product of corner tile IDs
	part1(tiles)

	// Part 2: piece together the entire image, and search for pattern
	part2(tiles)
}

// Part 1: find the four corner tiles with only 2 edges that match, and
// multiply their IDs together (simple algorithm, finds tiles with only
// two edges that match other tiles, does not try to piece together the
// whole picture)
func part1(tiles []Tile) {

	// Look at each pair of tiles
	var result int64 = 1
	nmatches := 0
	for i := 0; i < len(tiles); i++ {

		// Check this tile's edges against the edges (original and
		// reversed) of all other tiles
		t1 := tiles[i]
		matches := 0
		for j := 0; j < len(tiles); j++ {

			if i == j {
				continue
			}
			t2 := tiles[j]

			// Check each edge of tile 1 against this tile
			for _, e := range t1.edges {
				if isIn(e, t2.edges) || isIn(e, t2.reversed) {
					matches++
				}
			}
		}

		// Report the number of edge matches for this tile
		fmt.Printf("Tile %d: %d edges match other tiles\n", t1.number, matches)

		// Multiply ID if 2 matches (there should be 4 of these tiles)
		if matches == 2 {
			fmt.Println("  Multiplying by", t1.number)
			result *= t1.number
			nmatches++
		}
	}

	// Report the final result
	fmt.Println("Result =", result)
	if nmatches != 4 {
		fmt.Println("Warning: should be 4 matches, but there are", nmatches)
	}
}

// Part 2: piece together the whole image, flipping or rotating tiles as
// necessary to make the edges match. Then, search for a pattern within
// the combined image.
func part2(tiles []Tile) {

	// Start with the first tile, put it at position 0,0
	tiles[0].position = Position{0, 0}
	tiles[0].placed = true

	// Get an unplaced tile, and try to fit it next to already placed tiles,
	// flipping or rotating as necessary
	flips := []string{"none", "horizontal", "vertical", "both"} // possible flips
	rots := []int{0, 90, 180, 270}                              // possible rotations
	for {

		// Stop if no more unplaced tiles
		done := true
		for ui := 0; ui < len(tiles); ui++ {
			if !tiles[ui].placed {
				done = false
				break
			}
		}
		if done {
			break
		}

		// Go through tiles, process each unplaced one
		for ui := 0; ui < len(tiles); ui++ {

			// Skip placed tiles in the outer loop
			if tiles[ui].placed {
				continue
			}

			fmt.Println("Checking unplaced tile:", tiles[ui].number)

			// For the unplaced tile, try to find a placed tile that matches any
			// edge, in any configuration
			for i := 0; i < len(tiles); i++ {

				// Find the next already-placed tile
				t := &tiles[i]
				if !t.placed {
					continue
				}

				// Try every configuration of the unplaced tile, i.e., as-is,
				// flipped none/horiz/vertically, rotated 0/90/180/270 degrees,
				// and see if it matches any of the four edges of the placed
				// tile
				// TODO: only check positions that are not yet occupied
				for _, flip := range flips {
					for _, rot := range rots {
						u1 := flipRotate(&tiles[ui], flip, rot)
						if u1.left == t.right { // place to right
							u1.position = Position{t.position.row, t.position.col + 1}
							u1.placed = true
							tiles[ui] = u1 // save rotated state
						} else if u1.right == t.left { // place to left
							u1.position = Position{t.position.row, t.position.col - 1}
							u1.placed = true
							tiles[ui] = u1 // save rotated state
						} else if u1.bottom == t.top { // place above
							u1.position = Position{t.position.row - 1, t.position.col}
							u1.placed = true
							tiles[ui] = u1 // save rotated state
						} else if u1.top == t.bottom { // place below
							u1.position = Position{t.position.row + 1, t.position.col}
							u1.placed = true
							tiles[ui] = u1 // save rotated state
						}
						if tiles[ui].placed {
							u1 := tiles[ui]
							fmt.Printf("Placed %d in flip %s, rot %d at %d,%d\n",
								u1.number, flip, rot, u1.position.row, u1.position.col)
							break
						}
					}
					if tiles[ui].placed {
						break
					}
				}
				if tiles[ui].placed {
					break
				}
			}
			if tiles[ui].placed {
				break
			}
		}

	}

	// Sort the tiles by row, column
	sort.Slice(tiles, func(a, b int) bool {
		if tiles[a].position.row < tiles[b].position.row {
			return true
		} else if tiles[a].position.row > tiles[b].position.row {
			return false
		} else {
			return tiles[a].position.col < tiles[b].position.col
		}
	})

	// Show final positions of the tiles
	fmt.Println("Final tile positions:")
	for _, t := range tiles {
		fmt.Println(t.number, t.position)
	}

	// Strip the border off each tile image
	for i := 0; i < len(tiles); i++ {
		stripImage(&tiles[i])
	}

	// Combine the correctly positioned tiles to create a single large image
	img := []string{}
	for i, t := range tiles {
		if i == 0 || t.position.row != tiles[i-1].position.row {
			for _, r := range t.rows { // new tile row: add rows to end
				img = append(img, r)
			}
		} else {
			for j, r := range t.rows { // append to existing tile rows
				ir := len(img) - len(t.rows) + j
				img[ir] += r
			}
		}
	}

	fmt.Println("Final combined image:")
	for _, r := range img {
		fmt.Println(r)
	}

	// The pattern we're looking for
	pattern := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   "}

	// Search for pattern in each permutation of the image, and
	// calculate the number of hashes that are not part of patterns
	picHashes := countHashes(img)
	pattHashes := countHashes(pattern)
	for _, flip := range flips {
		for _, rot := range rots {
			t0 := Tile{number: 0, rows: img}
			t1 := flipRotate(&t0, flip, rot)
			img1 := t1.rows
			n := countPatterns(img1, pattern)
			fmt.Printf("Patterns found in image %s %d: %d\n", flip, rot, n)
			if n > 0 {
				fmt.Printf("  Answer = %d - %d * %d = %d\n", picHashes,
					pattHashes, n, picHashes-n*pattHashes)
			}
		}
	}
}

// Strip border from a tile's image
func stripImage(t *Tile) {
	rows := []string{}
	for i := 1; i < len(t.rows)-1; i++ {
		r := t.rows[i]
		r = r[1 : len(r)-1]
		rows = append(rows, r)
	}
	t.rows = rows
}

// Count occurrences of pattern in an image
func countPatterns(img, pattern []string) int {
	n := 0
	for r := 0; r < len(img); r++ {
		for c := 0; c < len(img[0]); c++ {
			if seaMonsterAt(img, pattern, r, c) {
				n++
			}
		}
	}
	return n
}

// Determine if there is a "sea monster" pattern at the given position
// "                  #"
// "#    ##    ##    ###"
// " #  #  #  #  #  #"
func seaMonsterAt(img, pattern []string, r, c int) bool {

	// Pattern must fit into image at this location
	if c+len(pattern[0]) >= len(img[0]) {
		return false
	}
	if r+len(pattern) >= len(img) {
		return false
	}

	// Check each row of the pattern against the image
	for pr := 0; pr < len(pattern); pr++ {
		for pc := 0; pc < len(pattern[pr]); pc++ {
			pb := pattern[pr][pc]       // char we're looking for
			ib := img[r+pr][c+pc]       // char in image
			if pb == '#' && ib != '#' { // fail if hashes don't match
				return false
			}
		}
	}
	return true
}

// Count the number of hashes in a list of strings
func countHashes(ss []string) int {
	n := 0
	for _, s := range ss {
		for i := 0; i < len(s); i++ {
			if s[i] == '#' {
				n++
			}
		}
	}
	return n
}

// Flip/rotate a tile, return a new copy
func flipRotate(t *Tile, flip string, degrees int) Tile {

	// Make a copy of the tile (edges assigned later, so leave blank)
	t1 := Tile{number: t.number, placed: t.placed, position: t.position}
	t1.rows = append(t1.rows, t.rows...)

	// Flip Horizontal: reverse the characters in each row
	if flip == "horizontal" || flip == "both" {
		for i := 0; i < len(t1.rows); i++ {
			t1.rows[i] = reverse(t1.rows[i])
		}
	}

	// Flip Vertical: just reverse the rows
	if flip == "vertical" || flip == "both" {
		rows := make([]string, len(t1.rows), len(t1.rows))
		for i := 0; i < len(t1.rows); i++ {
			rows[len(t1.rows)-i-1] = t1.rows[i]
		}
		t1.rows = rows
	}

	// Rotate (transpose) the rows, 90 degrees each time
	for degrees > 0 {
		t1.rows = transpose(t1.rows)
		degrees -= 90
	}

	// Extract the edges
	extractEdges(&t1)

	return t1
}

// Transpose (rotate) a list of strings 90 degrees
// 1. Extract each column, make each a row (same order)
// 2. Reverse each row
func transpose(rows []string) []string {
	res := []string{}
	for c := 0; c < len(rows[0]); c++ { // each col
		row := []byte{}
		for r := len(rows) - 1; r >= 0; r-- { // each row, in reverse
			row = append(row, rows[r][c])
		}
		res = append(res, string(row))
	}
	return res
}

// Read tiles, extract the edges
func readTiles(filename string) []Tile {

	// Read input file, split into lines
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	// Parse out separate tiles
	var tiles []Tile // list of tiles
	var t Tile       // current tile
	for _, l := range lines {

		// Start a new tile
		if strings.HasPrefix(l, "Tile ") {
			if t.number > 0 { // i.e., a tile has been defined, add to list
				tiles = append(tiles, t)
			}
			tnum, _ := strconv.ParseInt(l[5:len(l)-1], 10, 64)
			t = Tile{number: tnum}
		} else if len(l) > 0 { // or add row to current one
			t.rows = append(t.rows, l)
		}
	}

	// Add last tile
	if t.number > 0 {
		tiles = append(tiles, t)
	}

	// Extract the edges from each tile
	// For part 1, add lists of edges and reversed  edges
	for ti := 0; ti < len(tiles); ti++ {
		t := &tiles[ti]
		extractEdges(t)
		// For part 1:
		t.edges = append(t.edges, t.top, t.right, t.bottom, t.left)
		for _, e := range t.edges {
			t.reversed = append(t.reversed, reverse(e))
		}
	}

	return tiles
}

// Extract edges from a tile, updating the fields in the tile itself
func extractEdges(t *Tile) {

	// Top and bottom are just the first and last rows
	t.top = t.rows[0]
	t.bottom = t.rows[len(t.rows)-1]

	// Left and right are the first and last columns
	var left, right []byte
	for i := 0; i < len(t.rows); i++ {
		left = append(left, t.rows[i][0])
		right = append(right, t.rows[i][len(t.rows[i])-1])
	}
	t.left = string(left)
	t.right = string(right)
}

// Reverse a string
func reverse(s string) string {
	r := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		r[len(s)-i-1] = s[i]
	}
	return string(r)
}

// Is string in a list?
func isIn(s string, ss []string) bool {
	for _, x := range ss {
		if x == s {
			return true
		}
	}
	return false
}
