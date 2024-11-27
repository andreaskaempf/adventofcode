// Day 5 of Advent of Code 2021
//
// Trace lines on a grid, given starting and ending coordinates of
// each line, and count up where 2 or more lines cross (part 1 only
// includes horizontal and vertical lines, part 2 diagonal as well).
//
// This Rust implementation is different from the earlier one in Go,
// in that it represents the ocean floor as a sparse matrix using a
// hashmap (which can be summed up to get the number of crossings), 
// and also reads the input file directly, without the need
// to replace the arrows with commas externally.
//
// AK, 10 January 2022

use std::fs;
use std::collections::HashMap;

fn main() {

    // Set this to false for Part 2
    let part1 = false;

    // Create a hashmap of (x,y) pairs, to represent a sparse matrix of the number of
    // lines crossing each point
    let mut floor: HashMap<(i32,i32), i32> = HashMap::new();

    // Read the input file, each like: 723,326 -> 156,893
    // i.e., x1,y1 -> x2,y2
    let data = fs::read_to_string("input.txt").expect("File not found");
    for l in data.lines() {

        // Parse coordinates
        let l1 = l.replace(" -> ", ",");  // replace arrow with comma
        let parts: Vec<i32> = l1.split(",").map(|s| s.parse().expect("?")).collect();
        let x1 = parts[0];  // TODO: is there a more concise way?
        let y1 = parts[1];
        let x2 = parts[2];
        let y2 = parts[3];

        // Part 1: horizontal or vertical lines only
        if part1 && !(x1 == x2 || y1 == y2) {
            continue
        }

        // Calculate step size +/- 1 or 0 if horiz/vertical
        let dx = diff(x1, x2);
        let dy = diff(y1, y2);

        // Traverse line point by point
        let mut x = x1;
        let mut y = y1;
        loop {

            // Set the point to 1 if crossing
            let k = (x,y);
            if floor.contains_key(&k) {
                floor.insert(k, 1); // only care about whether 2 or more points
            } else {
                floor.insert(k, 0);
            }

            // Stop if reached end of this line
            if x == x2 && y == y2 {
                break;
            }

            // Next point on this line
            x += dx;
            y += dy;
        }
        
    }

    // Now add up the number of points with overlap
    let overlaps: i32 = floor.values().sum();
    println!("{} points overlap (s/b 5442 for Part 1, 19571 for Part 2)", overlaps);
}

// Return amount that needs to be added each step to get from a to b
// (will be 1 if increasing, -1 if decreasing, 0 if horiz/vertical)
fn diff(a: i32, b: i32) -> i32 {
    if a < b {
        1
    } else if a > b {
        -1
    } else {
        0
    }
}
