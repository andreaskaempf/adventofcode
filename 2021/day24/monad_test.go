// A few unit tests, no longer work because of changes in the
// code for monad3, but served their purpose to ensure that
// the simplified versions maintained the functionality of
// the full algorithm transcription.

package main

import (
	"fmt"
	"testing"
)

// Test similarity of two implementations of monad
func TestMonads(t *testing.T) {

	// Some 14-digit numbers to test
	cases := []string{"13579246899999", "11111111111111",
		"22222222222222", "33333333333333", "44444444444444",
		"55555555555555", "66666666666666", "77777777777777",
		"88888888888888", "99999999999999"}

	// Calculate each using each method, check same z
	for _, s := range cases {
		m1 := monad1(s)
		m2 := monad2(s)
		if m2["z"] != m1["z"] {
			t.Error("Monad1 and Monad2 not same")
			fmt.Println("monad1:", m1)
			fmt.Println("monad2:", m2)
		}

		//z3 := monad3(s, false)
		/*if z3 != m1["z"] {   //  no longer works
			t.Error("Monad2 and Monad3 not same")
			fmt.Println("monad2:", m1)
			fmt.Println("monad3:", z3)
		}*/
	}

	// Two known solutions
	d1 := intToDigits(92915979999498)
	d2 := intToDigits(21611513911181)
	if monad3(d1, false) != 0 {
		t.Error("monad3 failed (1)")
	}
	if monad3(d2, false) != 0 {
		t.Error("monad3 failed (2)")
	}
}

func TestConvert(t *testing.T) {

	if !same(intToDigits(123), []int{1, 2, 3}) {
		t.Error("intToDigits failed")
	}

	if digitsToString([]int{1, 2, 3}) != "123" {
		fmt.Println([]int{1, 2, 3}, " -> ", digitsToString([]int{1, 2, 3}))
		t.Error("digitsToString failed")
	}
}

func TestIncrement(t *testing.T) {
	if !same(increment([]int{1, 2, 3}), []int{1, 2, 4}) {
		t.Error("Increment failed (1)")
	}
	if !same(increment([]int{1, 9, 9}), []int{2, 1, 1}) {
		t.Error("Increment failed (2)")
	}
}

func TestDecrement(t *testing.T) {
	if !same(decrement([]int{1, 2, 3}), []int{1, 2, 2}) {
		fmt.Println([]int{1, 2, 3}, "->", decrement([]int{1, 2, 3}))
		t.Error("Decrement failed (1)")
	}
	if !same(decrement([]int{2, 1, 1}), []int{1, 9, 9}) {
		fmt.Println([]int{2, 1, 1}, "->", decrement([]int{2, 1, 1}))
		t.Error("Decrement failed (2)")
	}
}
