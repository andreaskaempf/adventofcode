// Advent of Code 2024, Day 21
//
// You are given a numeric keypad on which to type five 4-character
// codes. But you cannot type directly, but must mobilize a robot
// arm with a different keypad equipped with arrows and an Enter key.
// That arm must be controlled by another similar robot, which you can
// control using another keypad equipped with arrows. So two levels of
// indirection, 25 for Part 2. You must determine the length of each
// sequence that you type, in order to ultimately cause each code
// to be entered on the numeric keypad.
//
// This was very tricky, because upstream robots will have different
// path lengths depending on the paths of downstream robots. Ultimately,
// the solution for Part 2 involved recursively expanding paths, starting
// with the number pad and moving away from that. In order to make the
// problem tractable in terms of time and memory, the solution maintains
// a cache of sequence lengths at each level of recursion.
//
// This is possible because (A) the downstream keypads (after the number
// pad) always have to return the to A key at the end of each sequence,
// so you can treat sequences after the first as independent, and (B)
// you only need to know the final length, not the sequence itself. So
// the interim counters can be memoized, making the calculation very fast.
//
// AK, 22-27 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

// A pair of keys, e.g., start & end of path
type KeyPair struct {
	k1, k2 byte
}

// A sequence of key strokes
type KeySeq []byte

// List of shortest path between all keypairs on the numeric keypad, and
// on the direction keypad, precomputed
var numKeyPaths map[KeyPair][]KeySeq
var dirKeyPaths map[KeyPair][]KeySeq

// For caching in Part 2
type SeqLevel struct {
	level int
	seq   string // need to convert to string for key
}

var cache map[SeqLevel]int64

func main() {

	// Read the input file into a list of lines
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	codes := bytes.Split(data, []byte{'\n'})

	// Precompute all the shortest paths between keys on the numeric and direction keypads
	findAllShortest()

	// Do Part 1, either using first version with building up of
	// strings (slow and memory limited) or fast recursive addition
	// with counters for Part 2
	//ans1 := part1(codes)
	//fmt.Println("Part 1 =", ans1, "(using original logic)") // 126384/157892
	ans1a := part2(codes, 2) // alternative calculation
	fmt.Println("Part 1 =", ans1a, "(using Part 2 logic)")

	// Do Part 2
	ans2 := part2(codes, 25)      // 25 levels of recursion for Part 2, or 2 for Part 1
	fmt.Println("Part 2 =", ans2) // 197015606336332
}

//--------------------------------------------------------------------//
//                                 PART 1                             //
//--------------------------------------------------------------------//

// Part 1 solution, builds full length strings, not suitable for more
// than 2 levels of indirection. The new logic for Part 2 works for
// Part 1 as well, and is much faster (use 2 levels of indirection
// instead of 25).
func part1(codes [][]byte) int {

	// For each 4-character code, find all the shortest numeric
	// keypad paths for the entered code, then the shortest direction
	// keypads sequences for each of these.
	var ans int
	for _, code := range codes {

		// Robot 1: get all possible key sequences to directly input the code
		// on the numeric keypad
		fmt.Println("Processing code", string(code))
		seqs := NumPadEntry(code)
		fmt.Println("Robot 1:", len(seqs), "key patterns")

		// Repeat for number of robots, should be 2 for Part 1
		nrobots := 2
		for robot := 0; robot < nrobots; robot++ {
			fmt.Println("Robot", robot+2, ": expanding", len(seqs), "sequences from robot 1")
			seqs = expandSequences(seqs)
			fmt.Println(len(seqs), "sequences of length", len(seqs[0]))
		}

		// Convert to "complexity score" and add to answer:
		// numeric part of code * length of last sequence
		num, _ := strconv.Atoi(string(code[:3]))
		shortest := len(seqs[0])
		score := shortest * num
		fmt.Printf("Code seq length = %d, numeric = %d, score = %d\n\n", shortest, num, score)
		ans += score
	}

	return ans
}

// Get possible shortest path key sequences to directly input
// a code on the numeric key pad, uses shortest path lengths precomputed
// and stored in global dictionary
func NumPadEntry(code []byte) []KeySeq {

	// Get the possible paths for each transition on the number pad
	var prev byte = 'A' // key we're currently pointing at
	subseqs := [][]KeySeq{}
	for i := 0; i < len(code); i++ {
		this := code[i] // next character in code, 0-9 or A
		seqs := numKeyPaths[KeyPair{prev, this}]
		subseqs = append(subseqs, seqs)
		prev = this
	}

	// Combine all the subsequences into all possible combinations
	return combineAll(subseqs)
}

// Take a list of key sequences and expand each, creating a list of
// unique, shortest sequences
func expandSequences(seqs []KeySeq) []KeySeq {

	// Initialize variables to collect unique list of expansions
	haveSeq := map[string]bool{} // to check if exists
	newSeqs := []KeySeq{}        // list of new key sequences
	shortest := 9999999          // initialize to a big number

	// Process each (shorter) input sequence in list
	for _, s := range seqs {

		// Expand the sequence
		ss := expand(s)

		// Add to list, but if shorter than previous ones clear the list first,
		// so result is a list of the unique sequences with shortest length
		for _, s1 := range ss {
			if len(s1) < shortest {
				haveSeq = map[string]bool{string(s1): true}
				newSeqs = []KeySeq{s1}
				shortest = len(s1)
			} else if !haveSeq[string(s1)] {
				haveSeq[string(s1)] = true
				newSeqs = append(newSeqs, s1)
			}
		}
	}

	// Return just the list of sequences
	return newSeqs
}

// Expand one input key sequence by one level more remote, returning all
// possible sequences
func expand(input KeySeq) []KeySeq {

	// Get the expansions for each letter transition in the input sequence
	k0 := byte('A') // always start on the A key
	lists := [][]KeySeq{}
	for i := 0; i < len(input); i++ {
		k1 := input[i]
		kp := KeyPair{k0, k1}
		seqs := dirKeyPaths[kp]
		lists = append(lists, seqs)
		k0 = k1
	}

	return combineAll(lists)
}

// Given a list of lists, return a list of all possible combinations
// of the elements, concatenated
func combineAll(lists [][]KeySeq) []KeySeq {

	counters := make([]int, len(lists)) // all zeros to start
	res := []KeySeq{}
	done := false
	for !done { //counters[0] < len(lists[0]) {

		// Concatenate the current selection from each of the lists,
		// and add result to return value
		l := KeySeq{}
		for i := 0; i < len(lists); i++ {
			this := lists[i][counters[i]]
			l = append(l, this...)
		}
		//fmt.Println(string(l))
		res = append(res, l)

		// Increment counters starting with rightmost,
		// and "carry left"
		c := len(counters) - 1
		counters[c]++ // increment last column
		for c > 0 && counters[c] >= len(lists[c]) {
			counters[c] = 0
			counters[c-1]++ // carry left
			c--
		}
		if counters[0] >= len(lists[0]) {
			done = true
		}
	}
	return res
}

//--------------------------------------------------------------------//
//                                 PART 2                             //
//--------------------------------------------------------------------//

// Part 2 uses recursion with caching of counts per subsequence, to
// get around limitations with enumerating huge strings as used in Part 1
func part2(codes [][]byte, levels int) int64 {

	// Initialize cache
	cache = map[SeqLevel]int64{}

	// Do each code from the input
	var ans int64
	for _, code := range codes {

		// Get sequence of transitions on the number pad (e.g., 'A' -> '5').
		// For each transition, choose the shortest final sequence length,
		// by recursively running each through the designated number of
		// transformations (2 for Part 1, 25 for Part 2)
		var shortest int64
		for _, t := range transitions(code) {
			num_seq := expandNumeric(t, levels)
			shortest += expandDirectional(num_seq, levels)
		}

		// Multiply sequence length by numeric part of code to get score
		num, _ := strconv.ParseInt(string(code[:3]), 10, 64)
		score := shortest * num
		ans += score
	}

	return ans
}

// Calculate the optimal length for this key transition, to any level of
// indirection
func expandNumeric(kp KeyPair, levels int) KeySeq {

	// All possible shortest paths between these two keys (precomputed)
	paths := numKeyPaths[kp]

	// Pick the best one by cycling through levels of indirection (2 for
	// Part 1, 25 for Part 2)
	var best int64
	var bestSeq KeySeq
	for _, p := range paths {
		length := expandDirectional(p, levels)
		if best == 0 || length < best {
			best = length
			bestSeq = p
		}
	}
	return bestSeq
}

// Expand a sequence of two keys on the directional keypad,
// recursively, and return final length
func expandDirectional(seq KeySeq, level int) int64 {

	// Desired number of recursions reached, return final result
	if level == 0 {
		return int64(len(seq))
	}

	// Return memoized value if cached
	cacheKey := SeqLevel{level, string(seq)}
	cached, ok := cache[cacheKey]
	if ok {
		return cached
	}

	// Otherwise calculate this by recursively solving the optimal
	// path for each transition from here and adding up the lengths
	var length int64
	for _, t := range transitions(seq) {
		//dpt := directionPadTransitions[t]
		dpt := preferred(dirKeyPaths[t])
		length += expandDirectional(dpt, level-1)
	}

	// Save in cache
	cache[cacheKey] = length
	return length
}

// Return a list of all the consecutive transitions through a sequence, always
// starting at 'A'
func transitions(keys []byte) []KeyPair {
	trans := []KeyPair{}
	for i := 0; i < len(keys); i++ {
		k0 := byte('A') // inserted at beginning, since we start from A
		if i > 0 {
			k0 = keys[i-1]
		}
		k1 := keys[i]
		trans = append(trans, KeyPair{k0, k1})
	}
	return trans
}

// Apply overrides to direction key paths, to prefer left movement
// to vertical, and up to right (should really be done when these
// paths are enumerated)
func preferred(paths []KeySeq) KeySeq {
	if len(paths) == 1 {
		return paths[0]
	}
	p := string(paths[0])
	if p == "<v<A" {
		return KeySeq{'v', '<', '<', 'A'}
	} else if p == "^<A" {
		return KeySeq{'<', '^', 'A'}
	} else if p == "v<A" {
		return KeySeq{'<', 'v', 'A'}
	} else if p == ">^A" {
		return KeySeq{'^', '>', 'A'}
	} else {
		return paths[0]
	}
}

// Check if two sequences are the same, for debugging
func same(a, b []byte) bool {
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

//--------------------------------------------------------------------//
//                           PATH COMPUTATION                         //
//--------------------------------------------------------------------//

// An x,y point
type Point struct {
	x, y int
}

type Path []Point

// Working variable for building up list of key sequences
var _paths []KeySeq

// Get all shortest paths for each keypair, on both the numeric
// and direction keypads
func findAllShortest() {

	// Layout of numeric keypad, upside down so 0 on bottom
	numericKeypad := [][]byte{
		[]byte{'#', '0', 'A'}, // row 0
		[]byte{'1', '2', '3'}, // row 1
		[]byte{'4', '5', '6'}, // row 2
		[]byte{'7', '8', '9'}, // row 3
	}

	// Find all shortest paths, for every key pair
	numKeyPaths = findGridShortest(numericKeypad)

	// Layout of direction keypad, upside down so 0 on bottom
	directionKeypad := [][]byte{
		[]byte{'<', 'v', '>'}, // row 0
		[]byte{'#', '^', 'A'}, // row 1
	}

	// Find all shortest paths, for every key pair
	dirKeyPaths = findGridShortest(directionKeypad)
}

// Get all shortest paths between each pair of keys on a grid
// (keypad), avoiding # marks. Return map of
// fromkey,tokey -> list of shortest paths, defined as a list of point (coordinates)
func findGridShortest(grid [][]byte) map[KeyPair][]KeySeq {

	// Enumerate all points on the grid (buttons on keypad), excluding #
	points := []Point{}
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			p := Point{x: c, y: r}
			if at(p, grid) != '#' {
				points = append(points, p)
			}
		}
	}

	// Recursively trace all paths between each pair of keys, keep only the
	// shortest paths, and convert them from coordinates to arrow key chars
	keyPairPaths := map[KeyPair][]KeySeq{}
	for _, p1 := range points {
		for _, p2 := range points {
			kp := KeyPair{at(p1, grid), at(p2, grid)}
			_paths = []KeySeq{}            // initalize global variable _paths
			trace(p1, p2, grid, []Point{}) // get all paths, into global _paths
			_paths = shortestOnly(_paths)  // just the shortest
			keyPairPaths[kp] = _paths
		}
	}

	// Return dictionary of fromkey,tokey -> list of shortest paths
	return keyPairPaths
}

// Filter list of paths to only those with the shortest length
func shortestOnly(pathlist []KeySeq) []KeySeq {

	// Get length of shortest path
	var shortest = 0
	for _, p := range pathlist {
		if shortest == 0 || len(p) < shortest {
			shortest = len(p)
		}
	}

	// Collect just the paths of that length
	shortestPaths := []KeySeq{}
	for _, p := range pathlist {
		if len(p) == shortest {
			shortestPaths = append(shortestPaths, p)
		}
	}
	return shortestPaths
}

// Recursively search depth-first from given x,y to end character,
// add each key sequence found to global _paths variable
func trace(p, end Point, grid [][]byte, visited Path) {

	// Save path if at end
	if p == end {
		visited = append(visited, p)
		keyseq := coordPathToArrows(visited)
		keyseq = append(keyseq, 'A')
		_paths = append(_paths, keyseq)
		return
	}

	// Otherwise try in each unvisited direction
	visited = append(visited, p)
	for _, a := range adjacents(p) {
		if at(a, grid) == '#' || in(a, visited) {
			continue
		}
		vis := make([]Point, len(visited), len(visited))
		copy(vis, visited)
		trace(a, end, grid, vis)
	}
}

// Get adjacent points
func adjacents(p Point) []Point {
	return []Point{Point{p.x, p.y - 1}, Point{p.x + 1, p.y},
		Point{p.x, p.y + 1}, Point{p.x - 1, p.y}}
}

// Convert a path of coordinates to a list of arrow symbols
func coordPathToArrows(path Path) KeySeq {
	keys := KeySeq{}
	for i := 1; i < len(path); i++ {
		k0 := path[i-1]
		k1 := path[i]
		dx := k1.x - k0.x
		dy := k1.y - k0.y
		if dx == 1 { // move right
			keys = append(keys, '>')
		} else if dx == -1 { // move left
			keys = append(keys, '<')
		} else if dy == 1 { // move down
			keys = append(keys, '^')
		} else if dy == -1 { // move up
			keys = append(keys, 'v')
		}
	}
	return keys
}

// What is at a point on the grid
func at(p Point, grid [][]byte) byte {
	if p.y < 0 || p.y >= len(grid) || p.x < 0 || p.x >= len(grid[p.y]) {
		return '#' // out of bounds
	}
	return grid[p.y][p.x]
}

// Is point in a list?
func in(p Point, pp []Point) bool {
	for _, x := range pp {
		if x == p {
			return true
		}
	}
	return false
}

// Absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
