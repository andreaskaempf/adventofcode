// Advent of Code 2022, Day 16
//
// Given a network (graph) of closed "valves", each with a certain flow rate,
// connected by "tunnels", find the sequence of opening the valves (takes
// one minute, plus one minute per step to get there) that yields the highest
// possible total flow during a 30-minute period. For Part 2, same but try two
// decisions (one for you and one for the "elephant") each time step, over
// 26 minutes.
//
// This is an optimization problem. Used simple depth-first dynamic programming
// solution, recursively tries each feasible candidate unopened valve
// (maintained in a list that keeps getting smaller each recursive call)
// excluding those for which we wouldn't have enough time to get any flow. Same
// for Part 2, but tried all possible pairs of remaining valves, one for each
// actor (slow but works).
//
// AK, 16 and 26 Dec 2022

package main

import (
	"fmt"
	"strings"

	"github.com/yourbasic/graph"
)

// A node in the graph
type Node struct {
	id     string
	flow   int
	connTo []string
}

// Global list of nodes
var nodes []Node

// Index of node AA
var nodeAA int

// Minutes for the simulation: 30 for Part 1, 26 for Part 2
var minutes int = 30

// Global pointer to the graph
var g *graph.Mutable

// For memoization of graph shortest distances
type Pair struct {
	a, b int
}

var distances map[Pair]int

func main() {

	// Initialize dictionary for memoization of distances
	distances = map[Pair]int{}

	// Read the input file into a graph
	fname := "sample.txt"
	fname = "input.txt"
	readInput(fname)

	// Part 1: optimize total flow released over 26 minutes, for only
	// one actor
	valves := valvesWithFlow() //  needed for both parts
	fmt.Println("Part 1 (s/b 1651, 1647):", optimize1(nodeAA, valves, 1))

	// Part 2: assume two actors, who can act in parallel opening
	// valves, over 26 minutes instead of 30
	minutes = 26
	fmt.Println("Part 2 (s/b 1707, 2169):", optimize2(nodeAA, nodeAA, valves, 1, 1))
}

// Part 1: one recursive iteration of the optimization: given that
// you are at node "here" at time "t", try to open valves in list
// "candidates" and find the highest cumulative flow.
func optimize1(here int, candidates []int, t int) int {

	// Try each candidate, pruning the ones that would take too
	// long to reach
	best := 0
	for _, vi := range candidates {

		// Don't bother if not enough time to get there, i.e.,
		// by the time you got there, you could not open the valve
		// in time to get any flow
		dist := shortest(here, vi) // time to get to the next valve
		if (minutes-t)-dist < 1 {
			continue
		}

		// Get the value of opening this candidate valve now, until the
		// end of the simulation (takes one time step to open)
		thisValveFlow := nodes[vi].flow * (minutes - t - dist)

		// Recursively simulate moving from this node to the alternative,
		// and optimizing from there this node
		newCand := remove(candidates, vi) // makes a copy
		o := optimize1(vi, newCand, t+dist+1)

		// Is this the best found?
		if thisValveFlow+o > best {
			best = thisValveFlow + o
		}
	}

	return best
}

// Optimization for part 2: assume two actors who can potentially
// open valves at each step, by trying each possible combination
// of remaining open valves at each time step
func optimize2(here1, here2 int, candidates []int, t1, t2 int) int {

	// Stop if no more time or valves left
	if t1 > minutes || t2 > minutes || len(candidates) == 0 {
		return 0
	}

	// Each actor can visit one candidate each time step, so get
	// a list of all possible combinations
	// E.g., [1,2,3] => [1,2], [2,1], [1,3], [3,1], [2,3], [3,2]
	candPairs := [][]int{}
	for _, vi1 := range candidates {
		if here1 < 0 {
			vi1 = -1
		}
		for _, vi2 := range candidates {
			if here2 < 0 {
				vi2 = -1
			}
			if vi1 != vi2 {
				candPairs = append(candPairs, []int{vi1, vi2})
			}
		}
	}

	// For each candidate valve pair, try each combination of you
	// or the elephant doing it
	best := 0
	for _, pair := range candPairs {

		// The candidate valves for you and the elephant
		vi1 := pair[0]
		vi2 := pair[1]

		// Don't bother if not enough time to get there
		// (we set valve ID to -1 to indicate don't go there)
		var dist1, dist2 int
		if here1 >= 0 && vi1 >= 0 {
			dist1 = shortest(here1, vi1) // time to get to the next valve
			if (minutes-t1)-dist1 < 1 {
				vi1 = -1
			}
		}
		if here2 >= 0 && vi2 >= 0 {
			dist2 = shortest(here2, vi2)
			if (minutes-t2)-dist2 < 1 {
				vi2 = -1
			}
		}
		if vi1 < 0 && vi2 < 0 {
			continue
		}

		// Get the value of opening these two candidate valves now, until the
		// end of the simulation (takes one time step to open)
		thisValveFlow := 0
		if here1 >= 0 && vi1 >= 0 {
			thisValveFlow += nodes[vi1].flow * (minutes - t1 - dist1)
		}
		if here2 >= 0 && vi2 >= 0 {
			thisValveFlow += nodes[vi2].flow * (minutes - t2 - dist2)
		}

		// Recursively simulate moving from this node to the alternative,
		// and optimizing from there this node
		newCand := remove(candidates, vi1) // makes a copy
		newCand = remove(newCand, vi2)
		var o int
		if (vi1 >= 0 || vi2 >= 0) && len(newCand) > 0 {
			o = optimize2(vi1, vi2, newCand, t1+dist1+1, t2+dist2+1)
		}

		// Is this the best found?
		if thisValveFlow+o > best {
			best = thisValveFlow + o
		}
	}

	return best
}

// Get a list of the indices of all valves that have non-zero flow,
// since only these are of interest during the optimization (zero-flow
// valves only add time, they are not destinations)
func valvesWithFlow() []int {
	valves := []int{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].flow > 0 {
			valves = append(valves, i)
		}
	}
	return valves
}

// Shortest distance between two nodes, with memoization
func shortest(a, b int) int {
	pair := Pair{a, b}
	dist, ok := distances[pair]
	if !ok {
		_, d := graph.ShortestPath(g, a, b)
		dist = int(d)
		distances[pair] = dist
	}
	return dist
}

// Parse input, create graph
func readInput(filename string) {

	// Read the file, one line per node
	lines := readLines(filename)
	fmt.Println(len(lines), "lines read")

	// Process each line, build list of nodes
	// Valve CC has flow rate=2; tunnels lead to valves DD, BB
	nodes = []Node{} // initialize global variable
	nodeIndex := map[string]int{}
	for _, l := range lines {
		words := strings.Split(l, " ")
		rate := atoi(words[4][5:(len(words[4]) - 1)])
		n := Node{id: words[1], flow: rate}
		nodeIndex[n.id] = len(nodes)  // index of this node
		for _, c := range words[9:] { // list of connected nodes
			if c[len(c)-1] == ',' { // remove comma
				c = c[:len(c)-1]
			}
			n.connTo = append(n.connTo, strings.TrimSpace(c))
		}
		nodes = append(nodes, n)
	}

	// Find the first node AA (global variable)
	nodeAA = -1 // index of node we're at, start at AA
	for i := 0; i < len(nodes); i++ {
		if nodes[i].id == "AA" {
			nodeAA = i
		}
	}
	if nodeAA < 0 {
		panic("Node AA not found!")
	}

	// Great a graph for the nodes
	g = graph.New(len(nodes))
	for ni := 0; ni < len(nodes); ni++ {
		n := nodes[ni]
		for _, c := range n.connTo {
			ci := nodeIndex[c]
			g.AddBothCost(ni, ci, 1) // add bidirectional connection, weight 1
		}
	}
}

// Remove element from a list, returning new copy
func remove(elems []int, e int) []int {
	res := []int{}
	for _, x := range elems {
		if x != e {
			res = append(res, x)
		}
	}
	return res
}
