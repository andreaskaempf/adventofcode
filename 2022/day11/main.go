// Advent of Code 2022, Day 11
//
// Simulate transfer of objects between a bunch of monkeys, with worry
// level assigned to each object. Each monkey modifies the worry level
// according to some rules, then passes it to one of two monkeys, depending
// on whether the worry level is divisible by that monkey's "test" number.
// Count up the number of inspections each monkey makes during the simulation.
// The answer is the product of the two highest inspection counts. Trivial
// (if tedious) for 20 iterations in Part 1, but integer values overflow for
// 10,000 iterations in Part 2, unless you apply an adjustment that preserves
// the decision outcomes while keeping the numbers fom getting too large.
//
// AK, 11 Dec 2022

package main

import (
	"fmt"
	"sort"
	"strings"
)

// State of a monkey
type Monkey struct {
	id              int      // number of this monkey (0...)
	items           []int64  // worry levels of items it holds
	operation       []string // old*old, old+n, old*n
	test            int64    // divisible by this number
	ifTrue, ifFalse int      // next monkey to throw to
	inspections     int      // number of inspections made
}

// We need this number to prevent the weights from getting too big
var magic int64

func main() {

	// Set this to true for Part 1, false for Part 2
	part1 := false

	// Read and parse "monkeys" from input file, show starting state
	monkeys := readMonkeys("sample.txt")
	//monkeys := readMonkeys("input.txt")
	fmt.Println("Before simulation:")
	for _, m := range monkeys {
		fmt.Println(m)
	}

	// Do the simulation for 20 or 10k rounds
	niters := ifElse(part1, 20, 10000)
	for round := 1; round <= niters; round++ {

		// Do each monkey
		for mi := 0; mi < len(monkeys); mi++ {

			// Process each item the monkey holds
			m := &monkeys[mi]
			for len(m.items) > 0 {

				// Apply operation to the worry level
				m.inspections++
				wl := m.items[0] // current worry level
				wl = applyOperation(wl, m.operation)

				// Now integer-divide by 3 for Part 1
				if part1 {
					wl /= 3
				}

				// For part 2, we also need to adjust the worry level from
				// getting too big, by taking the modulo of it and our "magic"
				// number, which is all the test divisors multiplied together
				// (thanks to my son Alexander for helping me figure this out!)
				wl = wl % magic

				// Apply test to determine who to throw to
				dest := m.ifTrue    // assume divisible by test
				if wl%m.test != 0 { // route to other monkey if not
					dest = m.ifFalse
				}

				// Remove the worry level from this monkey's list, and
				// add it to the destination monkey's list
				m.items = m.items[1:]
				monkeys[dest].items = append(monkeys[dest].items, wl)
			}
		}
	}

	// Get the two highest inspections, answer is the product
	fmt.Println("\nAfter simulation:")
	ii := []int{}
	for _, m := range monkeys {
		fmt.Println(m)
		ii = append(ii, m.inspections)
	}
	sort.Ints(ii)
	fmt.Println("\nFinal number of inspections (s/b 1938, 47830, 52013, 52166):", ii)
	fmt.Println("Answer (on sample, s/b 10605 for Part 1, 2713310158 for Part 2):", ii[len(ii)-1]*ii[len(ii)-2])
}

// Apply an operation to a number: old*old, old+n, old*n
func applyOperation(wl int64, op []string) int64 {
	assert(op[0] == "old" && (op[1] == "+" || op[1] == "*"), "Invalid operation")
	if op[1] == "+" {
		return wl + atoi64(op[2])
	} else if op[2] == "old" {
		return wl * wl // this overflows in Part 2 without adjustment!
	} else {
		return wl * atoi64(op[2])
	}
}

// Read and parse "monkeys" from input file
func readMonkeys(filename string) []Monkey {

	// Create empty list of monkeys, and initial monkey
	monkeys := []Monkey{}
	m := Monkey{id: 0}

	// Initialize global variable of all the test divisors
	// multiplied together, for part 2
	magic = 1

	// Process each line of input file
	lines := readLines(filename)
	for _, l := range lines {

		// Blank line starts a new "monkey"
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			monkeys = append(monkeys, m)
			m = Monkey{id: len(monkeys)} // add current monkey to list
			continue
		}

		// Otherwise fill in fields about the current monkey
		words := strings.Split(l, " ")
		if words[0] == "Starting" { // list of numbers
			for i := 2; i < len(words); i++ {
				n := words[i]
				if n[len(n)-1] == ',' {
					n = n[:len(n)-1]
				}
				m.items = append(m.items, atoi64(n))
			}
		} else if words[0] == "Operation:" {
			m.operation = words[3:]
		} else if words[0] == "Test:" { // e.g., "divisible by 13"
			m.test = atoi64(words[3])
			magic *= m.test
		} else if words[1] == "true:" { // e.g., "throw to monkey 3"
			m.ifTrue = atoi(words[5])
		} else if words[1] == "false:" {
			m.ifFalse = atoi(words[5])
		} else if !strings.HasPrefix(l, "Monkey ") {
			panic("Invalid line: " + l)
		}
	}

	// Add the last monkey and return list
	monkeys = append(monkeys, m)
	return monkeys
}
