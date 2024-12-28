// Advent of Code 2024, Day 24
//
// You are given a set of 1/0 values for a bunch of x/y variables, each
// representing a bit.  For Part 1, evaluate all variables starting with z,
// recursively performing AND/OR/XOR operations in a hierarchy of equations,
// and finally assemble the results into a binary number, and convert binary to
// decimal.
//
// For Part 2, ignore the initial given values for the x/y values, and instead
// find 4 equation swaps that make the system work as an adder that outputs sum
// of any x and y bits into z. The x, y, and z variables represent bits, in
// reverse order from the usual (i.e., least significant bit first).
//
// Part 1 was a straightforward recursive equation interpreter. For Part 2, I
// used heuristically guided brute force, by testing all z circuits for those
// that would fail to add properly, gathering all the input variables for these
// circuits, and testing every pair of those for swaps that would reduce the
// number of errors output by the system.  Of these 20 pairs that reduced the
// error count, then tried all combinations of 4 swaps for one that would make
// the adder error-free.
//
// AK, 24 & 28 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

// A variable, with either a value (constants, e.g., x and y vars), or
// an operation with left & right sides, and operator XOR/OR/AND
type Var struct {
	name     string
	lhs, rhs string
	op       string
	value    int
}

// Global list of variables
var vars []Var

func main() {

	// Read the input file into initial register values, followed by operations
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	vars = []Var{} // global
	for _, l := range strings.Split(string(data), "\n") {
		parts := strings.Split(l, " ")
		var v Var
		if len(parts) == 2 { // x03: 1
			vname := parts[0][:3]
			v = Var{name: vname, value: atoi(parts[1])}
			vars = append(vars, v)
		} else if len(parts) == 5 { // x00 XOR y04 -> ntg
			vname := parts[4]
			v = Var{name: vname, lhs: parts[0], rhs: parts[2], op: parts[1], value: -1}
			vars = append(vars, v)
		}
	}

	// Part 1: evaluate values of all variables starting with z, assemble
	// them into a binary number, and convert binary to decimal
	part1()

	// Part 2: find 4 swaps that make the system work as an adder that outputs
	// sum of x and y into z. The x, y, and z variables represent bits, in
	// reverse order from the usual (i.e., least significant bit first).
	// com in separate function
	part2()
}

//------------------------------------------------------------------//
//                             PART 1                               //
//------------------------------------------------------------------//

// Part 1: evaluate the value of any variables starting with z,
// assemble them into a binary number,convert binary to decimal
func part1() {

	// Get list of variable names starting with z, sort by name
	zvars := getVarNames('z')
	sort.Strings(zvars)

	// Evaluate each z variable, i.e., each bit, and build up into
	// a binary number (in reverse order!)
	bits := []int{}
	for _, z := range zvars {
		bits = append(bits, evaluate(z, 0))
	}

	// Show result in decimal
	fmt.Println("Part 1 =", binToDec(bits)) // 2024, 53325321422566
}

// Evaluate value of a variable, recursively applying operations on inputs
func evaluate(vname string, depth int) int {

	// Catch endless loop, happens during testing in Part 2
	if depth > 100 {
		return -1
	}

	// For Part 2, change variable name if it being swapped
	vs, ok := swaps[vname]
	if ok {
		vname = vs
	}

	// Constant or already evaluated: return value
	v := getVar(vname) // returns a copy, do not change
	if v.value >= 0 {
		return v.value
	}

	// Otherwise calculate
	lhs := evaluate(v.lhs, depth+1)
	rhs := evaluate(v.rhs, depth+1)
	var y int
	if v.op == "AND" {
		y = AND(lhs, rhs)
	} else if v.op == "OR" {
		y = OR(lhs, rhs)
	} else if v.op == "XOR" {
		y = XOR(lhs, rhs)
	} else {
		panic("Bad op")
	}
	//setVar(v.name, y) // don't do this if you use value to determine if constant
	return y
}

// Interpret a binary number (list of bits, least significant first!), and
// convert to decimal
func binToDec(bits []int) int {
	factor := 1
	var n int
	for i := 0; i < len(bits); i++ { // bits are in reverse order!
		n += bits[i] * factor
		factor *= 2
	}
	return n
}

// Convert a decimal number to binary, return bits in reverse order (least
// significant bit first), only used in Part 2
func decToBin(n int) []int {
	s := fmt.Sprintf("%b", n) // string with binary representation
	bits := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		bits = append(bits, int(s[i]-'0'))
	}
	return bits
}

// Get variable names starting with letter
func getVarNames(prefix byte) []string {
	res := []string{}
	for _, v := range vars {
		if v.name[0] == prefix {
			res = append(res, v.name)
		}
	}
	return res
}

// Find variable by name (would be faster with map)
func getVar(vname string) Var {
	for i := 0; i < len(vars); i++ {
		if vars[i].name == vname {
			return vars[i]
		}
	}
	fmt.Println("Not found:", vname)
	return Var{} // if not found, return empty variable
}

// Find variable by name (would be faster with map)
func setVar(vname string, val int) {
	for i := 0; i < len(vars); i++ {
		if vars[i].name == vname {
			vars[i].value = val
			return
		}
	}
	fmt.Println("setVar: missing", vname)
}

// AND gates output 1 if both inputs are 1; if either input is 0, these gates output 0
func AND(a, b int) int {
	if a == 1 && b == 1 {
		return 1
	} else {
		return 0
	}
}

// OR gates output 1 if one or both inputs is 1; if both inputs are 0, these gates output 0
func OR(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	} else {
		return 0
	}
}

// XOR gates output 1 if the inputs are different; if the inputs are the same, these gates output 0
func XOR(a, b int) int {
	if a == b {
		return 0
	} else {
		return 1
	}
}

//------------------------------------------------------------------//
//                             PART 2                               //
//------------------------------------------------------------------//

// Used for collecting successful swaps
type Pair struct {
	a, b string
}

// Global variables used by Part 2
var swaps map[string]string // dictionary of swapped variables
var inputs []string         // used by getInputs() to get all equation inputs
var faulty []string         // used by testAdder

// Part 2: Determine which four pairs of gates need their outputs swapped
// so that your system correctly performs addition, i.e., x bits + y bits = z bits
func part2() {

	// Find which z variables incorrectly add 1+1 or 1+0
	// Bits that fail: z15, z16, z21, z30, z31, z36, z37
	nerrs := testAdder()
	fmt.Println(nerrs, "lines are faulty:", faulty)

	// Find all variables used by these z variables, as these are candidates
	// for swapping (includes z variables, but excludes x/y constants)
	candidates := []string{"rrn"}
	for _, f := range faulty {
		inputs = []string{} // clear out global list
		getInputs(f)
		candidates = append(candidates, inputs...)
	}

	// Get unique values and sort for informative output
	candidates = unique(candidates)
	sort.Strings(candidates) // to make progress readable
	fmt.Println(len(candidates), "candidates for swapping")
	fmt.Println(candidates)

	// Find all pairs of these variables, that if swapped reduce the number of
	// errors, using brute force to check each pair
	nfaulty := testAdder()              // initial number of faulty wires
	ncands := len(candidates)           // number of candidates
	npairs := ncands * (ncands - 1) / 2 // number of pairs
	successful := []Pair{}              // swaps that reduce errors
	n := 0                              // for showing progress
	fmt.Println("Finding swap pairs that reduce errors, at", time.Now().String())
	for i := 0; i < ncands; i++ {
		for j := i + 1; j < ncands; j++ {

			// Do the swap
			c1 := candidates[i]
			c2 := candidates[j]
			clearSwaps()
			swap(c1, c2)

			// Show progress
			n++
			pcnt := float64(n) / float64(npairs) * 100.0
			fmt.Printf("Swapping %s %s (%d / %d = %.2f%%), %d pairs found: ",
				c1, c2, n, npairs, pcnt, len(successful))
			fmt.Println(successful)

			// Test and record swap if reduced errors
			nfaulty2 := testAdder()
			if nfaulty2 < nfaulty {
				fmt.Printf("*** %s <-> %s reduced errors from %d to %d\n",
					c1, c2, nfaulty, nfaulty2)
				successful = append(successful, Pair{c1, c2})
			}
		}
	}
	fmt.Println(len(successful), "error-reducing pairs found at", time.Now().String())

	// We end up with 20 pairs, so try again to find combination of 4 pairs
	// that eliminates all errors. First, assemble all combinations of 4 swaps
	// where there is no overlap, i.e., one element is present in another pair
	quartets := [][]Pair{}
	for i := 0; i < len(successful); i++ {
		for j := i + 1; j < len(successful); j++ {
			for k := j + 1; k < len(successful); k++ {
				for l := k + 1; l < len(successful); l++ {
					p1 := successful[i]
					p2 := successful[j]
					p3 := successful[k]
					p4 := successful[l]
					if !overlaps(p1, p2, p3, p4) {
						q := []Pair{p1, p2, p3, p4}
						quartets = append(quartets, q)
					}
				}
			}
		}
	}
	fmt.Println(len(quartets), "non-overlapping combinations of 4 swaps to test")

	// Test each combination, find the one that makes the adder work perfectly
	fmt.Println("Testing", len(quartets), "at", time.Now().String())
	bestErrors := 999999
	var bestPairs []Pair
	for i, q := range quartets {

		// Do the swaps and test
		clearSwaps()
		for _, p := range q {
			swap(p.a, p.b)
		}
		nfaulty2 := testAdder()
		fmt.Println(i, "/", len(quartets), ":", q, "=>", nfaulty2, ", best =", bestErrors)

		// Update if improved result
		if nfaulty2 < bestErrors {
			bestErrors = nfaulty2
			bestPairs = q
			fmt.Println("Best found:", bestPairs, bestErrors)
		}
	}

	// If reached zero errors, we have a solution, turn into a list of strings,
	// sort them, and output separated by commas
	fmt.Println("Best set of four Swaps =", bestPairs, "with", bestErrors, "errors")
	fmt.Println("Finished at", time.Now().String())
	if bestErrors == 0 {

		// Turn pairs into a list
		ans := []string{}
		for _, n := range bestPairs {
			ans = append(ans, n.a, n.b)
		}

		// Output answer as comma-separated sorted list
		sort.Strings(ans)
		fmt.Println("Part 2 =", strings.Join(ans, ",")) // fkb,nnr,rdn,rqf,rrn,z16,z31,z37
	} else {
		fmt.Println("No solution found")
	}
}

// Determine if the adder works, by doing a couple of checks at each bit
// level, returning number of faults and setting global variable 'faulty'
// to list of faulty wires
func testAdder() int {
	x := 1
	faulty = []string{} // global
	for b := 0; b < 45; b++ {
		if !testAddition(x, x) || !testAddition(x, 1) {
			faulty = append(faulty, fmt.Sprintf("z%02d", b))
		}
		x *= 2
	}
	return len(faulty)
}

// Test addition using the logic gates: x bits + y bits = z bits
func testAddition(x, y int) bool {

	// Set all x and y registers to zero
	for i := 0; i < len(vars); i++ {
		if vars[i].name[0] == 'x' || vars[i].name[0] == 'y' {
			vars[i].value = 0
		}
	}

	// Convert each number to binary and set bits
	xbits := decToBin(x)
	for i := 0; i < len(xbits); i++ {
		vname := fmt.Sprintf("x%02d", i)
		setVar(vname, xbits[i])
	}

	ybits := decToBin(y)
	for i := 0; i < len(ybits); i++ {
		vname := fmt.Sprintf("y%02d", i)
		setVar(vname, ybits[i])
	}

	// Perform addition, by running simulation, extracting bits from z variables,
	// and turning back into a decimal number
	zvars := getVarNames('z')
	sort.Strings(zvars)
	bits := []int{}
	for _, v := range zvars {
		eval := evaluate(v, 0)
		if eval == -1 { // caught endless loop
			fmt.Println("Endless loop trying to evaluate", v)
			return false
		}
		bits = append(bits, eval)
	}
	ans := binToDec(bits)

	// Evaluate answer
	correct := ans == x+y
	return correct
}

// Get inputs to an equation
func getInputs(vname string) {

	// If an x/y variable, we have reached end of branch
	isXY := vname[0] == 'x' || vname[0] == 'y'
	if isXY {
		return
	}

	// Add this variable to list of inputs
	inputs = append(inputs, vname)

	// Trace both input variables
	v := getVar(vname)
	getInputs(v.lhs)
	getInputs(v.rhs)
}

// Check pairs for overlaps, i.e., one pair shares a point from another
func overlaps(p1, p2, p3, p4 Pair) bool {

	// Count the number of times each label occurs
	counts := map[string]int{}
	counts[p1.a]++
	counts[p1.b]++
	counts[p2.a]++
	counts[p2.b]++
	counts[p3.a]++
	counts[p3.b]++
	counts[p4.a]++
	counts[p4.b]++

	// Return true if any counts > 1
	for _, v := range counts {
		if v > 1 {
			return true
		}
	}
	return false
}

// Unique strings from list
func unique(l []string) []string {
	l2 := []string{}
	for _, s := range l {
		if !in(s, l2) {
			l2 = append(l2, s)
		}
	}
	return l2
}

// Swap two variables
func swap(v1, v2 string) {
	swaps[v1] = v2
	swaps[v2] = v1
}

// Clear all swaps
func clearSwaps() {
	swaps = map[string]string{}
}

// Recursively trace equations for a variable, and output complete abstract
// syntax tree, used for debugging and inspection of logic structure
func trace(v Var) string {

	// Return x & y variables by name
	if v.name[0] == 'x' || v.name[0] == 'y' {
		return fmt.Sprintf(v.name)
	}

	// Return operation
	lhs := getVar(v.lhs)
	rhs := getVar(v.rhs)
	return fmt.Sprintf("(%s %s %s)", trace(lhs), v.op, trace(rhs))

}

// Run the trace on all z variables
func traceAll() {
	for i := 0; i < 45; i++ {
		vname := fmt.Sprintf("z%02d", i)
		v := getVar(vname)
		fmt.Print("\n\n*** ", vname, " = ", trace(v))
	}
}

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}

// Is element in a list?
func in[T int | float64 | byte | string](c T, s []T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return true
		}
	}
	return false
}
