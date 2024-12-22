// Advent of Code 2024, Day 22
//
// Use an arcane series of calculations to calculate the next 2000 "secret
// numbers" starting with first for each of about 1600 players. For Part 1, add
// up the 2000th numbers. For Part 2, find optimal revenue that can be
// achieved, by deriving price from last digit of each secret number, the delta
// from each subsequent pair, and finding a sequence of four price changes
// common to all players, such that the revenue from that player is the price
// at the end of the first occurrence of the sequence. Used brute force, runs
// in about 5 minutes in Go.
//
// AK, 22 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Read the input file into a list of numbers
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := strings.Split(string(data), "\n")
	nums := []int64{}
	for _, r := range rows {
		n, _ := strconv.ParseInt(r, 10, 64)
		nums = append(nums, n)
	}
	fmt.Println(len(nums), "numbers")

	// Part 1: process each number from input list, get each one's
	// 2000ths secret number and add them up; also build up lists of the
	// secret numbers for Part 2
	var ans int64
	secretNums := [][]int64{}
	for _, n := range nums { // each secret number
		nn := []int64{n}
		for i := 0; i < 2000; i++ {
			n = nextSecret(n)
			nn = append(nn, n)
		}
		ans += n // Part 1
		secretNums = append(secretNums, nn)
	}
	fmt.Println("Part 1 =", ans) // 13753970725

	// For Part 2, first change each list of secret numbers into just
	// the last digit, then the list of changes in these
	prices := [][]int64{}
	changes := [][]int64{}
	for _, snums := range secretNums { // each list of secret numbers
		pp := []int64{} // list of prices
		cc := []int64{} // list of changes
		for i := 0; i < len(snums); i++ {
			price := lastDigit(snums[i])
			if i > 0 {
				chg := price - pp[len(pp)-1]
				cc = append(cc, chg)
			}
			pp = append(pp, price)
		}
		prices = append(prices, pp)
		changes = append(changes, cc)
	}

	// Find all unique sequences of 4 changes
	// TODO: this would be a lot faster using a map
	fmt.Println("Getting sequences")
	seqs := [][]int64{}
	for _, cc := range changes {
		for i := 0; i < len(cc)-4; i++ {
			seq := cc[i : i+4]
			alreadyHave := false
			for _, s := range seqs {
				if same(s, seq) {
					alreadyHave = true
					break
				}
			}
			if !alreadyHave {
				seqs = append(seqs, seq)
			}
		}
	}
	fmt.Println(len(seqs), "unique sequences")

	// Find the price change sequence that generates the most revenue. Do this
	// by checking each sequence, and the price for each player at the end of
	// the first occurance of that sequence, and add these up.
	fmt.Println("Checking sequences")
	var best int64
	for _, seq := range seqs {
		b := revenue(seq, prices, changes)
		if b > best {
			best = b
		}
	}
	fmt.Println("Part 2 =", best) // 1570
}

// Determine revenue for a given sequence of four price changes. This is done
// by finding the first occurrence of the pattern of four prices changes for
// each player, and taking the price in the last position of the sequence
// found, then adding these up.
func revenue(seq []int64, prices, changes [][]int64) int64 {
	var rev int64
	for x, chgs := range changes {
		i := findSeq(seq, chgs)
		if i > 0 {
			rev += prices[x][i+4] // price at last position in sequence
		}
	}
	return rev
}

// Find first occurrence of sequence in a list of numbers,
// return index or -1
func findSeq(seq, nums []int64) int64 {
	for i := 0; i < len(nums)-len(seq); i++ {
		found := true
		for j := 0; j < len(seq); j++ {
			if nums[i+j] != seq[j] {
				found = false
				break
			}
		}
		if found {
			return int64(i)
		}
	}
	return -1
}

// Take the last digit of a number
func lastDigit(n int64) int64 {
	s := fmt.Sprintf("%d", n)
	lc := s[len(s)-1]
	return int64(lc - '0')
}

// Get the next "secret number" after this one
func nextSecret(s int64) int64 {
	s = prune(mix(s, s*64))
	s = prune(mix(s, s/32))
	s = prune(mix(s, s*2048))
	return s
}

// For calculating secret numbers, see problem definition
func mix(secret, n int64) int64 {
	return n ^ secret
}

// For calculating secret numbers, see problem definition
func prune(n int64) int64 {
	return n % 16777216
}

// Check if two slices are the same
func same(s1, s2 []int64) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
