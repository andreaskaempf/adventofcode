// matrices.go
package main

//import "fmt"

// Compare two arrays of int arrays element-by-element, and report
// if they are the same
func matSame(a []P, b []P) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !vectorsSame(a[i], b[i]) {
			return false
		}
	}
	return true
}

// Compare two int arrays element-by-element, and report
// if they are the same
func vectorsSame(a []int64, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func containsVector(vlist []P, v P) bool {
	for _, p := range vlist {
		if vectorsSame(p, v) {
			return true
		}
	}
	return false
}

func copyMatrix(m []P) []P {
	res := []P{}
	for _, r := range m {
		res = append(res, copyVector(r))
	}
	return res
}

func copyVector(v P) P {
	res := make(P, len(v), len(v))
	for i := 0; i < len(v); i++ {
		res[i] = v[i]
	}
	return res
}

// Matrix multiply a matrix by a point
// https://mathinsight.org/matrix_vector_multiplication
func matMult(m []P, p P) P {
	res := P{}
	for _, r := range m {
		res = append(res, sumProd(r, p))
	}
	return res
}

// Do sum-product of two vectors, i.e., sum of product of each pair
func sumProd(v1 P, v2 P) int64 {
	if len(v1) != len(v2) {
		panic("sumProd: vectors not equal length")
	}
	var res int64
	for i := 0; i < len(v1); i++ {
		res += v1[i] * v2[i]
	}
	return res
}

// Rotating a cube n times along each axis, by multiplying each point
// by a matrix corresponding to 90 degrees of X, Y, Z rotation any
// number of times
// https://stackoverflow.com/questions/34050929/3d-point-rotation-algorithm
func rotate(s []P, x, y, z int) []P {

	// The matrices by which to multiply for each rotation type,
	// noting that sin(90) = 1, cos(90) = 0
	const sin90 = 1
	const cos90 = 0
	X := []P{P{1, 0, 0}, P{0, cos90, -sin90}, P{0, sin90, cos90}}
	Y := []P{P{cos90, 0, sin90}, P{0, 1, 0}, P{-sin90, 0, cos90}}
	Z := []P{P{cos90, -sin90, 0}, P{sin90, cos90, 0}, P{0, 0, 1}}

	// Rotate each point along each axis
	res := []P{}
	for ri := 0; ri < len(s); ri++ {
		r := copyVector(s[ri])
		for i := 0; i < x; i++ {
			r = matMult(X, r)
		}
		for i := 0; i < y; i++ {
			r = matMult(Y, r)
		}
		for i := 0; i < z; i++ {
			r = matMult(Z, r)
		}
		res = append(res, r)
	}

	return res
}

// Rotate cube in all directions, returning list of 24 cubes
func rotateAllWays(s0 []P) [][]P {
	found := [][]P{}
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				c := rotate(s0, x, y, z)
				exists := false
				for _, f := range found {
					if matSame(f, c) {
						//fmt.Println("Already found!")
						exists = true
						break
					}
				}
				if !exists {
					found = append(found, c)
				}
			}
		}
	}
	return found
}
