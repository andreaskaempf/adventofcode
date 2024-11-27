// These are unit tests for Day 13

package main

import (
	"fmt"
	"testing"
)

// Part 2 test examples, from problem definition
func TestPart2(t *testing.T) {

	// Example lists of numbers
	examples := [][]int64{
		[]int64{7, 19},
		[]int64{5, 7, 11},
		[]int64{7, 13, -1, -1, 59, -1, 31, 19},
		[]int64{17, -1, 13, 19},
		[]int64{67, 7, 59, 61},
		[]int64{67, -1, 7, 59, 61},
		[]int64{67, 7, -1, 59, 61},
		[]int64{1789, 37, 47, 1889},
	}

	// Expected answers
	ans := []int64{56, 20, 1068781, 3417, 754018, 779210, 1261476, 1202161486}

	// Test each example
	for i := 0; i < len(examples); i++ {
		ex := examples[i]
		sb := ans[i]
		res := part2(ex)
		if res != sb {
			t.Error("Error processing Part 2")
			fmt.Println("Input =", ex)
			fmt.Printf("Expected %d, got %d\n", sb, res)
		}
	}

}
