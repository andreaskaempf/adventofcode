// Advent of Code 2022, Day 22
//
// Simulate movement on a 2D map, according to a list of instructions, which
// can either be to move n steps, or to turn 90 degrees left or right. There
// are obstacles to avoid, and one wraps around to the other side when walking
// off and edge. For Part 1, the map is in 2D. For Part 3, the map gets folded
// into a cube.
//
// AK, 22 and 26 Dec 2022

package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

// Keep track of what is at each location
var tiles map[Point]byte

// Keep track of the minimum X and Y values for each row/col
var minX, maxX, minY, maxY map[int]int

// List of instructions, left and right encoded as -1 and -2
var instructions []int

const LEFT int = -1
const RIGHT int = -2

func main() {

	// Read the input file: map until blank line, then set of instructions
	fname := "sample.txt"
	fname = "input.txt"
	readMap(fname)

	// Do part 1
	part1()

	// Do part 2
	part2() // only works for input, not sample
}

// Part 1: follow instructions, moving around 2-d space, wrapping as necessary
func part1() {

	// Start in the first open tile on the first row, facing right
	y := 1
	x := minX[1]
	assert(tiles[Point{x, y}] == '.', "First tile is not open!")
	dir := 90 // start facing right

	// Process instructions, simplify by adding up all moves in each direction?
	for i := 0; i < len(instructions); i++ {

		// Process each turn or movement
		inst := instructions[i]
		if inst == RIGHT { // Rotate right
			dir = ifElse(dir == 270, 0, dir+90)
		} else if inst == LEFT { // Rotate left
			dir = ifElse(dir == 0, 270, dir-90)
		} else if dir == 90 { // Move right
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				x1++
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to beginning
					x1 = minX[y1]
				}
				if tiles[Point{x1, y1}] == '.' {
					x = x1
				} else {
					break
				}
			}

		} else if dir == 270 { // Move left
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				x1--
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to end
					x1 = maxX[y1]
				}
				if tiles[Point{x1, y1}] == '.' {
					x = x1
				} else {
					break
				}
			}

		} else if dir == 0 { // Move up
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				y1--
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to bottom
					y1 = maxY[x1]
				}
				if tiles[Point{x1, y1}] == '.' {
					y = y1
				} else {
					break
				}
			}

		} else if dir == 180 { // Move down
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				y1++
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to top
					y1 = minY[x1]
				}
				if tiles[Point{x1, y1}] == '.' {
					y = y1
				} else {
					break
				}
			}
		}

	}

	//fmt.Printf("Final x = %d, y = %d, dir = %d (s/b 8, 6, dir 90)\n", x, y, dir)
	facing := map[int]int{0: 3, 90: 0, 180: 1, 270: 2}
	score := 1000*y + 4*x + facing[dir]
	fmt.Println("Part 1: score (s/b 6032, 36518) =", score)
}

// Part 2: same as part 1, but wrap around cube instead of 2-d space.
// Done for main input only, and cube layout is hard-coded (so will not work
// for other layouts). Each face of the cube is assigned a letter:
//        _____ _____
//       |     |     |
//       |  A  |  B  |
//       |_____|_____|
//       |     |
//       |  C  |
//  _____|_____|
// |     |     |
// |  D  |  E  |
// |_____|_____|
// |     |
// |  F  |
// |_____|
//

func part2() {

	// Start in the first open tile on the first row (top left cell on face A),
	// facing right
	y := 1
	x := minX[1]
	assert(tiles[Point{x, y}] == '.', "First tile is not open!")
	dir := 90 // start facing right

	// Process instructions, simplify by adding up all moves in each direction?
	for _, inst := range instructions {

		// Get new coordinates and direction, and cell contents, if we did
		// that move. Update coordinates and direction if we would be moving
		// into an empty space (i.e., contains period).
		if inst == RIGHT { // Rotate right
			dir = ifElse(dir == 270, 0, dir+90)
		} else if inst == LEFT { // Rotate left
			dir = ifElse(dir == 0, 270, dir-90)
		} else {
			x1, y1, dir1, c := move(inst, x, y, dir)
			assert(c == '.' || c == '#', "Invalid char in map")
			if c == '.' {
				x = x1
				y = y1
				dir = dir1
			}
		}
	}

	// Final score
	facing := map[int]int{0: 3, 90: 0, 180: 1, 270: 2}
	score := 1000*y + 4*x + facing[dir]
	fmt.Println("Part 2 (s/b 144019):", score)
}

// Simulate a move for Part 2, given the number of steps to move in the
// current direction, current x,y location on the current cube face (oriented
// per original 2D layout), the current direction (relative to this cube face
// in original orientation). Returns the new x,y coordinates (relative to the
// original 2D layout), the direction (which may have changed as a result of
// wrapping around the cube), and the contents of the new cell. Stops moving
// before it hits a "wall" (cell with #).
func move(n, x, y, dir int) (int, int, int, byte) {

	// Initialize variables to relative coordinates on the current face
	face := whichFace(x, y)   // which face we are on A-F
	x1, y1 := relCoords(x, y) // coordinates relative to this face

	// Do each step of the move
	for k := 0; k < n; k++ {

		// Remember position before move, in case we hit a wall
		prevX := x1 // relative!
		prevY := y1
		prevDir := dir
		prevFace := face

		// Simulate one step in the current direction, on the current face
		if dir == 90 { // right
			x1++
		} else if dir == 270 { // left
			x1--
		} else if dir == 0 { // up
			y1--
		} else if dir == 180 { // down
			y1++
		} else {
			panic("Invalid direction!")
		}

		// If we are now off an edge of the face, we need to move to another face,
		// adjusting x and y and possibly direction
		if x1 < 1 { // moved off left edge

			if face == 'A' { // now on D, upside down
				face = 'D'
				y1 = faceSize - y1 + 1
				x1 = 1   // first column of D
				dir = 90 // now moving right

			} else if face == 'B' { // move B to A
				face = 'A'
				x1 = faceSize // y and dir do not change

			} else if face == 'C' { // now on D, rotated 90 deg right
				face = 'D'
				x1 = y1
				y1 = 1
				dir = 180 // now moving down

			} else if face == 'D' { // now on A, upside down
				face = 'A'
				x1 = 1
				y1 = faceSize - y1 + 1
				dir = 90 //  now moving right

			} else if face == 'E' { // now on D, nothing to do
				face = 'D'
				x1 = faceSize

			} else if face == 'F' { // now on A, rotated 90 degrees
				face = 'A'
				x1 = y1
				y1 = 1
				dir = 180 // now moving down
			}

		} else if x1 > faceSize { // moved off right edge

			if face == 'A' {
				face = 'B'
				x1 = 1

			} else if face == 'B' { // move B to E, upside down
				face = 'E'
				x1 = faceSize // right edge
				y1 = faceSize - y1 + 1
				dir = 270 // now moving left

			} else if face == 'C' {
				face = 'B'
				x1 = y1
				y1 = faceSize
				dir = 0 // now moving up

			} else if face == 'D' { // now on E continuing same direction
				face = 'E'
				x1 = 1

			} else if face == 'E' { // E -> B, upside down
				face = 'B'
				x1 = faceSize
				y1 = faceSize - y1 + 1
				dir = 270 // moving left

			} else if face == 'F' { // F -> E, rotated 90 degrees
				face = 'E'
				x1 = y1
				y1 = faceSize
				dir = 0 // now moving up
			}

		} else if y1 < 1 { // moved off top edge

			if face == 'A' {
				face = 'F'
				y1 = x1
				x1 = 1
				dir = 90 // now moving right

			} else if face == 'B' {
				face = 'F'
				y1 = faceSize
				dir = 0 // still moving up

			} else if face == 'C' {
				face = 'A'
				y1 = faceSize
				dir = 0 // still moving up

			} else if face == 'D' {
				face = 'C'
				y1 = x1
				x1 = 1
				dir = 90

			} else if face == 'E' {
				face = 'C'
				y1 = faceSize
				dir = 0 // still moving up

			} else if face == 'F' {
				face = 'D'
				y1 = faceSize
				dir = 0 // still moving up
			}

		} else if y1 > faceSize { // moved off bottom edge

			if face == 'A' {
				face = 'C'
				y1 = 1
				dir = 180 //still

			} else if face == 'B' {
				face = 'C'
				y1 = x1
				x1 = faceSize
				dir = 270 // now moving left

			} else if face == 'C' {
				face = 'E'
				y1 = 1
				dir = 180 // still moving down

			} else if face == 'D' {
				face = 'F'
				y1 = 1
				dir = 180

			} else if face == 'E' {
				face = 'F'
				y1 = x1
				x1 = faceSize
				dir = 270 // moving left

			} else if face == 'F' {
				face = 'B'
				y1 = 1
				dir = 180 // moving down
			}
		}

		// Convert back to absolute coordinates, check if we have hit a wall,
		// return last position if we have
		ax, ay := absCoords(x1, y1, face)
		if tiles[Point{ax, ay}] == '#' {
			x1, y1 = absCoords(prevX, prevY, prevFace)
			return x1, y1, prevDir, tiles[Point{x1, y1}]
		}

	}

	// Convert back to absolute coordinates and return result
	x1, y1 = absCoords(x1, y1, face)
	return x1, y1, dir, tiles[Point{x1, y1}]
}

// For part 2
const faceSize = 50 // width/height of each square face

// Which face (A-F) is a given absolute coordinate on? Assumes
// this layout, with each face 50 units square:
// 1   50 51 100 101-150
//
//	      _____ _____
//	     |     |     |
//	     |  A  |  B  |
//	     |_____|_____|
//	     |     |
//	     |  C  |
//	_____|_____|
// |     |     |
// |  D  |  E  |
// |_____|_____|
// |     |
// |  F  |
// |_____|
//
// Gives an error if the point is not on a valid face.

func whichFace(x, y int) rune {
	faces := [][]int{
		[]int{51, 100, 1, 50},    // A
		[]int{101, 150, 1, 50},   // B
		[]int{51, 100, 51, 100},  // C
		[]int{1, 50, 101, 150},   // D
		[]int{51, 100, 101, 150}, // E
		[]int{1, 50, 151, 200},   // F
	}
	for i := 0; i < len(faces); i++ {
		f := faces[i]
		if x >= f[0] && x <= f[1] && y >= f[2] && y <= f[3] {
			return rune(i + 'A')
		}
	}
	fmt.Printf("ERROR: point %d,%d is not on a face\n", x, y)
	panic("Invalid point")
}

// Older version, does not give error when a point is out of range
func whichFaceOld(x, y int) rune {
	if y <= faceSize {
		return ifElse(x > faceSize*2, 'B', 'A')
	} else if y <= faceSize*2 {
		return 'C'
	} else if y <= faceSize*3 {
		return ifElse(x > faceSize, 'E', 'D')
	} else {
		return 'F'
	}
}

// Return the relative coordinates, i.e., 1..50, works for any face
func relCoords(x, y int) (int, int) {
	assert(tiles[Point{x, y}] != 0, "relCoords: point not on map!")
	//return x % faceSize, y % faceSize
	return ((x - 1) % faceSize) + 1, ((y - 1) % faceSize) + 1
}

// Return the absolute coordinates for a pair of 1..50 relative coordinates
// and a face
func absCoords(x, y int, face rune) (int, int) {

	x0 := x
	y0 := y

	// Vertical adjustment
	if face == 'C' {
		y += faceSize
	} else if face == 'D' || face == 'E' {
		y += faceSize * 2
	} else if face == 'F' {
		y += faceSize * 3
	}

	// Horizontal adjustment
	if face == 'A' || face == 'C' || face == 'E' {
		x += faceSize
	} else if face == 'B' {
		x += faceSize * 2
	}

	// Check
	// Return adjusted coordinates
	if tiles[Point{x, y}] == 0 {
		fmt.Printf("absCoords(%d,%d): %d,%d not on map!\n", x0, y0, x, y)
	}
	if whichFace(x, y) == 0 {
		fmt.Printf("absCoords(%d,%d): %d,%d not on a face!\n", x0, y0, x, y)
	}
	return x, y
}

// Read the input file: map until blank line, then set of instructions,
// also set min/max X/Y
func readMap(fname string) {

	// Read the input file: map until blank line, then set of instructions
	readingMap := true
	tiles = map[Point]byte{}
	var x, y int
	instructions = []int{}
	for _, l := range readLines(fname) {
		if len(l) == 0 { // blank line means instructions come next
			readingMap = false
		} else if readingMap {
			for x = 0; x < len(l); x++ {
				if l[x] != ' ' {
					tiles[Point{x + 1, y + 1}] = l[x]
				}
			}
			y++
		} else { // instructions after blank line
			instructions = parseInstructions(l)
		}
	}

	// Get the min/max col in each row and min/max row in each column
	minX = map[int]int{}
	maxX = map[int]int{}
	minY = map[int]int{}
	maxY = map[int]int{}
	for p, _ := range tiles {
		if minX[p.y] == 0 || p.x < minX[p.y] {
			minX[p.y] = p.x
		}
		if p.x > maxX[p.y] {
			maxX[p.y] = p.x
		}
		if minY[p.x] == 0 || p.y < minY[p.x] {
			minY[p.x] = p.y
		}
		if p.y > maxY[p.x] {
			maxY[p.x] = p.y
		}
	}
}

// Parse string of instructions into a list of numbers, where -1 means turn
// left and -2 means turn right, other numbers are the number of steps to move
func parseInstructions(s string) []int {
	result := []int{}
	n := 0 // the current number
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' { // parse numbers, may be multi digit
			n = n*10 + int(c-'0')
		} else if c == 'L' || c == 'R' { // L and R are encoded as -1 and -2
			if n > 0 {
				result = append(result, n)
				n = 0
			}
			result = append(result, ifElse(c == 'L', -1, -2))
		} else {
			panic("Invalid instruction!")
		}
	}
	if n > 0 {
		result = append(result, n)
	}
	return result
}
