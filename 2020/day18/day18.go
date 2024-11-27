// Advent of Code 2020, Day 18
//
// Parse and evaluate arithmetic expressions with +, -, *, / and parentheses,
// with left-to right evaluation (no operator precedence) for Part 1, and
// mult/div having higher precedence for Part 2. Implemented Djikstra's
// Shunting Yard Algorithm. Part 2 was a trivial change to some precedence
// weights.
//
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
//
// AK, 17/10/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	// Evaluate each equation and add up answers
	// Sample.txt: 71, 51, 26, 437, 12240, 13632
	tot := 0
	data, _ := ioutil.ReadFile("input.txt")
	for _, expr := range strings.Split(string(data), "\n") {
		fmt.Println("Expr:", expr)
		tree := parse(expr)
		val := evaluate(tree)
		fmt.Println(" =", val)
		tot += val
	}

	fmt.Println("Total =", tot)
}

// Parse an expression, return list of tokens in postfix notation
func parse(expr string) []string {

	// Precendence of different operaters, same for part 1
	precedence := map[string]int{"+": 1, "-": 1, "*": 1, "/": 1}

	// Uncomment these line for Part 2, where addition and subtraction have
	// higher precedence
	precedence["+"] = 2
	precedence["-"] = 2

	// Output and operator stacks are just lists
	output := []string{}
	ops := []string{}

	// Tokenize the expression
	tokens := tokenize(expr)

	// Process tokens into stacks of operators and operands
	for _, t := range tokens {
		if isNumber(t) { // number: push onto output stack
			output = append(output, t)
		} else if t == "(" { // left paren: push onto operator stack
			ops = append(ops, t)
		} else if t == ")" { // right paren
			// while the operator at the top of the operator stack is not a left parenthesis:
			//   {assert the operator stack is not empty}
			//   pop the operator from the operator stack into the output queue
			for ops[len(ops)-1] != "(" {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			// {assert there is a left parenthesis at the top of the operator stack}
			// pop the left parenthesis from the operator stack and discard it
			if ops[len(ops)-1] == "(" {
				ops = ops[:len(ops)-1]
			} else {
				panic("Missing ( on top of op stack")
			}
		} else { // assume it's an operator
			for len(ops) > 0 && ops[len(ops)-1] != "(" && precedence[ops[len(ops)-1]] >= precedence[t] {
				output = append(output, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, t)
		}
	}

	// Pop the remaining items from the operator
	// stack into the output queue.
	for len(ops) > 0 {
		output = append(output, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}

	// Return the output queue, which is a list of numbers and operators
	// in postfix order, so they can be evaluated
	//fmt.Println("Ops =", ops)
	//fmt.Println("Output =", output)
	return output
}

// Evaluate a list of postfix notations, i.e., numbers onto stack, then
// operated upon
// TODO: Avoid the converstion back and forth between strings and numbers,
// by implementing token type
func evaluate(tokens []string) int {

	// Implement a simple stack of tokens using a list
	stack := []string{}

	// Process each token
	for _, t := range tokens {

		// Numbers just go on the stack
		if isNumber(t) {
			stack = append(stack, t)
		} else { // apply operator to top of stack

			// Pop x & y off top of stack
			x := atoi(stack[len(stack)-1])
			y := atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]

			// Apply operator, result back on stack
			if t == "+" {
				stack = append(stack, str(x+y))
			} else if t == "-" {
				stack = append(stack, str(x-y))
			} else if t == "*" {
				stack = append(stack, str(x*y))
			} else if t == "/" {
				stack = append(stack, str(x/y))
			}

		}
	}

	// Answer is on top of the stack
	return atoi(stack[len(stack)-1])
}

// Simple tokenizer, combines subsequent digits into numbers, skips spaces,
// and considers any other characters as tokens
func tokenize(expr string) []string {
	tokens := []string{}
	for i := 0; i < len(expr); i++ {
		c := string(expr[i])
		ntoks := len(tokens)
		if isNumber(c) {
			if ntoks > 0 && isNumber(tokens[ntoks-1]) {
				tokens[ntoks-1] = tokens[ntoks-1] + c
			} else {
				tokens = append(tokens, c)
			}
		} else if c != " " {
			tokens = append(tokens, c)
		}
	}
	return tokens
}

// Is a string a number (could be multi-digit)?
func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// Parse number
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse, assuming zero:", s)
		n = 0
	}
	return int(n)
}

// Convert number to string
func str(n int) string {
	s := fmt.Sprintf("%d", n)
	return s
}
