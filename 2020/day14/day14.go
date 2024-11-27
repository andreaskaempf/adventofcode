// Advent of Code 2020, day 14
//
// Read a "program" consisting of binary masks and instructions to set memory
// at given address to a value. For part 1, apply the mask to the value. For
// Part 2, first apply the mask to the address, then expand the address to all
// possible permutations where 'X' are changed to 1 and 0. For both parts,
// sum up the values in memory to get the answer.
//
// AK, 15/10/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	var mask string           // Current value of mask
	mem1 := map[int64]int64{} // Part 1: current number at each location
	mem2 := map[int64]int64{} // Same for Part 2

	// Read input file
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	// Go through line by line
	for _, l := range lines {

		// Parse lines with mask, set current mask
		if strings.HasPrefix(l, "mask = ") {
			mask = l[7:]
			continue
		}

		// Other lines must set memory, with address and value
		if !strings.HasPrefix(l, "mem[") {
			fmt.Println("Skipping invalid line:", l)
			continue
		}
		addr := atoi(l[4:strings.Index(l, "]")]) // address between brackets
		val := atoi(l[strings.Index(l, "=")+2:]) // value after equal sign

		// Part 1: apply mask to the current number, and set memory location
		n := applyMaskToNum(mask, val)
		mem1[addr] = n

		// Part 2: apply mask to address (using different rules than Part 1),
		// then expand address so all possible 1/0 values of 'X' digits,
		// and set memory location (unchanged)
		addrMasked := applyMaskToAddr(mask, addr)
		addrs := expandMask(addrMasked)
		for _, a := range addrs {
			a1 := btoi(a)
			mem2[a1] = val
		}
	}

	// Part 1: Sum up contents of memory
	var tot int64
	for _, v := range mem1 {
		tot += v
	}
	fmt.Println("Part 1:", tot)

	// Part 2: Sum up contents of memory
	var tot2 int64
	for _, v := range mem2 {
		tot2 += v
	}
	fmt.Println("Part 2:", tot2)
}

// Part 1: Apply binary mask to a number, settings 1/0 according to mask, and
// leaving X digits unchanged
func applyMaskToNum(mask string, n int64) int64 {

	// Convert the number to a binary string, with leading zeroes
	b := fmt.Sprintf("%036b", n) // pad leading zeros to 36 length
	bits := []byte(b)
	if len(b) != len(mask) {
		fmt.Println("Mask and binary not same size:", len(b), len(mask))
		return 0
	}

	// Go through the mask, and change any 1/0 digits
	for i := 0; i < len(mask); i++ {
		c := mask[i] // this is a byte, not a rune
		if c != 'X' {
			bits[i] = c
		}
	}

	// Convert binary digits back to decimal
	return btoi(string(bits))
}

// Part 2: Apply binary mask to an address number, settings any X in the mask
// to X in the address, leaving 1/0 unchanged; return new address as binary mask
func applyMaskToAddr(mask string, a int64) string {

	// Convert the address to a binary string, with leading zeroes
	b := fmt.Sprintf("%036b", a) // pad leading zeros to 36 length
	bits := []byte(b)
	if len(b) != len(mask) {
		panic("Mask and binary not same size:")
	}

	// Go through the address, and change any X from mask into X in the address
	for i := 0; i < len(mask); i++ {
		c := mask[i] // this is a byte, not a rune
		if c == '1' {
			bits[i] = '1'
		} else if c == 'X' {
			bits[i] = 'X'
		}
	}

	// Return the address as a pattern of 1/0/X
	return string(bits)
}

// For part 2, expand mask to include all possible values of X digits,
// returns a list of new masks with X digits replaced by 1 and 0
func expandMask(mask string) []string {

	// Start out with list containing only the starting mask
	masks := []string{mask}

	// Go through current set of masks, find the first X in each, and
	// create two new masks from it
	for i := 0; i < len(mask); i++ { // each digit
		masks2 := []string{}
		for _, m := range masks {

			// If mask does not have an X in that location, add it back to list unchanged
			if m[i] != 'X' {
				masks2 = append(masks2, m)
				continue
			}

			// Otherwise add two copies of the mask to the list, one with a
			// 1 in the X position, the other with 0
			m0 := []byte(m)
			m0[i] = '0'
			m1 := []byte(m)
			m1[i] = '1'
			masks2 = append(masks2, string(m0), string(m1))
		}

		// Update the main list of masks
		masks = masks2
	}

	// Return expanded list of masks
	return masks
}

// Parse decimal integer string
func atoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

// Parse binary integer string
func btoi(s string) int64 {
	i, _ := strconv.ParseInt(s, 2, 64)
	return i
}
