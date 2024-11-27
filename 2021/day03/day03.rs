// Advent of Code 2021, Day 3 (Rust version)
//
// Read a text file of binary digits, and make up a new number, consisting
// of the most frequent digit in each column. For Part 2, find the numbers
// that have the most/least frequent values in each successive bit, and
// multiply them to produce result.
// 
// AK, 6/01/2022

use std::fs;

fn main() {

    // Read rows of data from file
    //let fname = "sample.txt";
    let fname = "input.txt";
    println!("Reading {}", fname);
    let data = read_lines(fname);
    println!("{} lines", data.len());

    // Convert list of strings to a list of lists of 1/0 integers

    let m: Vec<Vec<i32>> = data.iter().map(|r| parse_binary(r)).collect();
    let ncols = m[0].len();  // number of columns
    //println!("{:?}", m);

    // Count up the most/least frequent digits in each column:
    // Gamma for this column is 1 if the column is mostly ones,
    // Epsilon the reverse
    let mut gamma_bits: Vec<i32> = Vec::new();
    let mut epsilon_bits: Vec<i32> = Vec::new();
    let ones = count_ones(&m); // count up ones in each column
    for i in 0..ncols {  // each column
        if ones[i] >= (m.len() as i32) / 2 {
            gamma_bits.push(1);
            epsilon_bits.push(0);
        } else {
            gamma_bits.push(0);
            epsilon_bits.push(1);
        }
    }

    // Convert to decimal and show result
    let gamma = binary_to_decimal(&gamma_bits);
    let epsilon = binary_to_decimal(&epsilon_bits);
    println!("Part 1: gamma = {}, epsilon = {}, answer = {}", gamma, epsilon, gamma * epsilon);

    // Part 2: Go through numbers, start at first column, keep only numbers
    // that meet bit criteria, until only one left:
    // - oxygen: most common value (keep 1 if equally common)
    // - co2: least common value (keep 0 if equally common)

    // Oxygen: filter list to numbers that have most frequent 1/0 in
    // the current column
    let mut m2 = copy_matrix(&m);
    let mut c = 0;  // current column
    while m2.len() > 1 && c < ncols {

        // Determine which bit is most common in this column
        let ones = count_ones(&m2);  // redo this for each column
        let mut bit = 0; // default value
        if (ones[c] as f32) / (m2.len() as f32) >= 0.5 {
             bit = 1;
        }

        // Filter out the rows that have this value for the column
        m2 = filter_rows(&m2, c, bit);
        //println!("  Col {}: {} left", c, m2.len());
        c += 1;  // next column
    }

    let ox = binary_to_decimal(&m2[0]);
    println!("Oxygen = {} (should be 3583)", ox);

    // CO2: filter list to numbers that have least frequent 1/0 in
    // the current column
    m2 = copy_matrix(&m);
    c = 0;  // current column
    while m2.len() > 1 && c < ncols {

        // Determine which bit is least common in this column
        let ones = count_ones(&m2);  // redo this for each column
        let mut bit = 0; // default value
        if (ones[c] as f32) / (m2.len() as f32) < 0.5 {
             bit = 1;
        }

        // Filter out the rows that have this value for the column
        m2 = filter_rows(&m2, c, bit);
        //println!("  Col {}: {} left", c, m2.len());
        c += 1;  // next column
    }

    let co2 = binary_to_decimal(&m2[0]);
    println!("CO2 = {} (should be 1601)", co2);

    // Compute the final answer
    println!("Answer = {} (should be 5736383)", ox * co2);
}

// For a list of 1/0 vectors, count up the 1s in each column
fn count_ones(m: &Vec<Vec<i32>>) -> Vec<i32> {
    let mut ones = Vec::new();
    let ncols = m[0].len();
    for i in 0..ncols {    // each column
        let mut o = 0;
        for r in m {
            o += r[i];
        }
        ones.push(o);
    }
    ones
}

// Given a list of 1/0 vectors, return a subset of the rows
// that contain the given value in the selected column
// TODO: Should be possible to use a filter, but this does not work:
// m2 = m2.iter().filter(|r| r[c] == bit).collect();
fn filter_rows(rows: &Vec<Vec<i32>>, col: usize, val: i32) -> Vec<Vec<i32>> {
    let mut res: Vec<Vec<i32>> = Vec::new();
    for r in rows {
        if r[col] == val {
            res.push(r.clone());
        }
    }
    res
}

// Make a deep copy of a matrix, i.e., a vector of integer vectors
fn copy_matrix(m: &Vec<Vec<i32>>) -> Vec<Vec<i32>> {
    let mut m2: Vec<Vec<i32>> = Vec::new();
    for r in m.iter() { // each row
        m2.push(r.clone());
    }
    m2
}

// Function to read lines from a file and return a vector of strings
fn read_lines(fname: &str) -> Vec<String> {

    // Read whole file into memory
    let data = fs::read_to_string(fname)
        .expect("Error reading file");

    // Split by newlines, return vector of Strings
    data.lines().map(|s| s.to_string()).collect()
}

// Parse a string of '1'/'0' chars into a vector of 1/0 ints
fn parse_binary(s: &str) -> Vec<i32> {
    let res: Vec<i32> = s.chars().map(|c| c.to_digit(10).unwrap_or(999) as i32).collect();
    res
}

// Convert a vector of binary 1/0 digits into decimal
fn binary_to_decimal(digits: &Vec<i32>) -> i32 {
    let mut result = 0;
    for d in digits {
        result = result * 2 + d;
    }
    result
} 
