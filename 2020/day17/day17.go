// Advent of Code 2020, Day 17
//
// Input is a set of "cubes" in 2-d space, either on or off. For part 1,
// this is extended to 3-d space, for part 2 4-d space. Simulate a set of
// simple rules, depending on current state of a cube and the number of "on"
// neighbours it has. Simulation is supposed to occur "simulataneously", so
// apply changes to future state, then roll them to current state after each
// iteration. Part 2 is a trivial set of modifications to Part 1, to make it
// 4-d instead of 3-d, so only the Part 2 code is shown here.
//
// AK, 17/10/2022

package main

import (
	"fmt"
	"io/ioutil"
)

// One point in 4-d space (use 3-d for part 1)
type Point struct {
	x, y, z, h int // can be negative
}

// The current and next state of each known point (unknown is off)
var current map[Point]int
var next map[Point]int

func main() {

	// Initialize global state maps
	current = map[Point]int{}
	next = map[Point]int{}

	// Read data set and convert to a set of points
	data, _ := ioutil.ReadFile("input.txt") // or sample.txt
	var x, y, z, h int                      // z and h are zero in input
	for _, b := range data {
		if b == '\n' { // start next row
			y++
			x = 0
		} else if b == '#' { // hash means on
			setCurrentState(x, y, z, h, 1)
			x++
		} else if b == '.' { // period means off
			//setCurrentState(x, y, 0, 0) // not really necessary
			x++
		} else {
			fmt.Println("Unknown character, ignoring:", b)
			x++
		}
	}

	// Run each iteration
	for iter := 1; iter <= 6; iter++ {

		// Look at each cube in current space, including 1 past current edge
		fmt.Println("Iteration", iter)
		min, max := getDims()
		for x := min.x - 1; x <= max.x+1; x++ {
			for y := min.y - 1; y <= max.y+1; y++ {
				for z := min.z - 1; z <= max.z+1; z++ {
					for h := min.h - 1; h <= max.h+1; h++ {

						// Get current state and number of active neighbors
						state := getCurrentState(x, y, z, h)
						nactive := activeNeighbours(x, y, z, h)

						// If a cube is active and exactly 2 or 3 of its neighbors
						// are also active, the cube remains active. Otherwise, the
						// cube becomes inactive.
						if state == 1 {
							if !(nactive == 2 || nactive == 3) {
								setNextState(x, y, z, h, 0)
							}
						}

						// If a cube is inactive but exactly 3 of its neighbors
						// are active, the cube becomes active. Otherwise, the cube
						// remains inactive.
						if state == 0 {
							if nactive == 3 {
								setNextState(x, y, z, h, 1)
							}
						}
					}
				}
			}
		}

		// After each iteration, roll over the next states back to the current
		rollOver()
	}

	// Count the number of active cubes
	// For Part 1, sample should be 112 after 6 iterations, input 336
	// For Part 2, 848 and 2620
	tot := 0
	for _, n := range current {
		tot += n
	}
	fmt.Printf("Part 2: %d active cubes\n", tot)
}

// Get current 1/0 state of cube at specific x/y/z/h
func getCurrentState(x, y, z, h int) int {
	c, ok := current[Point{x, y, z, h}]
	if ok {
		return c
	} else {
		return 0 // assume inactive if not defined
	}
}

// Set current 1/0 state of cube at specific x/y/z/h
func setCurrentState(x, y, z, h, state int) {
	current[Point{x, y, z, h}] = state
}

// Set next 1/0 state of cube at specific x/y/z/h
func setNextState(x, y, z, h, state int) {
	next[Point{x, y, z, h}] = state
}

// Roll next states over to current
func rollOver() {

	// Copy the changed values to the current map. Do NOT clear out
	// the current map first, as this would only copy things that changed
	// in the last iteration
	for p, st := range next {
		current[p] = st
	}

	// Clear out the map of change for the next iteration
	next = map[Point]int{}
}

// Get the current number of active neighbours for an x/y/z/h coordinate.
// Basically just look -1/0/1 in each dimension, but don't include the
// central cube itself.
func activeNeighbours(x, y, z, h int) int {
	diffs := []int{-1, 0, 1}
	var nactive int
	for _, dx := range diffs {
		for _, dy := range diffs {
			for _, dz := range diffs {
				for _, dh := range diffs {
					if !(dx == 0 && dy == 0 && dz == 0 && dh == 0) {
						nactive += getCurrentState(x+dx, y+dy, z+dz, h+dh)
					}
				}
			}
		}
	}
	return nactive
}

// Get dimensions of the data set, i.e., the min/max x/y/z of all currently defined points
func getDims() (Point, Point) {

	var min, max Point
	for p, _ := range current {

		// Min/max x values
		if p.x < min.x {
			min.x = p.x
		}
		if p.x > max.x {
			max.x = p.x
		}

		// Min/max y values
		if p.y < min.y {
			min.y = p.y
		}
		if p.y > max.y {
			max.y = p.y
		}

		// Min/max z values
		if p.z < min.z {
			min.z = p.z
		}
		if p.z > max.z {
			max.z = p.z
		}

		// Min/max h values
		if p.h < min.h {
			min.h = p.h
		}
		if p.h > max.h {
			max.h = p.h
		}
	}

	return min, max
}
