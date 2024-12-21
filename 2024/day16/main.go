// Advent of Code 2024, Day 16
//
// Given a maze, find the shortest path, with the twist that 90 degree turns
// have a cost of 1000. For Part 2, count up all the squares that were
// traversed while reaching the end at the lowest cost. Did this using Djikstra
// algorithm, adapted to deal with the added cost of turns, and keeping track
// of squares traversed.
//
// AK, 16-21 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

// A 2D point, with optional direction (0=up, 1=right,..)
type Point struct {
	x, y, dir int
}

// Rows of input data
var rows [][]byte

// Start and end points, and optimal distance found in Part 1
var start, end Point
var shortestPath int

// Distance matrix, global so we can use it for Part 2
var dist map[Point]int // lowest cost to each position

// For Part 2, a map of every point visited on any shortest path
var marked map[Point]int

// List of sources for each point, i.e., the points that entered this point
var sources map[Point][]Point // not used
var destinations map[Point][]Point

// Proxy for infinity, just a high number
var INFINITE int = 999999

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample1.txt" // 7036, 45
	fname = "sample2.txt"  // 11048, 64
	fname = "input.txt"    // 99460, 500
	data, _ := ioutil.ReadFile(fname)
	rows = bytes.Split(data, []byte("\n"))

	// Find start and end points
	for ri := 0; ri < len(rows); ri++ {
		for ci := 0; ci < len(rows[0]); ci++ {
			if rows[ri][ci] == 'S' {
				start = Point{ci, ri, -1}
			} else if rows[ri][ci] == 'E' {
				end = Point{ci, ri, 0}
			}
		}
	}
	fmt.Println("Start/end =", start, end)

	// Initialize global variables
	dist = map[Point]int{}             // lowest cost to each position
	sources = map[Point][]Point{}      // list of sources for each point
	destinations = map[Point][]Point{} // list of destinations from each point

	// Part 1: find the length of the shortest path from start to end
	shortestPath = traverse(start, end)
	fmt.Println("Part 1 =", shortestPath)

	// Part 2: count up how many squares were visited by any of the shortest
	// paths, uses destination links built up found by Djikstra algorithm in Part 1
	marked = map[Point]int{} // all cells visited towards best
	part2(start, []Point{})  // sets every visited cell in 'marked'
	fmt.Println("Part 2 =", len(marked))
}

// Adapted Djikstra algorithm from Day 18 to include directional state,
// i.e., the direction we were going in to get to the current cell.
// Returns length of shortest path, or high number if no route found.
func traverse(start, end Point) int {

	// Set up priority queue with just starting point
	q := []Point{start}

	// Set up maps of visited points, and shortest distance found to each point
	visited := map[Point]bool{}
	dist[start] = 0      // cost of getting to start is zero
	dist[end] = INFINITE // initial cost to end is infinity

	// Start exploring, while there are still points to explore
	for len(q) > 0 {

		// Sort the queue in ascending order, so shortest distance is at the
		// front, and remove the first point (i.e., closest one)
		sort.Slice(q, func(i, j int) bool {
			return dist[q[i]] < dist[q[j]]
		})
		p := q[0]
		q = q[1:]

		// Skip if already visited this
		if visited[p] {
			continue
		}

		// Mark this point as visited so we don't return there
		visited[p] = true

		// Look up/right/down/left, in that order
		adjacent := []Point{Point{p.x, p.y - 1, 0}, Point{p.x + 1, p.y, 1},
			Point{p.x, p.y + 1, 2}, Point{p.x - 1, p.y, 3}}
		for _, p1 := range adjacent {

			// Reject if hit a wall or already visited,
			if at(p1) == '#' || visited[p1] { // allow going into 'E'
				continue
			}

			// Add to queue
			q = append(q, p1)

			// Calculate cost of this next step, 1 plus 1000 if changing direction
			next_dist := dist[p] + 1
			if p1.dir != p.dir {
				next_dist += 1000
			}

			// Update lowest distance to this next point
			d := dist[p1] // will be zero if not found
			if d == 0 || next_dist < d {
				dist[p1] = next_dist
			}

			// Update lists of sources and destinations
			sources[p1] = append(sources[p1], p)
			destinations[p] = append(destinations[p], p1)
		}
	}

	// Return value is lowest distance to ending point, will be
	// infinite if not reached. There may be multiple directions into
	// the end point, so find lowest
	ld := INFINITE            // lowest distance
	for i := 0; i <= 3; i++ { // each direction
		p := Point{end.x, end.y, i}
		d := dist[p]
		if d > 0 && d < ld {
			ld = d
		}
	}
	return ld
}

// For Part 2, find out how many tiles were traversed on any of the routes
// leading to the lowest cost. Call this function with the start point and
// an empty list of points, will recursively call for each branch, and build
// up all paths that reach the end with lowest distance. All the cells along
// these paths are then marked.
func part2(p Point, path []Point) {

	// If reached end and cost is the best, mark all the points along the path
	if p.x == end.x && p.y == end.y {
		//fmt.Println("reached end, cost =", dist[p])
		if dist[p] == shortestPath {
			for _, q := range path {
				marked[Point{q.x, q.y, 0}]++ // ignore direction
			}
		}
		return
	}

	// Follow each destination that was taken out of this cell
	for _, p1 := range destinations[p] {
		path1 := make([]Point, len(path)) // todo: make one extra slot?
		copy(path1, path)
		path1 = append(path1, p)
		part2(p1, path1)
	}
}

// Return char at given location, assumes we will always be in range
func at(p Point) byte {
	return rows[p.y][p.x]
}
