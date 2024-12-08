// Advent of Code 2024, Day 08
//
// Given letters on a grid, extrapolate the diagonal distance between each pair
// of the same letter, in either direction, and count up the total number of
// cells occupied by new entries. For Part 2, extrapolate in a line either
// direction.
//
// AK, 08 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Point struct {
	x, y int
}

var rows [][]byte

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	// Find all the locations of each letter
	space := map[byte][]Point{}
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			c := rows[y][x]
			if c != '.' {
				p := Point{x, y}
				space[c] = append(space[c], p)
			}
		}
	}

	// For each letter, look at each pair of locations, and
	// extrapolate to create new antinodes
	antinodes1 := map[Point]bool{} // part 1
	antinodes2 := map[Point]bool{} // part 2
	for _, points := range space {
		//fmt.Println(c, points)
		for i := 0; i < len(points); i++ {
			for j := 0; j < len(points); j++ {
				if j != i {

					// Pair of points
					p1 := points[i]
					p2 := points[j]

					// Calculate the dx/dy difference
					diff := Point{p2.x - p1.x, p2.y - p1.y}

					// Extrapolate from first point
					a1 := Point{p1.x - diff.x, p1.y - diff.y}
					if inrange(a1) {
						antinodes1[a1] = true
						antinodes2[a1] = true
					}

					// Extrapolate from second point
					a2 := Point{p2.x + diff.x, p2.y + diff.y}
					if inrange(a2) {
						antinodes1[a2] = true
						antinodes2[a2] = true
					}

					// For Part 2, also find all points on the line
					p := a1
					for inrange(p) {
						antinodes2[p] = true
						p = Point{p.x - diff.x, p.y - diff.y}
					}
					p = a1
					for inrange(p) {
						antinodes2[p] = true
						p = Point{p.x + diff.x, p.y + diff.y}
					}
				}
			}
		}
	}

	fmt.Println("Part 1 (s/b 14, 318) =", len(antinodes1))
	fmt.Println("Part 2 (s/b 34, 1126) =", len(antinodes2))

}

func inrange(p Point) bool {
	return p.y >= 0 && p.y < len(rows) && p.x >= 0 && p.x < len(rows[0])
}
