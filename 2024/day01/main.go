// Advent of Code 2024, Day 01
//
//
//
// AK, 01 Dec 2024

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	// Read the input file into a list of byte vectors (remove any blank rows first)
	fname := "sample.txt"
	//fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	rows := bytes.Split(data, []byte("\n"))

	// Process each row
	for _, l := range rows {
		fmt.Println(string(l))
	}
}
