// Earlier approach, using optimization to find the best fit, did not work

package main

import (
	"fmt"
)

// Find the x, y and z offsets that minimize the total Manhattan distance
// between two lists of points
func minimizeDistance(s0 []P, s1 []P) (int64, int64, int64) {

	// Find the x offset first
	var xoff, incr int64
	bestDist := totDist(s0, s1, 0, 0, 0) // starting distance with no offset
	if totDist(s0, s1, 1, 0, 0) < bestDist {
		incr = 1 // increasing improves
	} else if totDist(s0, s1, -1, 0, 0) < bestDist {
		incr = -1 // decreasing will improve
	} else {
		incr = 0 // no improvement possible from zero
	}

	// Keep incrementing until no more improvement
	d := totDist(s0, s1, xoff+incr, 0, 0)
	for incr != 0 && d < bestDist {
		xoff += incr
		bestDist = d
		d = totDist(s0, s1, xoff+incr, 0, 0)
	}

	// Find the y offset next
	var yoff int64
	bestDist = totDist(s0, s1, xoff, 0, 0) // starting distance with no offset
	if totDist(s0, s1, xoff, 1, 0) < bestDist {
		incr = 1 // increasing improves
	} else if totDist(s0, s1, xoff, -1, 0) < bestDist {
		incr = -1 // decreasing will improve
	} else {
		incr = 0 // no improvement possible from zero
	}

	// Keep incrementing until no more improvement
	d = totDist(s0, s1, xoff, yoff+incr, 0)
	for incr != 0 && d < bestDist {
		yoff += incr
		bestDist = d
		d = totDist(s0, s1, xoff, yoff+incr, 0)
	}

	// Find the z offset next
	var zoff int64
	bestDist = totDist(s0, s1, xoff, yoff, 0) // starting distance with no offset
	if totDist(s0, s1, xoff, yoff, 1) < bestDist {
		incr = 1 // increasing improves
	} else if totDist(s0, s1, xoff, yoff, -1) < bestDist {
		incr = -1 // decreasing will improve
	} else {
		incr = 0 // no improvement possible from zero
	}

	// Keep incrementing until no more improvement
	d = totDist(s0, s1, xoff, yoff, zoff+incr)
	for incr != 0 && d < bestDist {
		zoff += incr
		bestDist = d
		d = totDist(s0, s1, xoff, yoff, zoff+incr)
	}

	// Try a few points around optimimum to make sure we can't improve it
	d = totDist(s0, s1, xoff, yoff, zoff)
	tests := []int64{-1500, -1000, -500, -100, -50, -10, -1, 1, 10, 50, 100, 500, 1000, 1500}
	for _, dx := range tests {
		for _, dy := range tests {
			for _, dz := range tests {
				d1 := totDist(s0, s1, xoff+dx, yoff+dy, zoff+dz)
				if d1 < d {
					fmt.Printf("WARNING: lower distance found at %d,%d,%d + %d,%d,%d (%d vs %d)\n",
						xoff, yoff, zoff, dx, dy, dz, d1, d)
				}
			}
		}
	}

	// Return the offsets found
	return xoff, yoff, zoff
}

// Measure the total Manhattan distance between two lists of points, with the points
// in the second list adjusted by the offsets
func totDist(s0 []P, s1 []P, xOffset, yOffset, zOffset int64) int64 {
	var totDist int64 = 0
	for _, p0 := range s0 {
		for _, p1 := range s1 {
			dist := abs(p0[0] - (p1[0] + xOffset))
			dist += abs(p0[1] - (p1[1] + yOffset))
			dist += abs(p0[2] - (p1[2] + zOffset))
			totDist += dist
		}
	}
	return totDist
}
