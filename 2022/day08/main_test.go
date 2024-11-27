// Unit tests for this Advent of Code submission

package main

import (
	"testing"
)

// Test part 1 against problem data set
func TestPart1(t *testing.T) {
	lines := readLines("sample.txt")
	t1 := part1(lines)
	if t1 != 21 {
		t.Errorf("Part 1: got %d instead of 21\n", t1)
	}
}

func TestPart2(t *testing.T) {

	// Test against sample data set, from problem statement
	lines := readLines("sample.txt")
	t1 := scenicScore(1, 2, lines)
	if t1 != 4 {
		t.Errorf("Test 1: score for 1,2 is %d instead of 4\n", t1)
	}
	t2 := scenicScore(3, 2, lines)
	if t2 != 8 {
		t.Errorf("Test 2: score for 3,2 is %d instead of 8\n", t2)
	}
	t3 := part2(lines)
	if t3 != 8 {
		t.Errorf("Test 3: max score on sample data set is %d , should be 8\n", t3)
	}

	// Test main data set for allowed range based on previous attempts
	// 560 and 14400 both too low
	lines = readLines("input.txt")
	t4 := part2(lines)
	if t4 != 671160 {
		t.Errorf("Test 3: max score for main data set is %d  instead of 671160\n", t4)
	}
}
