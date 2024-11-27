// Advent of Code 2021, Day 1 (Rust version)
//
// Read a text file of numbers, report the number of increases,
// first between subsequent lines, then between subsequent sums
// of 3 lines.
//
// To compile this Rust program: rustc day01.rs
// To run: ./day01 (or use cargo to build & run)
// 
// AK, 3/01/2022

use std::fs;

fn main() {

    // Read file as one string, convert to list of numbers
    let fname = "input.txt";
    println!("Reading {}", fname);
    let data = read_lines(fname);
    let nums = lines_to_integers(&data);
    println!("{} numbers", nums.len());

    // Part 1: Count the number of increases
    let mut prev_n = -1;    // to avoid counting the first number as an increase
    let mut increases = 0;
    let mut i = 0;
    while i < nums.len() {
        let n = nums[i]; 
        if prev_n >= 0 && n > prev_n { 
            increases += 1;
        }
        prev_n = n;
        i += 1;
    }

    println!("Part 1: there were {} increases (should be 1832)", increases);

    // Part 2: Count the number of increases between consecutive 3-value windows
    let mut prev_w = -1;  // to avoid counting the first number as an increase
    let mut w_increases = 0;
    let mut i = 2;  // end of first window
    while i < nums.len() {
        let w = nums[i] + nums[i-1] + nums[i-2]; 
        if prev_w >= 0 && w > prev_w {
            w_increases += 1;
        }
        prev_w = w;
        i += 1;
    }

    println!("Part 2: there were {} window increases (should be 1858)", w_increases);
}


// Function to read lines from a file and return a vector of strings
fn read_lines(fname: &str) -> Vec<String> {

    // Read whole file into memory
    let data = fs::read_to_string(fname)
        .expect("Error reading file");

    // Split by newlines, return vector of Strings
    data.lines().map(|s| s.to_string()).collect()
}

// Convert a vector of strings containing numbers, to a vector of integers
fn lines_to_integers(lines: &Vec<String>) -> Vec<i32> {
    lines.iter().map(|l| parse_int(l)).collect()
}


// Parse an integer string, return -1 if there was an error
fn parse_int(s: &str) -> i32 {
    match s.parse::<i32>() {
        Ok(n) => return n,
        Err(_) => return -1,
    }
}

