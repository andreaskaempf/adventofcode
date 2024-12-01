// Advent of Code 2024, Day 01
//
// Given rows of pairs of numbers, extract the two lists, sort them, and
// find the sum of the absolute difference (Part 1). For Part 2, sum up
// the product of each number in the left list, times the number of times
// it occurs in the right list.
//
// AK, 01 Dec 2024

use std::fs;

fn main() {

    // Read lines of number pairs, convert to pairs of numbers
    let fname = "input.txt";
    let data = fs::read_to_string(fname).expect("Error reading");
    let rows: Vec<_> = data.lines().map(parse_nums).collect();

    // Separate into two separate lists
    // TODO: is there an easier way? Tried map taking first then second 
    // element, but second list gives borrow error since into_iter() 
    // takes ownership or rows.
    let mut list1 =  Vec::<i32>::new();
    let mut list2 =  Vec::<i32>::new();
    for r in rows {
        list1.push(r[0]);
        list2.push(r[1]);
    }

    // Sort both lists
    list1.sort();
    list2.sort();

    // Part 1: add up the absolute difference between the series
    let mut part1 = 0;
    for i in 0..list1.len() {
        part1 += (list1[i] - list2[i]).abs();
    }
    println!("Part 1 = {}", part1);

    // Part 2: sum up the first column time the number of times the element 
    // appears in the second column
    let mut part2 = 0;
    for i in list1 {
        part2 += i * occurs(i, &list2);
    }
    println!("Part 2 = {}", part2);

}

// Count up the number of times an element appears in a list
fn occurs(n: i32, l: &Vec<i32>) -> i32 {
    let mut c = 0;
    for i in l {
        if *i == n { // WTF!!!
            c += 1;
        }
    }
    c
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
