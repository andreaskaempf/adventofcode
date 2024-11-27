// Advent of Code 2020, Day 16
//
// Read a file containing train ticket field names, and data for my ticket
// and a bunch of other tickets. In Part 1, identify and remove tickets that
// are invalid, because they do not match the allowed ranges for any field.
// In Part 2, infer which columns relate to which fields, and report the
// value of "departure" fields for my ticket.
//
// AK, 16/10/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Information about a field
type Field struct {
	Name                   string
	Min1, Max1, Min2, Max2 int   // from input
	PossibleCols           []int // inferred in part 2
	Col                    int   // assigned in part 2
}

// Global variables for fields and tickets
var fields []Field
var tickets [][]int

func main() {

	// Read and parse data
	readData("input.txt")

	// Part 1: Find fields that are invalid, i.e., not within any range,
	// sum them up for part 1, and keep just the good tickets for part 2
	var sumBad int
	goodTix := [][]int{}
	for _, t := range tickets { // each ticket
		validTicket := true
		for _, n := range t { // each number on ticket
			validField := false
			for _, f := range fields {
				if isValueValidForField(n, f) {
					validField = true
					break
				}
			}
			if !validField {
				sumBad += n
				validTicket = false
			}
		}
		if validTicket {
			goodTix = append(goodTix, t)
		}
	}

	fmt.Println("Part 1: Sum of bad fields =", sumBad)
	fmt.Printf("%d of %d tickets left\n", len(goodTix), len(tickets))

	// Just keep the good tickets and do Part 2
	tickets = goodTix
	part2()
}

// Part 2: infer which positional field is which, based on values within range
// I.e., each column could be field, X, Y or Z because all values are within
// range. The multiply the values on "my ticket" for all the columns starting
// with "departure"
func part2() {

	// Look at each field, and determine which columns could apply
	fmt.Println("\nPart 2: determining possible columns for each field")
	for i := 0; i < len(fields); i++ { // each field
		fields[i].PossibleCols = []int{}
		for c := 0; c < len(tickets[0]); c++ { // each column
			ok := true                  // assume valid
			for _, t := range tickets { // Check each ticket
				if !isValueValidForField(t[c], fields[i]) {
					ok = false
					break
				}
			}
			if ok {
				fields[i].PossibleCols = append(fields[i].PossibleCols, c)
			}
		}
		//fmt.Println(fields[i])
	}

	// Now iterate to assign columns to fields, basically using a process of
	// elimination, since there is always a field in the list for which only
	// one column is possible:
	// 1. find a field with only one possible column
	// 2. assign that field to that column
	// 3. remove that column number from all fields
	// 4. repeat until no remaining fields with one possible column
	fmt.Println("Assigning columns to fields")
	for {

		// Find a field that has only one possible column
		onePoss := -1
		for i, f := range fields {
			if len(f.PossibleCols) == 1 {
				onePoss = i
				break
			}
		}

		// If no field with one possible column, we are done
		if onePoss == -1 {
			break
		}

		// Assign this field to the single possible column
		col := fields[onePoss].PossibleCols[0]
		fields[onePoss].Col = col
		fmt.Printf("  %s assigned to column %d\n", fields[onePoss].Name, col)

		// Remove that column number from all the fields (including the one just assigned)
		for i := 0; i < len(fields); i++ {
			fields[i].PossibleCols = removeItem(col, fields[i].PossibleCols)
		}
	}

	// Now that we know the column for each field, multiply the values on
	// my ticket for all columns starting with "departure"
	myTicket := tickets[0] // first ticket in input is mine
	var ans int64 = 1
	for _, f := range fields {
		if strings.HasPrefix(f.Name, "departure") {
			ans *= int64(myTicket[f.Col])
		}
	}
	fmt.Println("Part 2 answer:", ans)
}

// Remove item from a list
func removeItem(n int, lst []int) []int {
	lst2 := []int{}
	for i := 0; i < len(lst); i++ {
		if lst[i] != n {
			lst2 = append(lst2, lst[i])
		}
	}
	return lst2
}

// Read and parse problem data
func readData(filename string) {

	// Read input file into list of strings
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")

	// Parse data to find fields and tickets.
	// Up to first blank lines: valid ranges for different fields
	// E.g., class: 1-3 or 5-7
	// After that, tickets are list of numbers (first one is ours)
	fields = []Field{}
	tickets = [][]int{}
	readingFields := true
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if readingFields {
			if len(l) == 0 { // blank line is end of field ranges
				readingFields = false
			} else {
				parts := strings.Split(l, ":")
				ranges := strings.Split(strings.TrimSpace(parts[1]), " ")
				range1 := strings.Split(ranges[0], "-")
				range2 := strings.Split(ranges[2], "-")
				f := Field{Name: parts[0],
					Min1: atoi(range1[0]), Max1: atoi(range1[1]),
					Min2: atoi(range2[0]), Max2: atoi(range2[1])}
				fields = append(fields, f)

			}
		} else if len(l) > 0 && !strings.Contains(l, ":") {
			t := parseTicket(l)
			tickets = append(tickets, t)
		}
	}
}

// Is a field value valid for given field?
func isValueValidForField(n int, f Field) bool {
	return (n >= f.Min1 && n <= f.Max1) || (n >= f.Min2 && n <= f.Max2)
}

// Convert string to number
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse:", s)
		return 0
	}
	return int(n)
}

// Parse a comma-delimited list of numbers
func parseTicket(s string) []int {
	tix := []int{}
	for _, n := range strings.Split(s, ",") {
		tix = append(tix, atoi(n))
	}
	return tix
}
