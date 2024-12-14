// Advent of Code 2024, Day 14
//
// Given a list of "robots" in a ~100x100 2D space, each with an x,y position
// and an dx,dy per-second velocity, simulate 100 iterations of the robots
// moving around, wrapping around at the edges. For Part 1, answer is the
// product of the number of robots in each of the four quadrants after 100
// iterations. For Part 2, it's the first iteration that yields ASCII art of a
// Christmas tree. First part was easy simulation. For Part 2, just drew maps
// of configuration after each iteration where there appeared to be lots of
// "robots" in one row, this was sufficient to find the answer by visual
// inspection.
//
// AK, 14 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// An x,y pair, so we can use as key for a map
type Point struct {
	x, y int
}

// A robot, with current position and velocity
type Robot struct {
	pos, v Point
}

// Global variables: list of robots, width & height
var robots []Robot
var width, height int

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

	// Parse list of robots
	robots = []Robot{}
	patt := regexp.MustCompile("-?[0-9]+")
	for _, l := range lines {
		x := patt.FindAllString(l, -1)
		if len(x) == 4 {
			r := Robot{Point{atoi(x[0]), atoi(x[1])}, Point{atoi(x[2]), atoi(x[3])}}
			robots = append(robots, r)
		} else {
			fmt.Println("Ignoring:", l, x)
		}
	}

	// Simulate movement for 100 seconds, in space that is 101 wide, 103 tall
	width = 101
	height = 103
	var part1_ans int
	for t := 0; t < 10000; t++ { // 100 for Part 1
		for i := 0; i < len(robots); i++ {

			// Move robot
			robots[i].pos.x += robots[i].v.x
			robots[i].pos.y += robots[i].v.y

			// If out of bounds, wrap around
			for robots[i].pos.x >= width {
				robots[i].pos.x -= width
			}
			for robots[i].pos.x < 0 {
				robots[i].pos.x += width
			}
			for robots[i].pos.y >= height {
				robots[i].pos.y -= height
			}
			for robots[i].pos.y < 0 {
				robots[i].pos.y += height
			}
		}

		// Calculate Part 1 at iteration 100
		if t == 100 {
			part1_ans = part1()
		}

		// For Part 2, had to find the first iteration that created a map
		// that looked like a Christmas tree, turned out to be at iteration
		// 7892. Found this by printing pictures of any configuration that had
		// lots of "robots" in any single row.
		if xmasTree() {
			fmt.Println("After", t+1)
			drawMap()
		}
	}

	// Answer for Part 1 is just the product of robots in each quadrant,
	// calculated at iteration 100
	fmt.Println("Part 1:", part1_ans) // 224969976
}

// Answer for Part 1, called at iteration 100, counts robots in each
// quadrant, calculates product
func part1() int {

	// Count number of robots in each quadrant, ignoring middle lines
	var q1, q2, q3, q4 int
	for _, r := range robots {

		// Determine if the robot is in the top/bottom or
		// left/right halves, ignoring center line
		hmid := (width - 1) / 2
		leftHalf := r.pos.x < hmid
		rightHalf := r.pos.x >= hmid+1
		vmid := (height - 1) / 2
		topHalf := r.pos.y < vmid
		bottomHalf := r.pos.y >= vmid+1

		// Add to count for relevant quadrant
		if leftHalf {
			if topHalf {
				q1++
			} else if bottomHalf {
				q2++
			}
		} else if rightHalf {
			if topHalf {
				q3++
			} else if bottomHalf {
				q4++
			}
		}
	}
	return q1 * q2 * q3 * q4
}

// Does this map possibly look like a Christmas tree?
func xmasTree() bool {

	// Count up number of robots in each row
	rcounts := map[int]int{}
	for _, r := range robots {
		rcounts[r.pos.y]++
	}

	// If any row has lots of robots, this might be a picture
	for _, c := range rcounts {
		if c > 30 { // found by trial and error
			return true
		}
	}
	return false
}

// Draw map as digits, used for debugging in Part 1 but also
// useful for Part 2!
func drawMap() {

	// Find number of robots at each coordinate
	c := map[Point]int{}
	for _, r := range robots {
		c[r.pos]++
	}

	// Draw each coordinate
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			n := c[Point{x, y}]
			if n == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d", n)
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
