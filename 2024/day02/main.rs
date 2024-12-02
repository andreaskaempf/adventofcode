// Advent of Code 2024, Day 02
//
// Given rows of numbers, check each for the following conditions:
// all numbers are increasing/decreasing, and by 1 to 3 each step.
// For Part 1, count up how many rows meet the condition. For Part 2,
// count up how many rows would meet the condition, if any digit were
// removed.
//
// AK, 02 Dec 2024

use std::fs;

fn main() {
    
    // Read input file, convert to lists of numbers
    //let fname = "sample.txt";
    let fname = "input.txt";
    let data = fs::read_to_string(fname).expect("Read error");
    let rows: Vec<_> = data.lines().map(parse_nums).collect();

    // Check each row, count matches
    let part1: i32 = rows.clone().into_iter().map(check).sum();
    let part2: i32 = rows.into_iter().map(check2).sum();
    println!("Part 1 = {}, Part 2 = {}", part1, part2);
}

// For Part 2, do the same check but try removing each number from the list,
// and resport success if any of these lists match
fn check2(row: Vec<i32>) -> i32 {

    // Try removing each number from the list
    for i in 0..row.len() {
        let mut rowx = row.clone();
        rowx.remove(i);
        if check(rowx) == 1 {
            return 1;
        }
    }
    return 0;
}

// Check a row: numbers must be all increasing or all decreasing,
// by between 1, 2 or 3 steps each time
fn check(row: Vec<i32>) -> i32 {

	// Assume a rows with less than 2 elements is not valid, since it can't
	// be increasing/decreasing
	if row.len() < 2 {
		return 0;
	}

	// Check the "direction" of the first to elements of this row, must be
	// incrasing or decreasing
	let direction = (row[1] - row[0]).signum();	
    if direction == 0 {
		return 0;
	}

	// Check this row: must be either all increasing or all decreasing,
	// always in the same direction
	for i in 1..row.len() {
		let delta = (row[i] - row[i-1]).abs();
		if (row[i] - row[i-1]).signum() != direction || delta < 1 || delta > 3 {
			return 0;
		}
	}

	// Passed all checks
	return 1;
}

// Parse a string with space-separated numbers into a list of ints
fn parse_nums(s: &str) -> Vec<i32> {
    s.split_whitespace().map(parse_int).collect()
}

// Parse an integer
// TODO: is there a simpler way to do this?
fn parse_int(n: &str) -> i32 {
    n.to_string().parse::<i32>().unwrap()
}
