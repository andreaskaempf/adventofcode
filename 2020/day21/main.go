// Advent of Code 2020, Day 21
//
// Read a list of ingredients and associated allergens, and determine which
// ingredients do not produce any allergies (Part 1), and a list of ingredients
// which produce allergies, sorted by allergen (Part 2). Quite easy using set
// operations.
//
// General approach:
//
// Given the following rules (sample in problem statement):
// 1: mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
// 2: trh fvjkl sbzzf mxmxvkd (contains dairy)
// 3: sqjhc fvjkl (contains soy)
// 4: sqjhc mxmxvkd sbzzf (contains fish)
//
// dairy: intersection of 1 and 2 => mxmxvkd
// fish: intersection of 1 and 4 => mxmxvkd sqjhc
// soy: no intersection => sqjhc fvjkl
//
// union of above: mxmxvkd sqjhc fvjkl
//
// all ingredients: mxmxvkd kfcds sqjhc nhms trh fvjkl sbzzf
//
// all ingreds - union => kfcds nhms trh sbzzf (correct answer)
//
// AK, 24/11/2022

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// A rule is a list of ingredients with an associated list of allergens
type Rule struct {
	ingreds, allerg []string
}

func main() {

	// Read input into a list of rules
	//rules := readInput("sample.txt")
	rules := readInput("input.txt")

	// Get the sets of all allergens and all ingredients
	allergens := []string{}
	ingreds := []string{}
	for _, r := range rules {
		allergens = append(allergens, r.allerg...)
		ingreds = append(ingreds, r.ingreds...)
	}
	allergens = unique(allergens)
	ingreds = unique(ingreds)

	// For each allergen, get the intersection of all ingredients
	// for all recipes with that allergen, and build up the union
	// of those ingredients
	union := []string{}
	for _, a := range allergens {

		// Get all ingredient lists that produce this allergen
		recipes := [][]string{}
		for _, r := range rules {
			if isIn(a, r.allerg) {
				recipes = append(recipes, r.ingreds)
			}
		}

		// Get an intersection of all these lists
		common := recipes[0]
		for i := 1; i < len(recipes); i++ {
			common = intersect(common, recipes[i])
		}
		fmt.Println(a, "=>", common) // for part 2

		// Add contents of intersection to the union
		union = append(union, common...)
	}
	union = unique(union)

	// Part 1 answer is the difference between all ingredients and the union
	// For sample.txt, should be kfcds, nhms, sbzzf, or trh
	ans := difference(ingreds, union)
	fmt.Println("Part 1 (should be kfcds, nhms, sbzzf, trh):", ans)

	// Count up the number of times these ingredients appear
	occ := 0 // number of times these ingredients occur
	for _, i := range ans {
		occ += occurences(i, rules)
	}
	fmt.Printf("Part 1: ingredients appear %d times\n", occ)

	// Part 2 is the list of ingredients, sorted by allergen
	// For sample, should be: mxmxvkd,sqjhc,fvjkl.
	// because mxmxvkd contains dairy.
	//         sqjhc contains fish.
	//         fvjkl contains soy.
	fmt.Println("Part 2 (need to manually reduce and sort):", union)

	// Outputs from part 2, reduced and sorted manually:
	//
	// fish => [cskbmx jrmr]
	// shellfish => [tzxcmr jrmr]
	// wheat => [jrmr cjdmk cskbmx fxzh]
	// nuts => [cjdmk cskbmx xlxknk]
	// dairy => [xlxknk jrmr cskbmx]
	// sesame => [jrmr]
	// peanuts => [cskbmx xlxknk bmhn]
	// soy => [fmgxh bmhn]
	//
	// Reduced and sorted:
	// dairy => [xlxknk]
	// fish => [cskbmx ]
	// nuts => [cjdmk]
	// peanuts => [bmhn]
	// sesame => [jrmr]
	// shellfish => [tzxcmr]
	// soy => [fmgxh]
	// wheat => [fxzh]
	//
	// Answer:  xlxknk,cskbmx,cjdmk,bmhn,jrmr,tzxcmr,fmgxh,fxzh

}

// Count the number of occurrences of ingredient in list of rules
func occurences(i string, rules []Rule) int {
	n := 0
	for _, r := range rules {
		if isIn(i, r.ingreds) {
			n++
		}
	}
	return n
}

// Read input file and parse into a list of Rules
func readInput(filename string) []Rule {
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	rules := []Rule{}
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		r := Rule{}
		words := strings.Split(l, " ")
		ingred := true
		for _, w := range words {
			if w == "(contains" {
				ingred = false
			} else if ingred {
				r.ingreds = append(r.ingreds, w)
			} else {
				if w[len(w)-1] == ')' || w[len(w)-1] == ',' {
					w = string(w[:len(w)-1])
				}
				r.allerg = append(r.allerg, w)
			}
		}
		rules = append(rules, r)
	}
	return rules
}

// SET FUNCTIONS

// Common elements between two lists (set intersection)
func intersect(l1, l2 []string) []string {
	res := []string{}
	for _, x := range l1 {
		if isIn(x, l2) {
			res = append(res, x)
		}
	}
	return res
}

// Difference between two sets (l1 must be larger, so l1 - l2)
func difference(l1, l2 []string) []string {
	res := []string{}
	for _, x := range l1 {
		if !isIn(x, l2) {
			res = append(res, x)
		}
	}
	return res
}

// Get unique values (set) from a list of strings
func unique(ss []string) []string {
	res := []string{}
	for _, s := range ss {
		if !isIn(s, res) {
			res = append(res, s)
		}
	}
	return res
}

// Check if string is in a list of strings
func isIn(s string, l []string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

// Remove element from a list (not used here)
func removeElement(x string, l []string) []string {
	l2 := []string{}
	for _, e := range l {
		if e != x {
			l2 = append(l2, e)
		}
	}
	return l2
}
