// Advent of Code 2021, Day 22
//
// Turn on/off points in space, defined by 3-d ranges (like rectangular
// cubes). Part 1 was done by brute force, had to replace with recursive
// calculation of volumes minus intersections for Part 2, due to memory
// constraints.
//
// AK, 22 Dec 2021 (Part 1) and 28 Dec 2021 (Part 2)

package main

import (
	"fmt"
)

// Part 1: Do brute-force simulation only within -50..50 subspace
func part1(instructions []Cube) int64 {

	// Set up a space with the centres of all 3x3 cubes, recording
	// the state of that cube
	space := map[Point]int32{}

	// Apply the instructions to the space
	for _, i := range instructions {
		fmt.Println("Processing", i)
		if i.x0 > i.x1 || i.y0 > i.y1 || i.z0 > i.z1 {
			panic("Malformed instruction")
		}
		for x := max(-50, i.x0); x <= min(50, i.x1); x++ {
			for y := max(-50, i.y0); y <= min(50, i.y1); y++ {
				for z := max(-50, i.z0); z <= min(50, i.z1); z++ {
					p := Point{x, y, z}
					space[p] = i.onOff
				}
			}
		}
	}

	// Now count up the number of on states
	var n int64 = 0
	for _, v := range space {
		n += int64(v)
	}
	return n
}

// Part 2: Recursively calculate volumes, minus any intersections
func part2(instructions []Cube) int64 {

	var total int64 = 0

	// Start with the last cube and work backwards
	for i := len(instructions) - 1; i >= 0; i-- {

		// Don't bother if this cube is "off"
		cube := instructions[i]
		if cube.onOff == 0 {
			continue
		}

		// Enumerate all intersections between this cube and any others,
		// moving forward from this cube
		intersects := []Cube{}
		for j := i + 1; j < len(instructions); j++ {
			cube1 := instructions[j]
			intersect := getIntersection(cube, cube1)
			if intersect != nil { // no intersection, skip
				intersect.onOff = 1 // so we count the volume
				intersects = append(intersects, *intersect)
			}
		}

		// Add volume of this cube, but without the intersections
		total = total + volume(cube) - part2(intersects)
	}

	return total
}

// Get a cube that is the intersection between two cubes
func getIntersection(a, b Cube) *Cube {

	x0 := max(a.x0, b.x0)
	x1 := min(a.x1, b.x1)
	y0 := max(a.y0, b.y0)
	y1 := min(a.y1, b.y1)
	z0 := max(a.z0, b.z0)
	z1 := min(a.z1, b.z1)
	if x1 < x0 || y1 < y0 || z1 < z0 {
		return nil
	}

	return &Cube{0, x0, x1, y0, y1, z0, z1}
}

// Calculate the volume of a cube
func volume(c Cube) int64 {
	return int64(c.x1-c.x0+1) * int64(c.y1-c.y0+1) * int64(c.z1-c.z0+1)
}

// Simple integer min/max
func min(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int32) int32 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	// Read instructions from a file
	//fname := "sample1.txt" // for part 1
	//fname := "sample2.txt" // for part 1
	fname := "sample3.txt" // for part 2
	fname = "input.txt"    // for part 2
	fmt.Println("Reading", fname)
	instructions := readData(fname)

	// Part 1: Do the simulation
	//n := part1(instructions)
	//fmt.Printf("Part 1: %d cubes are on (s/b be 39 for sample1, 590784 for sample2)\n", n)

	// Part 2: Do the simulation
	n := part2(instructions)
	var sb int64 = 2758514936282235 // expected result for sample3
	delta := n - sb
	frac := float64(delta) / float64(sb) * 100
	fmt.Printf("Part 2: %d cubes are on (should be %d for sample3)\n", n, sb)
	fmt.Printf("Delta = %d, %f%%\n", delta, frac)
}
