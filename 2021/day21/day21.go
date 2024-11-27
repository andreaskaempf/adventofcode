// Advent of Code 2021, Day 21
//
// Simulate a game of rolling dice and moving round a board, trivial in
// part 1. In part 2, fork a set of parallel "universes" with identical state
// after every throw of a 3-sided die, and get the number of
// universes in which the winner won.
//
// AK, 21/12/2021

package main

import (
	"fmt"
)

// Starting positions for the two players (puzzle input)
//const start1, start2 int = 4, 8 // sample
const start1, start2 int = 6, 3 // input

func part1() {

	// Current positions and scores of each player
	pos1 := start1
	pos2 := start2
	var score1, score2 int

	// Next value and total rolls of the "deterministic" die
	die := 1
	rolls := 0

	// Play the game
	for {
		// Player 1 rolls dice 3 times, adds up positions
		roll1 := die + (die + 1) + (die + 2)
		die += 3
		rolls += 3
		pos1 += roll1
		for pos1 > 10 {
			pos1 -= 10
		}
		score1 += pos1
		fmt.Printf("Player 1 rolled %d, moves to space %d, score = %d\n", roll1, pos1, score1)
		if score1 >= 1000 {
			fmt.Println("Player 1 wins with ", score1)
			break
		}

		// Player 2
		roll2 := die + (die + 1) + (die + 2)
		die += 3
		rolls += 3
		pos2 += roll2
		for pos2 > 10 {
			pos2 -= 10
		}
		score2 += pos2
		fmt.Printf("Player 2 to space %d, score = %d\n", pos2, score2)
		if score2 >= 1000 {
			fmt.Println("Player 2 wins with ", score2)
			break
		}
	}

	fmt.Printf("Answer: %d rolls, scores %d and %d => %d\n",
		rolls, score1, score2, min(score1, score2)*rolls)
}

// A "universe" (there will be many in a given state)
type Universe struct {
	pos0, pos1     int // players' positions (1-10)
	score0, score1 int // players' scores (0 to 21)
}

// Play game in each universe, with 3-sided die. Each player rolls the
// dice 3 times (1, 2, 3), adds up the values (6) and advances that
// many positions (circling back to position 1 if > 10). The position
// is then added to that player's score.
//
// Rolling the die always splits the universe into three copies:
// one where the outcome of the roll was 1, one where it was 2,
// and one where it was 3.
//
// Find the player that wins in more universes; in how many universes
// does that player win?
func part2() {

	// Start with one universe with two players at their starting positions
	univs := map[Universe]int64{}
	univs[Universe{start1, start2, 0, 0}] = 1

	// Number of wins for each of the two players
	var wins [2]int64

	// All combinations of 1,2,3
	c3 := combinations3()

	// Iterate until done
	iter := 0
	done := false
	player := 0 // for the two players
	for !done {

		// Switch players
		player = 1 - player

		// Play game in each universe, spawn a new game each throw
		univs2 := map[Universe]int64{} // for new universes
		for u, un := range univs {     // each many universes of the same state

			// Get current score and position for the first player, in
			// this universe
			oldScore := u.score0
			oldPos := u.pos0

			// Roll every combination of 3-sided dice 3 times, create
			// new universe each time
			for _, rolls := range c3 {

				// Advance position, circling back to 1 if > 10
				pos := oldPos + rolls[0] + rolls[1] + rolls[2]
				for pos > 10 {
					pos -= 10
				}

				// On the third roll, update score with position, stop if won,
				// otherwise create a set of parallel "universes" with this state
				score := oldScore + pos
				if score >= 21 {
					wins[player] += un
				} else {
					u1 := Universe{u.pos1, pos, u.score1, score}
					univs2[u1] += un // map returns 0 if key missing
				}
			}
		}

		// Reset the main list of universes to the current state
		univs = univs2

		// Stop if both players in all universes have scores of 21 or more
		done = true
		for u, _ := range univs {
			if u.score0 < 21 || u.score0 < 21 {
				done = false
				break
			}
		}

		// Show number of universes after this iteration
		iter++
		nunivs := countUnivs(univs)
		fmt.Printf("Iteration %2d: %20d universes\n", iter, nunivs)
		if nunivs > 10e15 {
			fmt.Println("Too big, aborting")
			break
		}
	}

	// Count up how many universes each player won in
	// Player 1 wins in 444,356,092,776,315 universes in sample,
	// Player 2 wins in 341,960,390,180,808 universes,
	// Total universes =786,316,482,957,123 (5 groups = 15 digits)
	fmt.Printf("Player 1 won in %d universes (%.1fT), Player 2 in %d (%.1fT)\n",
		wins[0], float64(wins[0])/1e12, wins[1], float64(wins[1])/1e12)
}

// Count up the total universes
func countUnivs(univs map[Universe]int64) int64 {
	var nu int64
	for _, n := range univs {
		nu += n
	}
	return nu
}

// Make a list of all combinations of 1,2,3
func combinations3() [][]int {
	comb3 := [][]int{}
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {
				comb3 = append(comb3, []int{d1, d2, d3})
			}
		}
	}
	return (comb3)
}

// Minimum of two ints
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	part1() // sample should be 739785 for part 1
	part2()
}
