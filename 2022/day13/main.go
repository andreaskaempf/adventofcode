// Advent of Code 2022, Day 13
//

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Read input file, split into rows
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("File not found")
	}
	lines := strings.Split(string(data), "\n")

	// Parse each group of lines
	//for _, l := range lines {
	//	parse(l)
	//	break
	//}
    expr := lines[12]
    tokens := tokenize(expr)
	tree := parse(tokens)
    fmt.Println(tree)
}

// An expression can be either a number, or a list of expressions
type Expression struct {
	isNum  bool         // true if number, otherwise list
	number int          // note that numbers can be zero
	list   []Expression // note that lists can be empty
}

// Parse an expression consisting of nested lists and comma-separated numbers,
// e.g., [[6,[],[10],[],1],[10,[0,[],[0,0,6,9,2]]],[6],[[3,[0],7,1],9]]
func parse(tokens []string) Expression {
    
      // Number: return a token with number
    if isNumber
	
    // Parse each token, build up structure
    
}

// Break string into list of tokens
func tokenize(str string) []string {
	var tokens []string
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c == '[' || c == ']' || c == ',' {
			tokens = append(tokens, string(c))
		} else if c >= '0' && c <= '9' {
			n := 0
			for isDigit(c) {
				n = n*10 + int(c-'0')
				i++
				c = str[i]
			}
			i--
			tokens = append(tokens, fmt.Sprintf("%d", n))
		}
	}
	return tokens
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
