// Advent of Code 2021, Day 18
//
// These are the unit tests for Day 18
//
// AK, 18/12/2021

package main

import (
	"fmt"
	"testing"
)

// Test basic expression operations
func TestParse(t *testing.T) {

	// Test parsing a simple expression
	shouldBe1 := []int{LB, 1, 2, RB}
	if !same(parse("[1,2]"), shouldBe1) {
		t.Error("Unable to parse [1,2]")
	}

	// For test cases, we need to be able to parse two-digit numbers
	if !same(parse("[13]"), []int{LB, 13, RB}) {
		t.Error("Unable to parse [13]")
	}
}

// Test addition
func TestAdd(t *testing.T) {

	// Test adding [1,2] + [[3,4],5]
	res2 := add(parse("[1,2]"), parse("[[3,4],5]"))
	shouldBe2 := parse("[[1,2],[[3,4],5]]")
	if !same(res2, shouldBe2) {
		t.Error("Unable to add expressions")
	}
}

// Test "exploding" of pairs in expressions
func TestExplode(t *testing.T) {

	// [[[[[9,8],1],2],3],4] becomes [[[[0,9],2],3],4]
	res, yes := explode(parse("[[[[[9,8],1],2],3],4]"))
	sb := parse("[[[[0,9],2],3],4]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (1)")
	}

	// [7,[6,[5,[4,[3,2]]]]] becomes [7,[6,[5,[7,0]]]]
	res, yes = explode(parse("[7,[6,[5,[4,[3,2]]]]]"))
	sb = parse("[7,[6,[5,[7,0]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (2)")
	}

	// [[6,[5,[4,[3,2]]]],1] becomes [[6,[5,[7,0]]],3]
	res, yes = explode(parse("[[6,[5,[4,[3,2]]]],1]"))
	sb = parse("[[6,[5,[7,0]]],3]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (3)")
	}
	// [[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]
	res, yes = explode(parse("[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"))
	sb = parse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (4)")
	}
	// [[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]] becomes [[3,[2,[8,0]]],[9,[5,[7,0]]]]
	res, yes = explode(parse("[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"))
	sb = parse("[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
	if !(yes && same(res, sb)) {
		t.Error("Unable to explode (5)")
	}
}

func TestSplit(t *testing.T) {

	// 10 becomes [5,5]
	if !same(splitNum(10), parse("[5,5]")) {
		t.Error("Unable to split (1)")
	}

	// 11 becomes [5,6]
	if !same(splitNum(11), parse("[5,6]")) {
		t.Error("Unable to split (2)")
	}

	// 12 becomes [6,6]
	if !same(splitNum(12), parse("[6,6]")) {
		t.Error("Unable to split (3)")
	}

	// [13] becomes [[5,6]]
	res0, yes0 := splitFirst(parse("[13]"))
	if !(yes0 && same(res0, parse("[[6,7]]"))) {
		t.Error("Unable to split first [13]")
		printExpr(parse("[13]"))
		printExpr(res0)
	}

	// [[[[0,7],4],[15,   [0,13]]],[1,1]] becomes
	// [[[[0,7],4],[[7,8],[0,13]]],[1,1]]
	x1 := parse("[[[[0,7],4],[15,[0,13]]],[1,1]]")
	sb1 := parse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")
	res1, yes1 := splitFirst(x1)
	if !(yes1 && same(res1, sb1)) {
		t.Error("Unable to split first (1)")
	}
	// [[[[0,7],4],[[7,8],[0,13]]],   [1,1]] (result from above) becomes
	// [[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]
	x2 := parse("[[[[0,7],4],[[7,8],[0,13]]],[1,1]]")
	sb2 := parse("[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]")
	res2, yes2 := splitFirst(x2)
	if !(yes2 && same(res2, sb2)) {
		t.Error("Unable to split first (2)")
	}
}

// Test reduce
func TestReduce(t *testing.T) {

	// Adding then reducing [[[[4,3],4],4],[7,[[8,4],9]]] + [1,1]
	// gives [[[[0,7],4],[[7,8],[6,0]]],[8,1]]
	x1 := parse("[[[[4,3],4],4],[7,[[8,4],9]]]")
	x2 := parse("[1,1]")
	x3 := add(x1, x2)
	res := reduce(x3)
	sb := parse("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")
	if !same(res, sb) {
		t.Error("Unable to reduce")
		fmt.Print("Result: ")
		printExpr(res)
		fmt.Print("Should be: ")
		printExpr(sb)
	}
}

// Test adding up lines from a file
func TestAddUp(t *testing.T) {

	res1 := addUpFile("sample1.txt")
	sb1 := parse("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")
	if !same(res1, sb1) {
		t.Error("Unable to add up sample1.txt")
	}

	res2 := addUpFile("sample2.txt")
	sb2 := parse("[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]")
	if !same(res2, sb2) {
		t.Error("Unable to add up sample2.txt")
	}
}

// Test magnitude

// To check whether it's the right answer, the snailfish teacher only checks
// the magnitude of the final sum. The magnitude of a pair is 3 times the
// magnitude of its left element plus 2 times the magnitude of its right
// element. The magnitude of a regular number is just that number.
//
// For example, the magnitude of [9,1] is 3*9 + 2*1 = 29;
// the magnitude of [1,9] is 3*1 + 2*9 = 21.
// Magnitude calculations are recursive:
// the magnitude of [[9,1],[1,9]] is 3*29 + 2*21 = 129.
//
// Here are a few more magnitude examples:
//
// [[1,2],[[3,4],5]] becomes 143.
// [[[[0,7],4],[[7,8],[6,0]]],[8,1]] becomes 1384.
// [[[[1,1],[2,2]],[3,3]],[4,4]] becomes 445.
// [[[[3,0],[5,3]],[4,4]],[5,5]] becomes 791.
// [[[[5,0],[7,4]],[5,5]],[6,6]] becomes 1137.
// [[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]] becomes 3488.

func TestMagnitude(t *testing.T) {

	if magnitude(parse("[9,1]")) != 29 {
		t.Error("Unable to calc magnitude (1)")
	}

	if magnitude(parse("[[9,1],[1,9]]")) != 129 {
		t.Error("Unable to calc magnitude (2)")
	}

	if magnitude(parse("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")) != 3488 {
		t.Error("Unable to calc magnitude (3)")
	}
}

// Test magnitude of an input file
func TestFileMag(t *testing.T) {

	expr := addUpFile("sample2.txt")
	mag := magnitude(expr)
	if mag != 4140 {
		t.Error("Unable to calculate magnitude of sample2.txt")
	}
}

// Compare two int arrays element-by-element, and report
// if they are the same
func same(a Expression, b Expression) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
