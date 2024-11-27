// Data structures and read input files

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// An 3D space, with on/off state
type Cube struct {
	onOff                  int32 // cube turned on by instruction
	x0, x1, y0, y1, z0, z1 int32
}

// A point in space (only used in Part 1)
type Point struct {
	x, y, z int32
}

// Read "instructions" from an input file
func readData(fname string) []Cube {

	insts := []Cube{}

	// Read each line of input file, each looking like:
	// on x=11..13,y=11..13,z=11..13
	f, _ := os.Open(fname)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		t := scanner.Text()

		// on/off beginning of line
		instr := Cube{}
		parts := strings.Split(t, " ")
		if parts[0] == "on" {
			instr.onOff = 1
		}

		// x/y/z coords (x=11..13)
		for _, xyz := range strings.Split(parts[1], ",") {
			lr := strings.Split(xyz, "=")
			lohi := strings.Split(lr[1], "..")
			if lr[0] == "x" {
				instr.x0 = atoi(lohi[0])
				instr.x1 = atoi(lohi[1])
			} else if lr[0] == "y" {
				instr.y0 = atoi(lohi[0])
				instr.y1 = atoi(lohi[1])
			} else if lr[0] == "z" {
				instr.z0 = atoi(lohi[0])
				instr.z1 = atoi(lohi[1])
			} else {
				panic("Bad letter")
			}

		}

		insts = append(insts, instr)
	}
	return insts
}

// Parse an integer
func atoi(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic("Bad number")
	}
	return int32(i)
}
