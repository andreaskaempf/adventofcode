// day15.go
//
// Find the lowest cost path through a graph, starting at the top right, and
// ending up at the bottom right, adding up any cells you enter, so as to
// minimize total cost. Uses the yourbasic/graph library, very fast.
//
// AK, 15/12/2021

package main

import (
	"bufio"
	"fmt"
	"github.com/yourbasic/graph"
	"os"
)

// Read data file into a matrix of integers (basically a list of rows, since Go
// does not have matrices)
func readMatrix(fname string) [][]int64 {

	// Open file
	f, err := os.Open(fname)
	if err != nil {
		return nil
	}

	// Read each line of input file
	m := [][]int64{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		r := []int64{}
		for i := 0; i < len(t); i++ {
			r = append(r, int64(t[i]-'0'))
		}
		m = append(m, r)
	}
	return m
}

// Convert the matrix into a graph using the yourbasic/graph library, by
// treating each horizontal and vertical pair of adjacent digits as an edge,
// with a weight equal to the value in the target cell. Then, use the library's
// shortest path implementation to find the lowest cost from top left to bottom
// right.
func buildGraph(m [][]int64) {

	// Build a graph with enough capacity for all nodes
	nr := len(m)
	nc := len(m[0])
	g := graph.New(nr * nc)

	// For each row, add right-pointing horizontal connections
	for ri := 0; ri < len(m); ri++ {
		for ci := 1; ci < len(m[0]); ci++ {
			cost := m[ri][ci]
			n0 := ri*len(m[0]) + (ci - 1)
			n1 := ri*len(m[0]) + ci
			//fmt.Printf("Adding row edge: %d -> %d = %d\n", n0, n1, cost)
			g.AddCost(n0, n1, cost)
		}
	}

	// For each column, add downward vertical connections
	for ci := 0; ci < len(m[0]); ci++ {
		for ri := 1; ri < len(m); ri++ {
			cost := m[ri][ci]
			n0 := (ri-1)*len(m[0]) + ci
			n1 := ri*len(m[0]) + ci
			//fmt.Printf("Adding col edge: %d -> %d = %d\n", n0, n1, cost)
			g.AddCost(n0, n1, cost)
		}
	}

	// Calculate shortest path from first to last node
	path, dist := graph.ShortestPath(g, 0, nr*nc-1)
	fmt.Println("Shortest path:", path)
	fmt.Println("Total distance:", dist)
}

// Expand the matrix by the factor provided as described in the
// problem statement, i.e., adding one to each horizontal and
// vertical "tile" added.
func expandMatrix(m [][]int64, n int) [][]int64 {

	// Expand each row
	nr := len(m)
	nc := len(m[0])
	fmt.Printf("Before expansion: %d x %d\n", nr, nc)

	// Expand horizontally, adding one for each expansion
	for ri := 0; ri < nr; ri++ { // each row
		for h := 1; h < n; h++ { // each horizontal duplication
			for ci := 0; ci < nc; ci++ {
				m[ri] = append(m[ri], m[ri][ci]+int64(h))
			}
		}
	}

	// Now duplicate each row 5 times, similarly increasing each time
	nc = len(m[0])
	for v := 1; v < n; v++ { // each vertical duplication
		for ri := 0; ri < nr; ri++ { // original each row
			r := []int64{}
			for ci := 0; ci < nc; ci++ {
				r = append(r, m[ri][ci]+int64(v))
			}
			m = append(m, r)
		}
	}

	// Adjust any cells > 9
	nr = len(m)
	for ri := 0; ri < nr; ri++ {
		for ci := 0; ci < nc; ci++ {
			for m[ri][ci] > 9 {
				m[ri][ci] -= 9
			}
		}
	}

	fmt.Printf("After expansion: %d x %d\n", nr, nc)

	// Return the expanded matrix (could have also used a pointer
	// to alter original)
	return m
}

// Main execution: run parts 1 and 2
func main() {

	// Read matrix from file
	//m := readMatrix("sample.txt")
	m := readMatrix("input.txt")

	// Build graph and compute shortest path on small matrix
	buildGraph(m)

	// Expand the matrix to 5x, increasing counts each cell, and do the graph
	// calculation again
	m = expandMatrix(m, 5)
	buildGraph(m)
}
