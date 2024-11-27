// Day 4 of Advent of Code 2021
//
// Simulate game of bingo, by applying list of called numbers
// until first board wins (for part 2, also find the last board
// that wins).
//
// AK, 4 December 2021

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// List of numbers read out, from first line of input file
var Called []int

// Data structure for a board (5 rows of 5), and a list of boards,
// to be read from input file
type Row []int
type Board []Row

var Boards []Board

// Convert string to a list of numbers
func strToInts(s string, delim string) []int {
	nums := []int{}
	s = strings.TrimPrefix(s, " ") // remove leading spaces
	if delim == " " {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	nn := strings.Split(s, delim)
	for _, n := range nn {
		i, _ := strconv.Atoi(n) // TODO: check err
		nums = append(nums, i)
	}
	return nums
}

// Read data from file, populates Called with a list of the numbers called,
// and Boards with a list of boards
func readData() bool {

	var b Board                  // current board, being read now
	f, _ := os.Open("input.txt") // should really check err
	scanner := bufio.NewScanner(f)
	var line int // line number (initially 0)
	//var lastLineBlank bool // whether prev line was blank (initially false)
	for scanner.Scan() {
		line++
		t := scanner.Text()

		// First line contains numbers called, comma separated, otherwise
		// assume it's a row from a board, separated by blank lines
		if line == 1 {
			Called = strToInts(t, ",")
		} else if len(t) == 0 { // blank line: finish board or start new
			if b != nil { // save current board in list
				Boards = append(Boards, b)
			}
			b = Board{} // start a new board
		} else { // add row to current board
			b = append(b, strToInts(t, " "))
		}
	}
	fmt.Printf("%d lines read: %d boards, %d numbers called\n", line, len(Boards), len(Called))
	return true
}

// Check all the boards: each should be 5x5, with
// numbers in the range 0-99
func checkBoards() bool {
	for bi, b := range Boards {
		if len(b) != 5 {
			fmt.Printf("Board %d has %d rows!\n", bi+1, len(b))
			return false
		}
		for ri, r := range b {
			if len(r) != 5 {
				fmt.Printf("Board %d, row %d has %d elements!\n", bi+1, ri+1, len(r))
				return false
			}
			for _, n := range r {
				if n < 0 || n > 99 {
					fmt.Printf("Board %d, row %d has invalid elements %d!\n", bi+1, ri+1, n)
					return false
				}
			}
		}
	}
	fmt.Println("Boards are okay")
	return true
}

// Print one board
func printBoard(b Board) {
	for _, r := range b {
		for _, n := range r {
			fmt.Printf("%3d ", n)
		}
		fmt.Printf("\n")
	}
}

// Print all the boards
func printBoards() {
	for i, b := range Boards {
		fmt.Printf("\nBoard %d:\n", i+1)
		printBoard(b)
		if i == 2 {
			break
		}
	}
}

// Determine if board wins, by having selections across row,
// column, or diagonal
func boardWins(b Board) bool {

	// Check rows
	for _, r := range b {
		wins := true
		for _, c := range r {
			if c < 100 {
				wins = false
				break
			}
		}
		if wins {
			return true
		}
	}

	// Check columns
	ncols := len(b[0])
	for ci := 0; ci < ncols; ci++ { // each column
		wins := true
		for _, r := range b { // each row
			if r[ci] < 100 {
				wins = false
				break
			}
		}
		if wins {
			return true
		}
	}

	// Didn't win if no winning row/column found
	return false
}

// Sum up all the unmarked numbers on a board, i.e., those to which 100
// has not been added
func sumUnmarked(b Board) int {
	sum := 0
	for _, r := range b {
		for _, n := range r {
			if n <= 100 {
				sum += n
			}
		}
	}
	return sum
}

// Play the game: choose each number called, then
// check each board for that number,  setting it to
// n+100 on the board. If a board has a marked row or column,
// declare winner by returning score.
func playGame(stopFirstWin bool) int {

	// For part 2, keep track of which boards have won
	wins := 0
	boardAlreadyWon := make([]bool, len(Boards), len(Boards))

	// Call each number
	for callSeq, called := range Called {

		fmt.Println("\nCalling", called)

		// Look at each board, row by row, column within
		// each row, and add 100 if number matches called
		for bi, b := range Boards {

			// Go through board, change any values that match call
			changed := false
			for _, r := range b { // each row
				for ci, c := range r {
					if c == called {
						r[ci] += 100
						changed = true
					}
				}
			}

			// If board was changed, show it, and check if it's a winner.
			// The score of the winning board is calculated by finding the
			// sum of all unmarked numbers on that board, and multiplying
			// that sum by the number that was just called.
			if changed && !boardAlreadyWon[bi] {
				fmt.Printf("Board %d was changed to call %d:\n", bi+1, called)
				printBoard(b)
				if boardWins(b) {
					unmarked := sumUnmarked(b)
					score := unmarked * called
					fmt.Printf("Board %d wins on call number %d!\n", bi+1, callSeq+1)
					fmt.Printf("Score = %d unmarked * %d call = %d\n", unmarked, called, score)
					// In part 1, we stop at first win
					if stopFirstWin {
						return score
					}
					// Otherwise, add to counter, and stop if this is the
					// last board to win
					wins++
					boardAlreadyWon[bi] = true
					if wins == len(Boards) {
						fmt.Println("This is the last board to win!")
						return score
					}
				}
			}
		}
	}
	return 0
}

// Main execution
func main() {

	// Read data and show it
	readData()
	//printBoards()
	if !checkBoards() {
		fmt.Println("Bad data, aborted")
		return
	}

	// Play the game
	//playGame(true)  // part 1: stop at the first win
	playGame(false) // part 2: stop at the last board that wins
}
