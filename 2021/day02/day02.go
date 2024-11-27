// Advent of Code, Day 2
//
// Reads a text file with "command value" on each line, and
// adjusts submarine state accordingly. Processes versions
// 1 and 2 of the problem at the same time.
//
// To compile and run:
// 	go build day02.go
//  ./day02
//
// Andreas Kaempf, 2 Dec 2021

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// State variables for problems 1 & 2 (all init to zero)
	var horiz1, depth1, horiz2, depth2, aim2 int

	// Read each line of input file
	f, _ := os.Open("input.txt") // should really check err
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		// Parse instruction and value
		words := strings.Split(scanner.Text(), " ")
		instruction := words[0]
		n, _ := strconv.Atoi(words[1]) // should check err

		// Process instructions for both problems
		if instruction == "forward" {
			horiz1 += n
			horiz2 += n
			depth2 += aim2 * n
		} else if instruction == "down" {
			depth1 += n
			aim2 += n
		} else if instruction == "up" {
			depth1 -= n
			aim2 -= n
		} else {
			fmt.Println("Bad instruction:", instruction)
		}
	}

	// Show results for both problems
	fmt.Printf("Problem 1: Ending pos = %d, depth = %d, product = %d\n", horiz1, depth1, horiz1*depth1)
	fmt.Printf("Problem 2: Ending pos = %d, depth = %d, aim = %d, product = %d\n", horiz2, depth2, aim2, horiz2*depth2)
}
