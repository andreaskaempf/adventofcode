// Transcription of problem input into Go, and reduction of this
// to iterated version. These were intermediate versions of the
// calculation, see day24.go for the final version, which is much
// simplified, and eliminates the need to keep track of x and y
// (only z is maintained between iterations).
//
// So this file is not required for the solution, just kept as
// back-up.

package main

import "fmt"

// Conversion of the full problem input into Go, using search and replace
func monad1(inpString string) map[string]int {

	// The registers
	var w, x, y, z int

	// Convert the input into an array of digits
	input := parseInput(inpString)
	if len(input) == 0 {
		return map[string]int{}
	}

	// Direct transcription of the problem into Go
	w = input[0]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 12
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 4
	y = y * x
	z = z + y

	w = input[1]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 11
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y

	w = input[2]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 13
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 5
	y = y * x
	z = z + y

	w = input[3]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 11
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y

	w = input[4]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 14
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 14
	y = y * x
	z = z + y

	w = input[5]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -10
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 7
	y = y * x
	z = z + y

	w = input[6]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 11
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y

	w = input[7]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -9
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 4
	y = y * x
	z = z + y

	w = input[8]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -3
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 6
	y = y * x
	z = z + y

	w = input[9]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 1)
	x = x + 13
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 5
	y = y * x
	z = z + y

	w = input[10]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -5
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 9
	y = y * x
	z = z + y

	w = input[11]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -10
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 12
	y = y * x
	z = z + y

	w = input[12]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -4
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 14
	y = y * x
	z = z + y

	w = input[13]
	x = x * 0
	x = x + z
	x = x % 26
	z = div(z, 26)
	x = x + -5
	x = eql(x, w)
	x = eql(x, 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 14
	y = y * x
	z = z + y

	// Input is valid if z is zero at the end
	return map[string]int{"w": w, "x": x, "y": y, "z": z} //z == 0
}

// Implementation of the div and eql functions
func div(a, b int) int {
	return a / b
}

func eql(a, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}

// Iterative implementation of the original algorithm, basically breaks it into
// a loop, with lists of numbers for the bits that change.
func monad2(inpString string) map[string]int {

	// The registers (all start at zero)
	var w, x, y, z int

	// Convert the input into an array of digits
	input := parseInput(inpString)
	if len(input) == 0 {
		return map[string]int{}
	}

	// Constants from the each block of the problem
	A := [14]int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}          // z = div(z, A)
	B := [14]int{12, 11, 13, 11, 14, -10, 11, -9, -3, 13, -5, -10 - 4, -5} // x = x + B[i]
	C := [14]int{4, 11, 5, 11, 14, 7, 11, 4, 6, 5, 9, 12, 14, 14}

	// Direct transcription of the problem into Go
	for i := 0; i < 14; i++ {
		w = input[i]
		x = x * 0
		x = x + z
		x = x % 26
		z = div(z, A[i]) // 1
		x = x + B[i]     // 12
		x = eql(x, w)
		x = eql(x, 0)
		y = y * 0
		y = y + 25
		y = y * x
		y = y + 1
		z = z * y
		y = y * 0
		y = y + w
		y = y + C[i] // 4
		y = y * x
		z = z + y
	}

	// Input is valid if z is zero at the end
	return map[string]int{"w": w, "x": x, "y": y, "z": z} //z == 0
}

// Parse string of 14 digits into an array of integers,
// return empty slice if not valid (i.e., not 14 long,
// not all digits, or contains zeros)
func parseInput(input string) []int {

	// Convert the input into an array of digits
	result := []int{}
	for i := 0; i < len(input); i++ {
		n := int(input[i] - '0')
		if n < 1 || n > 9 {
			fmt.Println("Invalid digit:", input)
			return []int{}
		}
		result = append(result, n)
	}

	// Check length (can't do this on string because of runes)
	if len(result) != 14 {
		fmt.Println("Invalid length:", input)
		return []int{}
	}

	// Okay
	return result
}
