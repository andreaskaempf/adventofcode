// Advent of Code 2022, Day 21
//
// Given a list of variable names, each with either a numbers or
// a simple formula, recursively evaluate the root node (Part 1),
// and find the value for one cell that makes the two sides of the
// root node equal (solved using gradient descent).
//
// AK, 21 Dec 2022

package main

import (
	"fmt"
	"strings"
)

// A monkey has a name, and either a simple formula, or a number
type Monkey struct {
	name         string
	lhs, rhs, op string
	num          int64
}

// A dictionary of monkeys, so can find by name
var monkeys map[string]*Monkey

func main() {

	// Read the input file into a dictionary of monkeys
	fname := "sample.txt"
	fname = "input.txt"
	monkeys = map[string]*Monkey{}
	for _, l := range readLines(fname) {
		words := strings.Split(l, " ")
		mname := words[0][:len(words[0])-1]
		m := Monkey{name: mname}
		if len(words) == 2 {
			m.num = atoi64(words[1])
		} else {
			m.lhs = words[1]
			m.op = words[2]
			m.rhs = words[3]
		}
		monkeys[m.name] = &m
	}

	// Part 1: just get the value for "root"
	fmt.Println("Part 1:", process("root"))

	// Part 2: find the value for monkey "humn" that would make the
	// rhs and lhs for "root" equal (use gradient search).
	// 3769668716710 too high, s/b 3769668716709 (subtract one as
	// in sample output)
	lhs := monkeys["root"].lhs
	rhs := monkeys["root"].rhs
	var delta int64 = 1000000000    // change by this much each step
	var guess int64 = 3000000000000 // starting guess
	iters := 0
	for {
		iters++
		monkeys["humn"].num = guess
		diff := process(lhs) - process(rhs)
		//fmt.Println("guess =", guess, ", delta =", delta, ", diff =", diff)
		if diff == 0 {
			fmt.Println("Part 2: found", guess-1, "after", iters, "iterations")
			break
		} else if abs(diff) < delta*10 {
			delta /= 10
		} else if diff > 0 {
			guess += delta
		} else {
			guess -= delta
		}
	}

}

// Determine if this monkey is a number
func isNumber(m *Monkey) bool {
	return len(m.op) == 0
}

// Process (recursively evaluate) a monkey, returning result
func process(name string) int64 {
	m := monkeys[name]
	if isNumber(m) {
		return m.num
	}
	lhs := process(m.lhs)
	rhs := process(m.rhs)
	if m.op == "+" {
		return lhs + rhs
	} else if m.op == "-" {
		return lhs - rhs
	} else if m.op == "*" {
		return lhs * rhs
	} else if m.op == "/" {
		return lhs / rhs
	} else {
		panic("Bad operator")
	}
}
