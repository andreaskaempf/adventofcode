// Advent of Code 2024, Day 23
//
// Given a list of connected node pairs, find all groups of 3 that are
// connected to each other, where at least on of the nodes starts with 't'
// (Part 1).  For Part 2, find the largest set of nodes that are all connected
// to each other, sort names and join with commas. Did this directly, no graph
// library (would not have helped).
//
// AK, 23 Dec 2024

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// All nodes, and direct connections for each node
var nodes []string
var conns map[string][]string

func main() {

	// Read list of connections
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)

	// Create map of connections from each node, both ways
	conns = map[string][]string{}
	for _, l := range strings.Split(string(data), "\n") {
		parts := strings.Split(l, "-")
		n1 := parts[0]
		n2 := parts[1]
		conns[n1] = append(conns[n1], n2)
		conns[n2] = append(conns[n2], n1)
	}

	// Get just the keys, i.e., the node names
	nodes = []string{}
	for n, _ := range conns {
		nodes = append(nodes, n)
	}

	// Part 1: find all trios of nodes connected to each other, where at least
	// one node starts with 't'
	var ans int
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			for k := j + 1; k < len(nodes); k++ {
				n1 := nodes[i]
				n2 := nodes[j]
				n3 := nodes[k]
				if connected(n1, n2) && connected(n1, n3) && connected(n2, n3) &&
					(t(n1) || t(n2) || t(n3)) {
					ans++
				}
			}
		}
	}
	fmt.Println("Part 1 =", ans) // 1077

	// Part 2: find the largest set of nodes that are all connected to each other,
	// sort names and join with commas
	biggest := []string{}
	for _, n := range nodes {
		friends := findFriends(n)
		if len(friends) > len(biggest) {
			biggest = friends
		}
	}
	sort.Strings(biggest)
	fmt.Println("Part 2:", strings.Join(biggest, ",")) // bc,bf,do,dw,dx,ll,ol,qd,sc,ua,xc,yu,zt
}

// Find community directly connected to this node, where each node is also
// connected to every other node in the set
func findFriends(n string) []string {

	// List of known friends, and those that need to be checked
	friends := []string{n} // people known to be connected to everybody else
	todo := conns[n]       // those that still need to be checked
	for len(todo) > 0 {

		// Pop the next person from the todo list
		p := todo[0]
		todo = todo[1:]

		// Reject if not connected to all others already in circle
		connToAll := true
		for _, f := range friends {
			if !connected(p, f) {
				connToAll = false
				break
			}
		}
		if !connToAll {
			continue
		}

		// Add this person to friends list, and all his connections to
		// the todo list
		friends = append(friends, p)
		todo = append(todo, conns[p]...)
	}

	// Return the list of nodes, each is connected to each of the others
	return friends
}

// Are these two nodes connected?
func connected(n1, n2 string) bool {
	return in(n1, conns[n2])
}

// Does this string start with t?
func t(s string) bool {
	return s[0] == 't'
}

// Is this string in the list?
func in(s string, l []string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}
