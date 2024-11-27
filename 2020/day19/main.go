// Advent of Code 2020, Day 19
//
// AK, 18/10/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Structure for a rule, consists of either a single character,
// or one or two lists of numbers of other rules that must be matched
type Rule struct {
	num  int   // rule number
	char byte  // matching character (zero if not used)
	L    []int // sub-rule numbers, left part
	R    []int // sub-rule numbers, right part
}

func main() {

	// Read rules and messages from input file
	filename := "sample.txt" // 2 matches (part 1)
	//filename := "sample2.txt" // 2/12 matches (part 1/2)
	// filename := "input.txt" // 156/363
	//readData("sample.txt")
	rules, messages := readData(filename)
	fmt.Println(rules)
	fmt.Println(messages)

	// Test each message against rule zero
	for _, msg := range messages {
		fmt.Println(msg, "=>", match(msg, rules[0], rules))
	}
}

func match(msg string, rule Rule, rules []Rule) bool {

	// If both empty, return true (pattern has been matched)
	if len(msg) == 0 && len(rules) == 0 {
		return true
	}

	// If one is empty but not the other, return false (cannot match)
	if len(msg) == 0 || len(rules) == 0 {
		return false
	}

	// Take the first part of this rule
	r := rules[rule[0]]

	// If it's a letter and matches start of string, check rest, otherwise fail
	if r.char != 0 {
		if msg[0] == r.char {
			return match(s[1:], rule[1:], rules)
		} else {
			return false
		}
	}

	// Check remaining patterns: for each term in the current rule, check
	// it plus the remaining rules
	//for t in r:
	//    if test(s, t + seq[1:]):
	//        return True

	// All failed
	return false
}

// Read data file into a list of rules and message strings
func readData(filename string) (map[int]Rule, []string) {

	// Read file into list of strings
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	// At first blank line, stop reading rules and start reading messages
	rules := map[int]Rule{}
	messages := []string{}
	readingRules := true
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			readingRules = false
			continue
		} else if readingRules {
			r := parseRule(l)
			rules[r.num] = r
		} else {
			messages = append(messages, l)
		}
	}
	return rules, messages
}

// Parse a rule
// E.g., 99: "a"
// or 88: 24 103 | 36 6
func parseRule(s string) Rule {
	words := strings.Split(s, " ")
	n := atoi(words[0][:len(words[0])-1])
	r := Rule{num: n}
	parsingR := false
	for i := 1; i < len(words); i++ {
		w := words[i]
		if w[0] == '"' { // quote means rule is a character
			r.char = w[1]
		} else if words[i] == "|" { // bar means start of right sub-rules
			parsingR = true
		} else if parsingR { // right sub-rules, list of numbers
			r.R = append(r.R, atoi(w))
		} else { // left sub-rules, list of numbers
			r.L = append(r.L, atoi(w))
		}
	}
	return r
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
