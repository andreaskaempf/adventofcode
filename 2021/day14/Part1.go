// Advent of Code, Day 14
//
// Expand a sequence of letters according to rules. This works
// for Part 1 (10 iterations), but runs out of memory for Part 2
// (40 iterations), so had change to a different approach (see
// Python version in day14.py).
//
// AK, 14 Dec 2021

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Each letter is kept inside a structure, along with
// a list of the letters that will be inserted after it
type Letter struct {
	letter      byte
	insertAfter []byte
}

// The current list of letters
var letters []Letter

// Translation rules are read into pattern -> insert dictionary
var rules map[string]byte

// Read the data, populating the initial code, and the list of rules
func readData(filename string) bool {

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		return false
	}

	// Read each line
	scanner := bufio.NewScanner(f)
	line := 0
	rules = map[string]byte{}
	for scanner.Scan() {

		line++
		t := scanner.Text()

		// First line has code
		if line == 1 {
			for i := 0; i < len(t); i++ {
				c := t[i]
				letters = append(letters, Letter{c, nil})
			}
		} else if len(t) > 0 { // Subsequent lines have rules
			from := t[:2]
			to := t[6]
			rules[from] = to
		}
	}

	return true
}

// Process one iteration: go through letters, and at each position,
// insert any letters according to the rules. At the end, rebuild the
// list to incorporate the expansions
func expand() {

	// Go through letters up to one before last, and insert
	// letters after this one according to rules
	for i := 0; i < len(letters)-1; i++ {

		// Find pattern for this letter and the next
		l := letters[i]
		pattern := string(l.letter) + string(letters[i+1].letter)
		ins, ok := rules[pattern]
		if !ok { // no rule
			continue
		}

		// If inserting the same letter, increase count, otherwise
		// add a new letter
		letters[i].insertAfter = append(letters[i].insertAfter, ins)

	}

	// Now recreate the string by inserting letters
	ll := []Letter{}
	for _, l := range letters {
		ll = append(ll, Letter{l.letter, nil})
		for _, b := range l.insertAfter {
			ll = append(ll, Letter{b, nil})
		}
	}
	letters = ll
}

func printLetters() {
	for _, l := range letters {
		fmt.Printf("%s", string(l.letter))
	}
	fmt.Println("")
}

// Count frequency of letters
func countFreq() {

	// Count frequency of each letter
	freq := map[byte]int{}
	for _, l := range letters {
		freq[l.letter] += 1
	}

	// Find highest, lowest, difference
	var max, min int
	for _, v := range freq {
		if max == 0 {
			min = v
			max = v
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Printf("Min %d, max %d, diff = %d\n", min, max, max-min)
}

func main() {

	// Read data
	//f := "sample.txt"
	f := "input.txt"
	if !readData(f) {
		fmt.Println("Count not read data")
		return
	}

	// Do 10 iterations (40 for Part 2, but crashes)
	//printLetters()
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration", i+1)
		expand()
		//printLetters()
	}

	countFreq()
}
