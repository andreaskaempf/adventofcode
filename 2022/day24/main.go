// Advent of Code 2022, Day 24
//
// Find shortest path from entry to exit of a rectangular field, avoiding
// "blizzards" that move every time step. For Part 2, also move back to
// entry then back to exit, and add up all the steps.
//
// Used dynamic programming, depth-first search with memoization of
// previously found best values for each position+time combination
// (since the map changes each time step, you can't just use the position).
//
// AK, 24 Dec 2022

package main

import (
	"fmt"
)

// A terrain is the map at one point in time, with all its blizzards
type Terrain struct {
	blizzards []Blizzard // can't use a map, because may be more than one
}

// We will pre-compute all the terrains for each day
var terrains []Terrain

// Info about a blizzard
type Blizzard struct {
	p         Point // the location
	direction byte  // < > ^ or v
}

// A point in 2D space
type Point struct {
	x, y int
}

// Min/max coordinates (starting from 1)
var minX, maxX, minY, maxY int

// Locations of the doors
var entry, exit Point

// Best optimization result found
var bestOverall int

// For keeping track of points we have visited
type PointTime struct {
	p    Point
	time int
}

// History of points we have visited
var visited map[PointTime]int

func main() {

	// Read the input file into a map of positions of the blizzard,
	// which becomes the first terrain
	fname := "sample.txt"
	fname = "sample2.txt"
	//fname = "input.txt" // uncomment to run on real input
	lines := readLines(fname)
	t := Terrain{}
	for y := 1; y < len(lines)-1; y++ {
		l := lines[y]
		for x := 1; x < len(l)-1; x++ {
			if l[x] != '.' {
				b := Blizzard{Point{x, y}, l[x]}
				t.blizzards = append(t.blizzards, b)
			}
		}
	}
	terrains = append(terrains, t)

	// Set the min/max x and y, and locations of entry and exit
	minX = 1
	minY = 1
	maxX = len(lines[0]) - 2
	maxY = len(lines) - 2
	entry = Point{1, 0}
	exit = Point{maxX, maxY + 1}

	// Precompute the terrain at each step of the simulation
	nterrains := (maxX + maxY) * 4 // this will be the maximum number of time steps
	nterrains *= 3                 // for Part 2
	fmt.Printf("Precomputing %d terrains\n", nterrains)
	for i := 1; i < nterrains; i++ { // increase as necessary

		// Make a new terrain and simulate movement of each blizzard
		t0 := terrains[len(terrains)-1] // start with the last terrain
		t1 := Terrain{}                 // initialize a new one
		for _, b := range t0.blizzards {
			p1 := Point{b.p.x, b.p.y}
			if b.direction == '>' {
				p1.x = ifElse(p1.x == maxX, 1, p1.x+1)
			} else if b.direction == '<' {
				p1.x = ifElse(p1.x == 1, maxX, p1.x-1)
			} else if b.direction == 'v' {
				p1.y = ifElse(p1.y == maxY, 1, p1.y+1)
			} else if b.direction == '^' {
				p1.y = ifElse(p1.y == 1, maxY, p1.y-1)
			} else {
				panic("Invalid direction!")
			}
			b := Blizzard{p1, b.direction}
			t1.blizzards = append(t1.blizzards, b)

		}

		// Add blizzard to the new terrain
		terrains = append(terrains, t1)
	}

	// Part 1: minimal number of minutes to get from entry to exit
	// 176 too low, 258 too low, 373 right
	fmt.Println("Part 1: Optimizing")
	ans1 := optimize(entry, exit, []Point{entry}, 0)
	fmt.Println("Part 1 (s/b 18 for sample 2, 373 for input) =", ans1)

	// Parts 2: add the minimal amounts of time to go back to the
	// entrance, then back to the exit.
	// Important: don't start back at terrain 0, but continue from
	// where left off.
	// For sample 2, should be 18 + 23 + 13 = 54 minutes

	fmt.Println("Part 2: Optimizing back to entry")
	bestOverall = -1              // need to reinitialize this
	visited = map[PointTime]int{} // need to reinitialize this
	ans2 := optimize(exit, entry, []Point{exit}, ans1-1)
	fmt.Printf("  took %d minutes\n", ans2)

	fmt.Println("Part 2: Optimizing back to exit")
	bestOverall = -1              // need to reinitialize this
	visited = map[PointTime]int{} // need to reinitialize this
	ans3 := optimize(entry, exit, []Point{exit}, ans2-1)
	fmt.Printf("  took %d minutes\n", ans3)
	fmt.Println("Part 2 (s/b 54) =", ans3)

}

// For the optimization, do a depth-first recursive search, subject to movement
// of the blizzards at each step, and return the best possible time to the
// destination
func optimize(here, dest Point, path []Point, t int) int {

	// Initialize first iteration (t = 0), needs to be "manually"
	// if you want to re-run optimization from last starting point,
	// i.e., not from zero
	if t == 0 {
		bestOverall = -1
		visited = map[PointTime]int{}
	}

	// If you have successfully reached the destination, return
	// the time it took, report of best result so far, and optionally
	// show the path taken
	if here == dest {
		if bestOverall == -1 || t < bestOverall {
			fmt.Println("New best time", t)
			//fmt.Println("  path:",  path)  // uncomment to show path
			bestOverall = t
		}
		return t
	}

	// No more terrains left, dead end (or increase the number of
	// precomputed terrains)
	if t >= len(terrains)-1 {
		return 0
	}

	// If already visited this point at this time, return result from previous
	// visit. This "memoization" *hugely* speeds up the optimization, from many
	// minutes to under a hundredth of a second for the sample.
	vis, ok := visited[PointTime{here, t}]
	if ok {
		return vis
	}

	// Find out where can we go from here, that will not be blocked in
	// the next time step
	t1 := terrains[t+1]            // look at next period's terrain
	candidates := []Point{}        // list of candidates
	if empty(here.x, here.y, t1) { // consider staying put, unless blizzard
		candidates = append(candidates, here)
	}
	if here.y >= minY && here.y <= maxY && here.x > minX && empty(here.x-1, here.y, t1) { // left
		candidates = append(candidates, Point{here.x - 1, here.y})
	}
	if here.y >= minY && here.y <= maxY && here.x < maxX && empty(here.x+1, here.y, t1) { // right
		candidates = append(candidates, Point{here.x + 1, here.y})
	}
	if here.y > minY && empty(here.x, here.y-1, t1) { // up
		candidates = append(candidates, Point{here.x, here.y - 1})
	}
	if here.y < maxY && empty(here.x, here.y+1, t1) { // down, non exit
		candidates = append(candidates, Point{here.x, here.y + 1})
	}
	if dest == exit && here.x == exit.x && here.y == maxY { // down, to exit
		candidates = append(candidates, Point{here.x, here.y + 1})
	}
	if dest == entry && here.x == entry.x && here.y == minY { // up, to entry
		candidates = append(candidates, Point{here.x, here.y - 1})
	}

	// Recursively try out each possible move from this place/time
	best := 0 // best this time step
	for _, c := range candidates {
		path1 := path
		//path1 := copyPath(path) // uncomment these two lines to show each
		//path1 = append(path1, here) // path found (for debugging)
		o := optimize(c, dest, path1, t+1)    // optimize from here/now
		if best == 0 || (o > 0 && o < best) { // is this destination better?
			best = o
		}
	}

	// Record the current best from here/now, and return best result
	visited[PointTime{here, t}] = best
	return best
}

// Is a location empty?
func empty(x, y int, t Terrain) bool {
	for _, b := range t.blizzards {
		if b.p.x == x && b.p.y == y {
			return false
		}
	}
	return true
}

// Everything below his is just for debugging, can be removed

// Make a copy of a list of paths, for keeping track of journey (for debugging)
func copyPath(path []Point) []Point {
	path1 := []Point{}
	for _, p := range path {
		path1 = append(path1, p)
	}
	return path1
}

// Draw map (for debugging)
func draw(t Terrain) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			conts := contents(x, y, t)
			if len(conts) == 0 {
				fmt.Print(".")
			} else if len(conts) == 1 {
				fmt.Print(string(conts[0]))
			} else {
				fmt.Printf("%d", len(conts))
			}
		}
		fmt.Println()
	}
}

// Return the contents(s) of given location on a terrain
// Only used for drawing
func contents(x, y int, t Terrain) []byte {
	res := []byte{}
	for _, b := range t.blizzards {
		if b.p.x == x && b.p.y == y {
			res = append(res, b.direction)
		}
	}
	return res
}
