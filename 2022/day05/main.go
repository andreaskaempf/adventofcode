// Advent of Code 2022, Day 05
//
// Simulate execution of instructions to move crates from one tower
// to another, one at a time (Part 1), or in groups (Part 2)
//
// AK, 5 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read the input file
	//filename := "sample.txt"  // uncomment filename required
	filename := "input.txt"
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines), "lines read")

	// Stacks from the problem example, and from input data
	// (transposed manually)

	// Problem example:
	//stack := [][]byte{[]byte("ZN"), []byte("MCD"), []byte("P")}

	// From problem input (uncomment to use this)
	stack := [][]byte{[]byte("WDGBHRV"), []byte("JNGCRF"), []byte("LSFHDNJ"),
		[]byte("JDSV"), []byte("SHDRQWNV"), []byte("PGHCM"), []byte("FJBGLZHC"),
		[]byte("SJR"), []byte("LGSRBNVM")}

	// Process each line: skip over stacks of crates, then follow instructions
	// to move each item from top of one stack to top of another
	crates := true // file starts with crates
	for _, l := range lines {

		// Skip over crates, until after first blank line
		if len(l) == 0 {
			crates = false
			continue
		}
		if crates {
			continue
		}

		// Parse an instruction, e.g., "move 1 from 2 to 1"
		words := strings.Split(l, " ")
		q := atoi(words[1])       // qty to move
		src := atoi(words[3]) - 1 // source tower (adjust for zero indexing)
		dst := atoi(words[5]) - 1 // destination

		// Execute instruction, moving crates one at a time for Part 1
		/*for i := 0; i < q; i++ {
			o := stack[src][len(stack[src])-1]          // object to move
			stack[dst] = append(stack[dst], o)          // add to destination
			stack[src] = stack[src][:len(stack[src])-1] // remove from source
		}*/

		// For part 2, move multiple crates at once instead of one at a time
		oo := stack[src][len(stack[src])-q:]        // object(s) to move
		stack[dst] = append(stack[dst], oo...)      // add to destination
		stack[src] = stack[src][:len(stack[src])-q] // remove from source

	}

	// For answer, show top of each stack in ending config
	fmt.Print("Top of stacks: ")
	for i := 0; i < len(stack); i++ {
		fmt.Print(string(stack[i][len(stack[i])-1]))
	}
}

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}
