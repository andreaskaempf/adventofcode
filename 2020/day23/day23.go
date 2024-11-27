package main

import (
	"container/ring"
	"fmt"
)

func main() {

	// List of cups
	//cups := []int{3, 8, 9, 1, 2, 5, 4, 6, 7} // sample
	cups := []int{1, 5, 7, 6, 2, 3, 9, 8, 4} // puzzle input

	// Set to true for part 2
	var part2 bool // default false
	part2 = true

	// Create and populate a ring (circular list)
	n := len(cups)
	million := 1000000 //1000000
	if part2 {
		n = million
	}
	r := ring.New(n)
	var maxVal int // the highest cup value
	for i := 0; i < len(cups); i++ {
		v := cups[i]
		if v > maxVal {
			maxVal = v
		}
		r.Value = v
		r = r.Next()
	}

	// For part 2, number the remaining cups sequentially starting with max+1
	// until there are 1 million
	if part2 {
		v := maxVal + 1
		for i := len(cups); i < million; i++ {
			r.Value = v
			if v > maxVal {
				maxVal = v
			}
			v++
			r = r.Next()
		}
	}

	// Do 100 iterations, or 1 million for part 2
	niter := 100
	if part2 {
		niter = million
	}
	for i := 1; i <= niter; i++ {

		fmt.Println("\nIteration", i)
		printRing(r, "Starting:")

		// 1. The crab picks up the three cups that are immediately clockwise
		// of the current cup. They are removed from the circle; cup spacing is
		// adjusted as necessary to maintain the circle.
		removed := r.Unlink(3) // removes the 3 cups AFTER the current cup
		printRing(removed, "Removed:")
		printRing(r, "  now:")

		// 2. The crab selects a destination cup: the cup with a label equal to
		// the current cup's label minus one. If this would select one of the
		// cups that was just picked up, the crab will keep subtracting one
		// until it finds a cup that wasn't just picked up. If at any point in
		// this process the value goes below the lowest value on any cup's
		// label, it wraps around to the highest value on any cup's label
		// instead.

		// Get the next value to search for (current minus one)
		destVal := r.Value.(int) - 1
		if destVal <= 0 {
			destVal = maxVal
		}
		already := destVal // to detect endless loops
		fmt.Println("Destination (current cup value - 1):", destVal)

		// Try to find it in remaining cups (don't bother if it's in removed list)
		var destEl *ring.Ring = nil              // initialized to nil
		if ringSearch(removed, destVal) == nil { // don't bother searching if in removed list
			destEl = ringSearch(r, destVal) // find this cup
		}

		// If not found, try each lower value
		for destEl == nil { // TODO: risk of infinite loop?

			// Get the next value to search for
			destVal--
			if destVal == already {
				fmt.Println("Endless loop detected")
				return
			}
			if destVal <= 0 {
				destVal = maxVal
			}

			if ringSearch(removed, destVal) == nil { // don't bother searching if in removed list
				destEl = ringSearch(r, destVal)
			} else {
				destEl = nil
			}
		}

		// 3. The crab places the cups it just picked up so that they are
		// immediately clockwise of the destination cup. They keep the same
		// order as when they were picked up.
		destEl.Link(removed)
		printRing(r, "After reinserting removals:")

		// 4. The crab selects a new current cup: the cup which is immediately
		// clockwise of the current cup.
		r = r.Next()
		printRing(r, "Ring now:")

	}
	printRing(r, "\nFinal configuration:")

	// Calculate answer, the sequence starting after 1, omitting the 1
	r = ringSearch(r, 1)
	printRing(r, "Answer (omit the first 1):")
	fmt.Println("  part 1 should be 58427369") //58427369 correct for part 1 with input
}

// Print ring, maximum of 20 values
func printRing(r *ring.Ring, label string) {
	fmt.Print(label)
	n := r.Len()
	if n > 20 {
		n = 20
	}
	for i := 0; i < n; i++ {
		fmt.Print(" ", r.Value.(int))
		r = r.Next()
	}
	if r.Len() > 20 {
		fmt.Print(" ...")
	}
	fmt.Println()
}

func ringSearch(r *ring.Ring, value int) *ring.Ring {
	fmt.Print("Searching for ", value)
	for i := 0; i < r.Len(); i++ {
		if r.Value.(int) == value {
			fmt.Println(value, " ... found at position", i, ", delta =", i-value)
			return r
		}
		r = r.Next()
	}
	fmt.Println(" ... not found!")
	return nil
}
