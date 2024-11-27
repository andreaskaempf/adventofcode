// Advent of Code 2020, Day XX
//
// AK, x/x/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read both decks of cards
	//data, _ := ioutil.ReadFile("sample.txt")
	data, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(data), "\n")
	var player1, player2 []int
	reading2 := false
	for _, l := range rows {
		if l == "Player 1:" || len(l) == 0 {
			continue
		} else if l == "Player 2:" {
			reading2 = true
		} else if reading2 {
			player2 = append(player2, atoi(l))
		} else {
			player1 = append(player1, atoi(l))
		}
	}

	// Simulate rounds until one player has no cards left
	var round, card1, card2 int
	for len(player1) > 0 && len(player2) > 0 {

		round++
		fmt.Println("Round", round)
		fmt.Println("  Player 1:", player1)
		fmt.Println("  Player 2:", player2)

		// Draw cards
		card1 = player1[0]
		player1 = player1[1:]
		card2 = player2[0]
		player2 = player2[1:]

		// Higher value keeps both cards
		if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
	}

	// Calculate score
	winner := player1
	if len(player2) > 0 {
		winner = player2
	}
	score := 0
	for i := 0; i < len(winner); i++ {
		score += (i + 1) * winner[len(winner)-i-1]
	}
	fmt.Println("Winning score =", score)
}

// Parse number
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse, assuming zero:", s)
		n = 0
	}
	return int(n)
}
