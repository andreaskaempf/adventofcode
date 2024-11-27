// Advent of Code 2022, Day 02
//
// Simulate play of Rock/Scissors/Paper. In part 2, choose each move's
// strategy based on the pre-determined outcome in the input.
//
// AK, 2 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Read the input file
	//filename := "sample.txt"
	filename := "input.txt"
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines), "lines read")

	// Part 1: play game, assuming that:
	// First col: A for Rock, B for Paper, and C for Scissors
	// Second col:  X for Rock, Y for Paper, and Z for Scissors
	score := 0
	opShapes := map[byte]string{'A': "Rock", 'B': "Paper", 'C': "Scissors"}
	myShapes := map[byte]string{'X': "Rock", 'Y': "Paper", 'Z': "Scissors"}
	shapeScore := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3}
	for _, l := range lines {
		if len(l) < 3 {
			break
		}
		op := opShapes[l[0]]
		me := myShapes[l[2]]
		score += shapeScore[me]
		if op == me { // draw
			score += 3
		} else if defeats(me, op) {
			score += 6
		}
	}
	fmt.Println("Part 1 (15 for sample, 13809 with input):", score)

	// Part 2: find the combination of shapes that results in the outcome
	// predicted by second column:
	// X means you need to lose,
	// Y means you need to end the round in a draw, and
	// Z means you need to win
	score = 0
	for _, l := range lines {
		if len(l) < 3 {
			break
		}
		op := opShapes[l[0]]
		outcome := l[2] // X/Y/Z
		myShape := outcomeShape(op, outcome)
		score += shapeScore[myShape]
		if outcome == 'Y' { // draw
			score += 3
		} else if outcome == 'Z' { // win
			score += 6
		}
	}
	fmt.Println("Part 2 (12 for sample, 12316 with input):", score)
}

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
func defeats(me, op string) bool {
	return (me == "Rock" && op == "Scissors") ||
		(me == "Scissors" && op == "Paper") ||
		(me == "Paper" && op == "Rock")
}

// Get the shape you need to win/lose/draw against opponent's shape
// X = lose, Y = draw, Z = win
// Rock > Scissors > Paper
// and Paper > Rock
func outcomeShape(op string, outcome byte) string {
	if outcome == 'Z' { // lose
		if op == "Rock" {
			return "Paper"
		} else if op == "Scissors" {
			return "Rock"
		} else {
			return "Scissors"
		}
	} else if outcome == 'X' { // lose
		if op == "Paper" {
			return "Rock"
		} else if op == "Rock" {
			return "Scissors"
		} else {
			return "Paper"
		}
	} else { // draw
		return op
	}
}
