// Advent of Code 2020, Day 24
//
// Given black/white tiles on a hexoganal grid, follow set of movement
// directions and flip over tiles, then count number of black tiles.
// For Part 2, simulate 100 days of flipping tiles based on state and
// number of adjacent black tiles.
//
// Useful: https://www.redblobgames.com/grids/hexagons/
//
// AK, 5/09/2023

package main

import (
	"fmt"
	"os"
	"strings"
)

// A point is a coordinate in a hexagonal grid
type Point struct {
	x, y int
}

func main() {

	// Read input file
	fname := "sample.txt"
	fname = "input.txt"
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	// Convert to string and split lines
	lines := strings.Split(string(data), "\n")

	// Create an empty map of coordinates
	coords := map[Point]int{} // empty map of coords

	// Do part 1: follow instructions to flip tiles
	for _, line := range lines {

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Start at 0,0 and move the point according to the instructions
		insts := parseLine(line)
		p := Point{0, 0}
		for _, inst := range insts {
			p = move(p, inst)
		}

		// Flip tile at this location
		coords[p] = 1 - coords[p]
	}

	// Part 1: show number of black tiles
	fmt.Println("Part 1 (s/b 10 or 266):", sum(coords))

	// For part 2, simulate 100 days:
	// 1. Any black tile with zero or more than 2 black tiles
	//    immediately adjacent to it is flipped to white.
	// 2. Any white tile with exactly 2 black tiles immediately adjacent
	//    to it is flipped to black.
	// The rules are applied simultaneously to every tile; put another
	// way, it is first determined which tiles need to be flipped, then
	// they are all flipped at the same time.
	for day := 0; day < 100; day++ {

		// Accumulate changes based on state of tile and number of adjacent black tiles
		changes := map[Point]int{}     // changes to be applied at end of day
		for x := -100; x <= 100; x++ { // every tile on the floor, make
			for y := -100; y <= 100; y++ { // big enough to cover all points
				p := Point{x, y}
				nblack := countNeighbors(coords, p)
				if coords[p] == 1 && (nblack == 0 || nblack > 2) {
					changes[p] = 0
				}
				if coords[p] == 0 && nblack == 2 {
					changes[p] = 1
				}
			}
		}

		// Apply changes at end of each day
		for p, c := range changes {
			coords[p] = c
		}
	}

	// Part 2: count up the black tiles
	fmt.Println("Part 2 (s/b 2208 or 3627):", sum(coords))
}

// Sum up the values of a map
func sum(coords map[Point]int) int {
	count := 0
	for _, v := range coords {
		count += v
	}
	return count
}

// Convert a line of instructins to a list of directions
func parseLine(line string) []string {
	var directions []string
	for i := 0; i < len(line); i++ {
		if line[i] == 'e' || line[i] == 'w' {
			directions = append(directions, line[i:i+1])
		} else {
			directions = append(directions, line[i:i+2])
			i++
		}
	}
	return directions
}

// Move a point in a hexagonal grid
func move(p Point, dir string) Point {
	switch dir {
	case "e":
		p.x++
	case "w":
		p.x--
	case "ne":
		p.x++
		p.y--
	case "nw":
		p.y--
	case "se":
		p.y++
	case "sw":
		p.x--
		p.y++
	}
	return p
}

// Count the number of neighbors of a point that are black
func countNeighbors(coords map[Point]int, p Point) int {
	count := 0
	for _, dir := range []string{"e", "w", "ne", "nw", "se", "sw"} {
		n := move(p, dir)
		count += coords[n]
	}
	return count
}
