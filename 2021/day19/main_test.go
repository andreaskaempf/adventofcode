// Advent of Code 2021, Day 19
//
// A few unit tests
//
// AK, 18/12/2021

package main

import (
	//"fmt"
	"testing"
)

// Test matrix multiplication
func TestMatrix(t *testing.T) {

	// Test example from math site
	A := []P{P{1, -1, 2}, P{0, -3, 1}}
	x := P{2, 1, 0}
	res := matMult(A, x)
	sb := P{1, -3}
	if !vectorsSame(res, sb) {
		t.Error("Unable to multiply vectors (1)")
	}

	// Test example from cube rotation
	// | 0 -1  0 |   |1|   |0|
	// | 1  0  0 | * |0| = |1|
	// | 0  0  1 |   |0|   |0|
	A = []P{P{0, -1, 0}, P{1, 0, 0}, P{0, 0, 1}}
	x = P{1, 0, 0}
	res = matMult(A, x)
	sb = P{0, 1, 0}
	if !vectorsSame(res, sb) {
		t.Error("Unable to multiply vectors (2)")
	}

}

func TestRotate(t *testing.T) {

	// Unrotated cube
	rot0 := []P{P{-1, -1, 1}, P{-2, -2, 2}, P{-3, -3, 3}, P{-2, -3, 1}, P{5, 6, -4}, P{8, 0, 7}}

	// Various rotations provided as examples
	rot1 := []P{P{1, -1, 1}, P{2, -2, 2}, P{3, -3, 3}, P{2, -1, 3}, P{-5, 4, -6}, P{-8, -7, 0}}
	rot2 := []P{P{-1, -1, -1}, P{-2, -2, -2}, P{-3, -3, -3}, P{-1, -3, -2}, P{4, 6, 5}, P{-7, 0, 8}}
	rot3 := []P{P{1, 1, -1}, P{2, 2, -2}, P{3, 3, -3}, P{1, 3, -2}, P{-4, -6, 5}, P{7, 0, 8}}
	rot4 := []P{P{1, 1, 1}, P{2, 2, 2}, P{3, 3, 3}, P{3, 1, 2}, P{-6, -4, -5}, P{0, 7, -8}}

	// Get all possible rotations for the first cube (should be 24)
	rots := rotateAllWays(rot0)
	if len(rots) != 24 {
		t.Error("Rotation did not create 24 cubes")
	}

	// Make sure each of the four is in the set
	var found1, found2, found3, found4 bool
	for _, rot := range rots {
		if matSame(rot, rot1) {
			found1 = true
		} else if matSame(rot, rot2) {
			found2 = true
		} else if matSame(rot, rot3) {
			found3 = true
		} else if matSame(rot, rot4) {
			found4 = true
		}
	}
	if !(found1 && found2 && found3 && found4) {
		t.Error("Rotation did not find all four examples")
	}
}
