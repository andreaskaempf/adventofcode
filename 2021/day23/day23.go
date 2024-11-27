// Advent of Code, Day 23
//
// Find the most economical solution to a board game, involving
// 4 (later 8) pieces from random starting tunnels to ordered
// destination tunnels via a corridor, sort of like the Towers
// of Hanoi. Solved the first part (8 pieces) on paper, second
// part using recursive depth-first search, eliminating branches
// that exceeded best solution found so far.
//
// Summary of algorithm for depth-first search:
//
// Start with initial state

// Evaluate state, identify all possible moves
// - for every piece, that is not locked or already in final position, ...
// - look at every possible move:
//   - if in corridor, can it move to its destination (without blocking
//     a piece already there)?
//   - if in its tunnel, can it move to the corridor?
//   - in either case, check that it is not blocked
// - if no possible moves, check if final state reached => solution
//
// Try every possible move from that state
// - create a copy of the current state
// - move the piece there, counting steps
// - add this move to this copy of the game
// - recursively evaluate the resulting state copy
//
// AK, 23-26 December 2021
package main

import (
	"fmt"
)

// Display additional messages, including representation of board
// at every step
var verbose = false

// Structure for a chamber within a board, each containing letters or nil,
// element 0 closest to opening
type Chamber [cdepth]byte

// Structure for a board state
type Board struct {

	// Four chambers, A on left to D on right,
	// element 0 closest to opening
	chambers [4]Chamber

	// The corridor, with 11 positions:
	// 0, 1       left of first opening
	// 2, 4, 6, 8 openings to chambers A, B, C, D
	// 3, 5, 7    holding spaces between chambers
	// 9, 10      right of last opening
	corridor [11]byte

	// List of moves already made to get to this state
	moves []Move
}

// One move, for piece X from position A to position B, with number of steps
// Position:
// 0-10 = corridor positions
// 100-103 = first chamber
// 200-203 = second chamber
// 300-303 = third chamber
// 400-403 = third chamber
type Move struct {
	piece     byte // the piece letter
	src, dest int  // source and destination position
	steps     int  // number of steps taken
	cost      int  // the cost of this move
}

// Cost per move of the different species
var costPerMove = map[byte]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

// Corridor positions for the four chamber entrances
var corriDoor = map[byte]int{'A': 2, 'B': 4, 'C': 6, 'D': 8}

// The position indicators for the first position in each chamber
var chamberPosition = map[byte]int{'A': 100, 'B': 200, 'C': 300, 'D': 400}

// The chamber number for each specieies
var speciesChamber = map[byte]int{'A': 0, 'B': 1, 'C': 2, 'D': 3}

// The best solution found during the simulations
var bestSolution int // initially zero

// Make a deep copy of a board
func copyBoard(b0 Board) Board {

	// Create new board and copy chambers and corridor
	b1 := Board{}
	for c := 0; c < 4; c++ {
		for i := 0; i < cdepth; i++ {
			b1.chambers[c][i] = b0.chambers[c][i]
		}
	}
	for i := 0; i < 11; i++ {
		b1.corridor[i] = b0.corridor[i]
	}

	// Copy the moves
	for _, m := range b0.moves {
		m1 := Move{m.piece, m.src, m.dest, m.steps, m.cost}
		b1.moves = append(b1.moves, m1)
	}

	return b1
}

// Search the board for all pieces, and find all the possible moves
// for each of those pieces; returns nil if none left (i.e., all pieces
// in their destination position, so the game is over)
func possibleMoves(b Board) []Move {

	// List of moves to be created
	moves := []Move{}

	// Check the corridor first: can any pieces be moved
	// to destination chamber? (cannot move from corridor to another
	// position on the corridor, and there is no point moving from the
	// corridor to a chamber other than one's own)
	for pos0 := 0; pos0 < len(b.corridor); pos0++ {

		// Skip if nothing in this cell
		if b.corridor[pos0] == 0 {
			continue
		}

		// Get the position of the door for this letter
		species := b.corridor[pos0] // the creature we are going to move
		door := corriDoor[species]  //  corridor location of target door, e.g., 2 for A

		// Try walking toward the target door, fail if any obstacles
		steps := 0
		step := 1 // 1 for right, -1 for left
		if door < pos0 {
			step = -1
		}
		obstacle := false
		pos1 := pos0
		for {
			pos1 += step
			steps++
			if b.corridor[pos1] != 0 {
				obstacle = true
				break
			} else if pos1 == door {
				break
			}
		}
		if obstacle {
			continue
		}

		// We are at the mouth of the destination door, is it empty
		// except for others of our species?
		chamber := speciesChamber[species] // Chamber number, i.e., 0 for A, 1 for B, ...
		if b.chambers[chamber][0] != 0 {   // chamber is full
			continue
		}
		pos1 = chamberPosition[species] - 1 // e.g., 99 for the first position in chamber A
		ok := true
		for i := 0; i < cdepth; i++ {
			occupant := b.chambers[chamber][i]
			if occupant == 0 { // Empty position, move to it
				steps++
				pos1++
			} else if occupant != species {
				ok = false // a different species, can't move on top of it
				break
			}
		}

		// We have found a feasible move, add it to the list
		if ok {
			cost := steps * costPerMove[species]
			m := Move{piece: species, src: pos0, dest: pos1, steps: steps, cost: cost}
			moves = append(moves, m)
		}
	}

	// Now check each of the four chambers:
	// - We are only interested in the creature closest to the entrance
	// - If that creature is in its destination chamber, and any creatures below
	//   it are also of its own species, not valid to move from here
	// - Otherwise, the creature may move into any position in the corridor
	//   that is free, and not opposite a door
	for c := 0; c < 4; c++ { //  each chamber

		// Find the creature closest to the entrance of this chamber
		pos0 := -1
		for i := 0; i < cdepth; i++ {
			if b.chambers[c][i] != 0 { // found a non-empty cell
				pos0 = i
				break
			}
		}

		// If chamber is empty, move to the next one
		if pos0 == -1 {
			continue
		}

		// If the creature is in its destination chamber, and it's the only
		// one there, or all the creatures below if are also the same species,
		// don't move it
		species := b.chambers[c][pos0]
		isHome := (c == 0 && species == 'A') || (c == 1 && species == 'B') || (c == 2 && species == 'C') || (c == 3 && species == 'D')
		if isHome {
			for i := pos0; i < cdepth; i++ {
				if b.chambers[c][i] != species {
					isHome = false
					break
				}
			}
		}
		if isHome {
			continue
		}

		// Now we have found a creature that can move, so create possible moves,
		// to every corridor cell it can reach without being blocked
		steps := pos0 + 2                    // i.e., one step to get into corridor from location inside door
		corriPos := c*2 + 2                  // corridor position of the door
		for x := corriPos - 1; x >= 0; x-- { // move left
			if b.corridor[x] != 0 { // obstacle, can't move further
				break
			}
			if !(x == 2 || x == 4 || x == 6 || x == 8) { // not opp door
				cost := steps * costPerMove[species]
				m := Move{piece: species, src: (c+1)*100 + pos0, dest: x, steps: steps, cost: cost}
				moves = append(moves, m)
			}
			steps += 1
		}
		steps = pos0 + 2
		for x := corriPos + 1; x < len(b.corridor); x++ { // move right
			if b.corridor[x] != 0 { // obstacle, can't move further
				break
			}
			if !(x == 2 || x == 4 || x == 6 || x == 8) { // not opp door
				cost := steps * costPerMove[species]
				m := Move{piece: species, src: (c+1)*100 + pos0, dest: x, steps: steps, cost: cost}
				moves = append(moves, m)
			}
			steps += 1
		}

	}

	// Return moves
	return moves
}

func evaluate(b Board) {

	// Check integrity of the board (panics if there is a problem)
	checkBoard(b)

	// Add up total cost from moves
	tc := 0
	for _, m := range b.moves {
		tc += m.cost
	}

	// Show the board
	if verbose {
		fmt.Println("\nEvaluating board:", b.chambers, b.corridor, tc)
		printBoard(b)
	}

	// If board is in final state, we are done!
	if finalState(b) {
		if verbose {
			fmt.Printf("SOLUTION FOUND, COST = %d (best %d)\n", tc, bestSolution)
		}
		if bestSolution == 0 || tc < bestSolution {
			fmt.Printf("NEW BEST SOLUTION, COST = %d\n", tc)
			bestSolution = tc
			if verbose {
				fmt.Println("Summary of moves:")
				for _, m := range b.moves {
					printMove(m)
				}
			}
		}
		return
	}

	// If the total cost exceeds the best solution found so far, don't proceed
	if bestSolution > 0 && tc > bestSolution {
		if verbose {
			fmt.Printf("Cost %d exceeds best solution so far %d, stopping branch\n", tc, bestSolution)
		}
		return
	}

	// Get all possible moves for this board, finish if none left
	moves := possibleMoves(b)
	if len(moves) == 0 {
		if verbose {
			fmt.Println("No more moves possible from this state, total cost =", tc)
		}
		return
	}

	// For each  possible move, create a copy of the current state,
	// execute the move on it, and recursively evaluate the resulting state for
	// possible moves
	for _, m := range moves {
		if verbose {
			//fmt.Printf("%d possible moves from this state\n", len(moves))
			fmt.Print("Trying: ")
			printMove(m)
		}
		b1 := copyBoard(b)
		b1.moves = append(b1.moves, m)
		doMove(&b1, m)
		evaluate(b1)
	}
}

// 0-10 = corridor positions
// 100-103 = first chamber
// 200-203 = second chamber
// 300-303 = third chamber
// 400-403 = third chamber
func doMove(b *Board, m Move) {

	// If src is in corridor (src = 0..10)
	if m.src <= 10 {
		creature := b.corridor[m.src]
		if m.dest >= 400 { // chamber D
			b.chambers[3][m.dest-400] = creature
		} else if m.dest >= 300 { // chamber C
			b.chambers[2][m.dest-300] = creature
		} else if m.dest >= 200 { // chamber B
			b.chambers[1][m.dest-200] = creature
		} else { // chamber A
			b.chambers[0][m.dest-100] = creature
		}
		b.corridor[m.src] = 0
	} else { // move from tunnel
		var creature byte = 0
		if m.src >= 400 { // chamber D
			creature = b.chambers[3][m.src-400]
			b.chambers[3][m.src-400] = 0
		} else if m.src >= 300 { // chamber C
			creature = b.chambers[2][m.src-300]
			b.chambers[2][m.src-300] = 0
		} else if m.src >= 200 { // chamber B
			creature = b.chambers[1][m.src-200]
			b.chambers[1][m.src-200] = 0
		} else { // chamber A
			creature = b.chambers[0][m.src-100]
			b.chambers[0][m.src-100] = 0
		}
		b.corridor[m.dest] = creature
	}
}

// Print one move, for debugging
func printMove(m Move) {
	fmt.Printf("Move %s from %d -> %d (%d steps, cost %d)\n",
		string(m.piece), m.src, m.dest, m.steps, m.cost)
}

// Determine whether the board is in its final state, i.e.,
// all the pieces are in their correct chambers
func finalState(b Board) bool {

	// Corridor should be empty
	for i := 0; i < len(b.corridor); i++ {
		if b.corridor[i] != 0 {
			return false
		}
	}

	// Each chamber should have 4 of its species
	for ci := 0; ci < 4; ci++ {
		expected := byte('A' + ci)
		for i := 0; i < cdepth; i++ {
			if b.chambers[ci][i] != expected {
				return false
			}
		}
	}

	// Otherwise full solution found
	return true
}

// Print a representation of the board (for debugging)
func printBoard(b Board) {

	// Print the corridor
	for i := 0; i < len(b.corridor); i++ {
		c := b.corridor[i]
		if c == 0 {
			c = '.'
		}
		fmt.Print(string(c))
	}
	fmt.Print("\n")

	// Print each row
	for i := 0; i < cdepth; i++ {
		fmt.Print("  ")
		for ci := 0; ci < 4; ci++ {
			c := b.chambers[ci][i]
			if c == 0 {
				c = '.'
			}
			fmt.Print(string(c), " ")
		}
		fmt.Print("\n")
	}

}

// Check integrity of board, by counting the number of pieces
func checkBoard(b Board) {

	npieces := map[byte]int{}

	for _, c := range b.corridor {
		npieces[c]++
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < cdepth; j++ {
			npieces[b.chambers[i][j]]++
		}
	}
	if npieces['A'] != cdepth || npieces['B'] != cdepth || npieces['C'] != cdepth || npieces['D'] != cdepth {
		panic("Board is corrupted")
	}
}

// Depth of each chamber (set this to 2 for Part 1, 4 for Part 2)
const cdepth = 4

// Create and return the initial state, just need to populate
// the four chambers (leave corridor, moves empty).
func initialState() Board {
	b := Board{}

	// Part 1 sample input (2 deep), solution 12521
	/*b.chambers[0] = Chamber{'B', 'A'}
	b.chambers[1] = Chamber{'C', 'D'}
	b.chambers[2] = Chamber{'B', 'C'}
	b.chambers[3] = Chamber{'D', 'A'}*/

	// Part 1 problem input (2 deep), solution 14546
	/*b.chambers[0] = Chamber{'D', 'C'}
	b.chambers[1] = Chamber{'A', 'A'}
	b.chambers[2] = Chamber{'D', 'B'}
	b.chambers[3] = Chamber{'C', 'B'}*/

	// Part 2 problem input (4 deep), solution 42308
	b.chambers[0] = Chamber{'D', 'D', 'D', 'C'}
	b.chambers[1] = Chamber{'A', 'C', 'B', 'A'}
	b.chambers[2] = Chamber{'D', 'B', 'A', 'B'}
	b.chambers[3] = Chamber{'C', 'A', 'C', 'B'}

	return b
}

func main() {

	// Get initial state and evaluate it recursively
	b0 := initialState()
	evaluate(b0)

	// Show best solution achieved
	if bestSolution == 0 {
		fmt.Println("No solution was found")
	} else {
		fmt.Println("Best solution found had cost of", bestSolution)
	}
}
