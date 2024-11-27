// Advent of Code 2021, Day 2 (Rust version)
//
// Read a text file of numbers, report the number of increases,
// first between subsequent lines, then between subsequent sums
// of 3 lines.
// 
// AK, 4/01/2022

use std::fs;

fn main() {

    // Read rows of data from file
    let fname = "input.txt";
    println!("Reading {}", fname);
    let data = read_lines(fname);
    println!("{} lines", data.len());

    // State variables for problems 1 & 2 (all init to zero)
    let mut horiz1 = 0;
    let mut depth1 = 0;
    let mut horiz2 = 0;
    let mut depth2 = 0;
    let mut aim2 = 0;

    // Process each row
    for r in data {

        // Split out instruction and number (e.g., "forward 2")
        let parts: Vec<&str> = r.split_whitespace().collect();
        let instruction = parts[0];
        let n = parse_int(parts[1]);

        // Process instructions for both problems
        if instruction == "forward" {
            horiz1 += n;
            horiz2 += n;
            depth2 += aim2 * n;
        } else if instruction == "down" {
            depth1 += n;
            aim2 += n;
        } else if instruction == "up" {
            depth1 -= n;
            aim2 -= n;
        } else {
            println!("Bad instruction: {}", instruction)
        }
    }

    // Show results for both problems
    println!("Problem 1: Ending pos = {}, depth = {}, product = {}", horiz1, depth1, horiz1*depth1);
    println!("Problem 2: Ending pos = {}, depth = {}, aim = {}, product = {}", horiz2, depth2, aim2, horiz2*depth2);
}

// Function to read lines from a file and return a vector of strings
fn read_lines(fname: &str) -> Vec<String> {

    // Read whole file into memory
    let data = fs::read_to_string(fname)
        .expect("Error reading file");

    // Split by newlines, return vector of Strings
    data.lines().map(|s| s.to_string()).collect()
}

// Parse an integer string, return -1 if there was an error
fn parse_int(s: &str) -> i32 {
    match s.parse::<i32>() {
        Ok(n) => return n,
        Err(_) => return -1,
    }
}

