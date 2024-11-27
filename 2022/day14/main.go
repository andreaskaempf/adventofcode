// Advent of Code 2022, Day 14
//
// Simulate grains of sand dropping from a hole into 2-d space. For Part 1,
// count how many grains of sand before they start dropping of edges of
// existing rock. For Part 2, add a "foor" below  bottom layer of rock, and
// count how many grains of sand before a pyramid is built, and the hole at
// the top becomes blocked.
//
// AK, 14 Dec 2022

package main

import (
	"fmt"
	"strings"
)

// A point in 2-d space
type Point struct {
	x, y int
}

func main() {

	// Read data into a list of paths, each consisting of x,y points
	//fname := "sample.txt"
	fname := "input.txt"
	paths := [][]Point{} // a list of lists of points
	for _, l := range readLines(fname) {
		path := []Point{}
		for _, s := range strings.Split(l, " -> ") {
			xy := strings.Split(s, ",")
			path = append(path, Point{atoi(xy[0]), atoi(xy[1])})
		}
		paths = append(paths, path)
	}

	// Turn paths into a sparse 2-d array of points already blocked by "rock"
	space := map[Point]int{}
	for _, path := range paths {
		for i := 1; i < len(path); i++ { // each segment

			// Get start & end of segment, and dx/dy
			p0 := path[i-1] // start of segment
			p1 := path[i]   // end of segment
			var dx, dy int  // default zero
			if p1.x > p0.x {
				dx = 1
			} else if p1.x < p0.x {
				dx = -1
			}
			if p1.y > p0.y {
				dy = 1
			} else if p1.y < p0.y {
				dy = -1
			}

			// Mark each point on the segment as filled in space
			p := p0 // this point
			for {
				space[p] = 1 // mark location as filled with rock
				p.x += dx    // adjust position by one
				p.y += dy
				if (dx > 0 && p.x > p1.x) || (dx < 0 && p.x < p1.x) ||
					(dy > 0 && p.y > p1.y) || (dy < 0 && p.y < p1.y) {
					break // stop if out of bounds of segment
				}
			}
		}
	}

	// Find the bottom of the existing rock, and lowest point overall
	bottom := map[int]int{} // for Part 1: lowest point at each x
	lowest := 0             // for Part 2: lowest point overall
	for p, _ := range space {
		if p.y > bottom[p.x] { // for part 1
			bottom[p.x] = p.y
		}
		if p.y > lowest { // for part 2
			lowest = p.y
		}
	}

	// For Part 2, add a layer of rock at lowest - 2
	for i := 0; i < 1000; i++ {
		space[Point{i, lowest + 2}] = 3
	}

	// Start simulation
	n := 0             // number of grains released
	part1done := false // haven't reported part 1 yet
	for {

		// Create a grain of sand at starting location
		g := Point{500, 0}
		n++

		// Try to move it down, diag left or right
		for {
			blocked := false
			if !filled(space, g.x, g.y+1) {
				g.y += 1
			} else if !filled(space, g.x-1, g.y+1) {
				g.x--
				g.y++
			} else if !filled(space, g.x+1, g.y+1) {
				g.x++
				g.y++
			} else {
				space[g] = 2 // mark as filled with a grain of sand
				blocked = true
			}

			// Part 1: report grains of sand before start falling below edges
			if g.y > bottom[g.x] && !part1done { // i.e., past bottom
				fmt.Println("Part 1 (24, 1133):", n-1)
				part1done = true // so we don't report it again
				//visualize(space)
			}

			// Stop when grain of sand could not be moved
			if blocked {
				break
			}
		}

		// Part 2: stop when grain of sand could not be moved
		if g.x == 500 && g.y == 0 {
			fmt.Println("Part 2 (93, 27566):", n)
			//visualize(space) // Uncomment to see visualization
			break
		}
	}
}

// Determine if point in space is filled, i.e., that element in the
// sparse matrix is not empty
func filled(space map[Point]int, x, y int) bool {
	_, ok := space[Point{x, y}]
	return ok
}

// Visualize the space,for debugging
func visualize(space map[Point]int) {
	for r := 0; r < 20; r++ {
		for c := 480; c < 530; c++ {
			x, ok := space[Point{c, r}]
			if ok && x > 0 {
				fmt.Printf("%d", x)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("|\n")
	}
}
