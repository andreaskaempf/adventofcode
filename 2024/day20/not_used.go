// Legacy code from earlier attempts at this problem.

package main

import (
	"fmt"
	"maps"
	"sort"
)

// Temporarily open wall positions, for 2 cycles
var opened map[Point]int

// Find shortest path from start to end, walking through walls at t=happy_time
// or happy_time+1
func djikstra(happy_time int) int {

	// Set up priority queue with just starting point
	q := []Point{start}

	// Set up maps of visited points, and shortest distance found to each point
	visited := map[Point]bool{}
	dist := map[Point]int{}
	dist[start] = 0    // cost of getting to start is zero
	dist[end] = 999999 // initial cost to end is infinity

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

		// Look up/right/down/left
		for _, p1 := range adjacents(p) {

			// Reject if already visited
			if visited[p1] {
				continue // does not make a difference
			}

			// Reject if hitting a wall, unless at "happy time" or one tick after
			t := dist[p]
			happy := happy_time >= 0 && (t == happy_time || t == happy_time+1)
			if at(p1) == '#' && !happy { // cannot walk through walls
				continue // except at "happy time"
			}

			// Add to queue
			q = append(q, p1)

			// Update lowest distance to this next point
			next_dist := dist[p] + 1
			d, ok := dist[p1]
			if !ok || next_dist < d {
				dist[p1] = next_dist
			}
		}
	}

	// Return value is lowest distance to ending point, will be
	// infinite if not reached
	return dist[end]
}

// Get shortest path from this point, allowing walk through obstacles/walls
// at ht and ht+1 ("happy time")
func shortestPath1(p Point, d int, visited map[Point]bool, ht int) int {

	// Stop when reached end
	if p.x == end.x && p.y == end.y {
		return d
	}

	// Explore all adjacent points from here
	visited[p] = true
	d1 := 999999 // shortest distance from adjacencies
	for _, p1 := range adjacents(p) {

		// Skip if already visited
		if visited[p1] {
			continue
		}

		// Can't go through walls, except at happy time
		happy := ht >= 0 && (d == ht || d == ht+1)
		if at(p1) == '#' && !happy { // can't go here if a wall
			continue
		}

		// Proceed to next step and update shortest distance
		vis := maps.Clone(visited)            // make a new copy of visited map
		dx := shortestPath1(p1, d+1, vis, ht) // recursive call from next step
		if dx < d1 {
			d1 = dx
		}
	}

	return d1
}

// Get shortest path from this point, allowing for temporarily opened points
// TODO time contraint on temporarily opened points
func shortestPath(p Point, d int, visited map[Point]bool) int {

	// Stop when reached end
	if p.x == end.x && p.y == end.y {
		return d
	}

	// Explore all adjacent points from here
	visited[p] = true
	d1 := 0 // shortest distance from adjacencies
	for _, p1 := range adjacents(p) {

		// Skip if already visited
		if visited[p1] {
			continue
		}

		// If temporarily opened, make sure it's still open, also allow
		// if not on a wall
		nPasses, tempOpen := opened[p1]
		if tempOpen {
			if nPasses > 0 {
				opened[p1]-- // reduce number of passes through here
			} else { // passes used up, can't go through here any more
				continue
			}
		} else if at(p1) == '#' { // can't go here if a wall
			continue
		}

		// Proceed to next step and update shortest distance
		vis := maps.Clone(visited)       // make a new copy of visited map
		dx := shortestPath(p1, d+1, vis) // recursive call from next step
		if d1 == 0 || dx < d1 {
			d1 = dx
		}
	}

	return d1
}

func old_main() {

	// Find starting length of path (will be same as number of steps along default path)
	shortest := djikstra(-1)
	fmt.Println("Dijkstra shortest =", shortest)
	fmt.Println("Recursive shortest =", shortestPath(start, 0, map[Point]bool{}))

	// Find "cheats", i.e., any pair of wall cells that if removed would
	// shorten the path. Do this by finding all pairs of steps along the path
	// that are exactly 2 or 3 apart (i.e., 1 or 2 walls in-between), and
	// calculating the difference in distance from removing these.
	cheats := map[int]int{}  // for debugging
	opened = map[Point]int{} // global
	saved100 := 0            // for Part 1, the number of cheats that saved 100+ steps
	for i := 0; i < len(path); i++ {
		for j := i + 1; j < len(path); j++ {

			// Only look at pairs of points that are 2 or 3 apart on path
			p1 := path[i]
			p2 := path[j]
			d := dist(p1, p2)
			if d < 2 || d > 3 { // allow 2 or 3
				continue // only looks at horiz or vertical, not diagonal
			}

			// Get point(s) in-between the path points, skip if both are periods
			w1, w2 := pointsInBetween(p1, p2)
			if at(w1) == '.' && (w2 == Point{0, 0, 0} || at(w2) == '.') {
				continue
			}

			// Temporarily "open up" these positions for up to 2 cycles, and
			// measure resulting distance
			opened[w1] = 2
			opened[w2] = 2 //may be 0,0,0 if only one point in-between
			d = shortestPath(start, 0, map[Point]bool{})
			delete(opened, w1)
			delete(opened, w2)

			// If resulting distance is shorter, report successful cheat
			if d > 0 && d < shortest {
				saved := shortest - d
				fmt.Println("Cheat at", w1, w2, ", saved", saved)
				if saved >= 100 {
					saved100++
				}
				cheats[saved]++
			}
		}
	}
	fmt.Println("s/b 2:14, 4:14, 6:2 8:4 10:2 12:3 20 36 38 40 64 each 1")
	fmt.Println(cheats)
	//fmt.Println("Part 1 =", saved100)

}

func part1(opentime int) {

	// Find starting length of path (will be same as number of steps along default path)
	shortest := djikstra(0)
	fmt.Println("Dijkstra shortest =", shortest)
	fmt.Println("Recursive shortest =", shortestPath(start, 0, map[Point]bool{}))

	// Find "cheats", i.e., any pair of wall cells that if removed would
	// shorten the path. Do this by finding all pairs of steps along the path
	// that are exactly 2 or 3 apart (i.e., 1 or 2 walls in-between), and
	// calculating the difference in distance from removing these.
	cheats := map[int]int{}  // for debugging
	opened = map[Point]int{} // global
	saved100 := 0            // for Part 1, the number of cheats that saved 100+ steps
	for _, p := range path { // every point on path
		for _, w1 := range adjacents(p) { // every point next to it
			for _, w2 := range adjacents(w1) { // every point next to the adjacency

				// Temporarily "open up" these positions for up to 2 cycles, and
				// measure resulting distance
				opened[w1] = 2
				opened[w2] = 2 //may be 0,0,0 if only one point in-between
				d := shortestPath(start, 0, map[Point]bool{})
				delete(opened, w1)
				delete(opened, w2)

				// If resulting distance is shorter, report successful cheat
				if d > 0 && d < shortest {
					saved := shortest - d
					fmt.Println("Cheat at", w1, w2, ", saved", saved)
					if saved >= 100 {
						saved100++
					}
					cheats[saved]++
				}
			}
		}
	}
	fmt.Println("s/b 2:14, 4:14, 6:2 8:4 10:2 12:3 20 36 38 40 64 each 1")
	fmt.Println(cheats)
	//fmt.Println("Part 1 =", saved100)
}

// Is there a cheat between these two points on the path?
// Horizontal or vertical distance must be 2 or 3, and there
// must be at least one wall between them. If this is the case,
// remove those two cells, and find new path.
func pointsInBetween(a, b Point) (Point, Point) {

	// Must be 3 apart, to allow wall of 1 or 2 thick
	d := dist(a, b) // only horiz or vertical
	if d < 2 || d > 3 {
		panic("pointsInBetween must have points that are 2 or 3 apart")
	}

	// Find the two points in-between
	var w1, w2 Point
	x1 := min(a.x, b.x)
	y1 := min(a.y, b.y)
	if a.x == b.x { // vertical, same x
		w1 = Point{x1, y1 + 1, 0}
		w2 = Point{x1, y1 + 2, 0}
	} else if a.y == b.y { // horizontal, same y
		w1 = Point{x1 + 1, y1, 0}
		w2 = Point{x1 + 2, y1, 0}
	} else {
		panic("Invalid check")
	}
	if d == 2 {
		//w2 = Point{0, 0, 0}
	}
	return w1, w2
}

// Is there a cheat between these two points on the path?
// Horizontal or vertical distance must be 2 or 3, and there
// must be at least one wall between them. If this is the case,
// remove those two cells, and find new path.
func isCheat(a, b Point) (Point, Point, int) {

	// Must be 2 or 3 apart, to allow wall of 1 or 2 thick
	d := dist(a, b) // only horiz or vertical
	w0 := Point{0, 0, 0}
	if !(d == 2 || d == 3) {
		return w0, w0, -1
	}

	// Find the two points in-between
	var w1, w2 Point
	x1 := min(a.x, b.x)
	y1 := min(a.y, b.y)
	if a.x == b.x { // vertical, same x
		w1 = Point{x1, y1 + 1, 0}
		w2 = Point{x1, y1 + 2, 0}
	} else if a.y == b.y { // horizontal, same y
		w1 = Point{x1 + 1, y1, 0}
		w2 = Point{x1 + 2, y1, 0}
	} else {
		panic("Invalid check")
	}

	// At least one must be a wall
	//WHY NOT FINDING ANY PAIRS WITH 2 #?
	isWall := at(w1) == '#' || at(w2) == '#'
	if !isWall {
		return w0, w0, -1
	}

	// Set these two points to periods to 'free' them, and calculate
	// distance before restoring the points
	oldW1 := at(w1)
	oldW2 := at(w2)
	rows[w1.y][w1.x] = '.'
	rows[w2.y][w2.x] = '.'
	d = djikstra(0)
	rows[w1.y][w1.x] = oldW1
	rows[w2.y][w2.x] = oldW2

	// Return the two wall points and shortest distance
	//plot(w1, w2)
	return w1, w2, d
}

func plot(p1, p2 Point) {
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			ch := string(at(Point{x, y, 0}))
			if p1.x == x && p1.y == y {
				fmt.Printf("\033[41m%s\033[0m", ch)
			} else if p2.x == x && p2.y == y {
				fmt.Printf("\033[42m%s\033[0m", ch)
			} else {
				fmt.Print(ch)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

// Distance between two points, just horizontal or vertical
func dist(a, b Point) int {
	if a.x == b.x { // same col
		return abs(a.y - b.y)
	} else if a.y == b.y { //  same row
		return abs(a.x - b.x)
	} else {
		return -1 // no valid distance, e.g. diagonal
	}
}

// Get adjacent points
func adjacents(p Point) []Point {
	return []Point{Point{p.x, p.y - 1, 0}, Point{p.x + 1, p.y, 0},
		Point{p.x, p.y + 1, 0}, Point{p.x - 1, p.y, 0}}
}
