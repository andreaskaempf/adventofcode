// Advent of Code 2022, Day 10
//
// Simulate accumulator register during a series of given ADD or NOOP
// instructions, and report the accumulator values during selected clock
// cycles (Part 1). For part 2, simulate drawing of pixels on a screen,
// using the sequence of acculator values, and report the eight letters
// that appear.
//
// AK, 10 Dec 2022

package main

import (
	"fmt"
	"strings"
)

func main() {

	// Part 1: process each line and simulate the add/noop instructions,
	// building up accumulator over each cycle (addx is two cycles)
	lines := readLines("input.txt") // for part 1: "sample.txt"
	X := 1                          // the register starts at 1
	acc := []int{}                  // value of the register *during* each cycle
	for _, l := range lines {
		words := strings.Split(l, " ")
		if words[0] == "noop" {
			acc = append(acc, X) // no change, one cycle
		} else { // expect "addx <n>"
			acc = append(acc, X, X) // takes two cycles, still at old value
			X += atoi(words[1])     // the new value
		}
	}

	// Part 1: sumproduct of certain cycles and register values
	part1 := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		part1 += i * acc[i-1]
	}
	fmt.Println("Part 1 (s/b 13140):", part1)

	// For part 2, draw pixels on a 40x6 virtual screen, one pixel per cycle
	screen := make([]int, 6*40, 6*40)
	var h int                       // current horizontal position
	for t := 0; t < len(acc); t++ { // each cycle
		if abs(acc[t]-h) <= 1 { // if acc close to horizontal position,
			screen[t] = 1 // turn on pixel
		}
		h++         // next horizontal position on screen
		if h > 39 { // loop back to beginning of row
			h = 0
		}
	}

	// Print the final screen, shows EZFPRAKL
	fmt.Println("Part 2:")
	for p := 0; p < len(screen); p++ { // each pixel
		if p > 0 && p%40 == 0 {
			fmt.Print("\n") // start new row
		}
		fmt.Print(ifElse(screen[p] == 0, " ", "X"))
	}
}
