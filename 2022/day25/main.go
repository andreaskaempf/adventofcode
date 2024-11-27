// Advent of Code 2022, Day 25
//
// Decode a series of numbers which are in base 5 with some special
// characters representing negative numbers. Then, add up the decimal
// equivalents, and return the result encoded back into the special
// representation.
//
// AK, 25 Dec 2022

package main

import (
	"fmt"
)

// Pairs, for testing
type Pair struct {
	SNAFU   string
	decimal int
}

func main() {

	// Uncomment this to run the tests, which were used to figure out
	// the encoding/decoding
	//doTests()

	// Part 1: read the input file, convert each row to decimal and add them
	// up, convert sum back to SNAFU
	// Input: 28115957264952  =>  122-12==0-01=00-0=02
	fname := "sample.txt"
	fname = "input.txt" // uncomment to run on real input
	part1 := 0
	for _, l := range readLines(fname) {
		part1 += convert(l)
	}
	fmt.Println("Part 1 (s/b 4890 => 2=-1=0):", part1, "=>", SNAFU(part1))

}

// Run a series of tests based on the problem input, decoding then
// re-encoding each
func doTests() {

	// Test cases from sample input (SNAFU -> Decimal)
	testCases := []Pair{
		Pair{"20", 10},
		Pair{"2=", 8},
		Pair{"1=-0-2", 1747},
		Pair{"12111", 906},
		Pair{"2=0=", 198},
		Pair{"21", 11},
		Pair{"111", 31},
		Pair{"20012", 1257},
		Pair{"112", 32},
		Pair{"1=-1=", 353},
		Pair{"1-12", 107},
		Pair{"12", 7},
		Pair{"1=", 3},
		Pair{"122", 37},
	}

	// Do test cases from sample
	for _, t := range testCases {

		// Test SNAFU -> Decimal
		d := convert(t.SNAFU)
		fmt.Print(t.SNAFU, "  ->  ", d)
		if d != t.decimal {
			fmt.Println(" *** s/b", t.decimal)
		} else {
			fmt.Println(" ok")
		}

		// Test decimal -> SNAFU
		snafu := SNAFU(t.decimal)
		fmt.Print(t.decimal, "  ->  ", snafu)
		if snafu != t.SNAFU {
			fmt.Println(" *** s/b", t.SNAFU)
		} else {
			fmt.Println(" ok")
		}
		fmt.Println()

	}
}

// Convert SNAFU to decimal:
//  1. Uses powers of five instead of ten. Starting from the right, you have
//     a ones place, a fives place, a twenty-fives place, etc.
//  2. Instead of using digits four through zero, the digits are 2, 1, 0,
//     minus (written -), and double-minus (written =). Minus is worth -1,
//     and double-minus is worth -2."
//     So 4 -> 2,
//     3 -> 1,
//     2 -> 0
//     1 -> -  (worth -1)
//     0 -> =  (worth -2)
func convert(s string) int {

	translate := map[byte]int{'-': -1, '=': -2} // others not required for decoding
	n := 0                                      // result
	for i := 0; i < len(s); i++ {
		c := s[i]
		dig, ok := translate[c]
		if !ok {
			dig = int(c - '0')
		}
		n = n*5 + dig
	}
	return n
}

// Convert decimal number to SNAFU:
// - Places values are 5 instead of 10
// - Convert -1 to '-' and -2 to '=' as in convert, but also
// - But also: 4 -> 2, 3 > 1, 2 -> 0
func SNAFU(d int) string {

	// Convert to base 5 (digits will be in reverse order)
	digits := []int{}
	for {
		rem := d % 5
		digits = append(digits, rem)
		d /= 5
		if d == 0 {
			break
		}
	}

	// Reverse digits, turn into bytes
	res := []byte{}
	for i := len(digits) - 1; i >= 0; i-- {
		res = append(res, byte(digits[i]))
	}

	// Pad with an extra zero at the front
	res = append([]byte{0}, res...)

	// Turn any 3 into '=' and carry 1
	// Turn any 4 into '-' and carry 1
	// Turn any 5 into 0 and carry 1
	for i := len(res) - 1; i > 0; i-- {
		if res[i] == 3 {
			res[i-1]++
			res[i] = '='
		} else if res[i] == 4 {
			res[i-1]++
			res[i] = '-'
		} else if res[i] == 5 {
			res[i-1]++
			res[i] = 0
		}
	}

	// Remove leading zero
	if res[0] == 0 {
		res = res[1:]
	}

	// Convert numbers to characters, then return string
	for i := 0; i < len(res); i++ {
		if res[i] != '=' && res[i] != '-' {
			res[i] += '0'
		}
	}
	return string(res)
}
