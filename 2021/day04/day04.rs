// Day 4 of Advent of Code 2021
//
// Simulate game of bingo, by applying list of called numbers
// until first board wins (for part 2, also find the last board
// that wins).
//
// AK, 8 January 2022

use std::fs;

fn main() {

    // Read the list of numbers to draw and all the boards from input file
    let (draws, mut boards) = read_data("input.txt");

    // Draw each number: update all the boards, then report any
    // new winners
    let mut winners: Vec<usize> = vec![]; // boards that have already won
    for n in draws {
        
        // Draw this number on all boards, setting drawn numbers to zero
        for bi in 0..boards.len() {
            draw_number(n, &mut boards[bi]);
        }

        // Check if any new boards have won. Report the first score for Part 1,
        // last score for Part 2.
        for bi in 0..boards.len() {
            if ! winners.contains(&bi) {   // ignore boards that have already won
                if board_wins(&boards[bi]) {
                    let sum = sum_numbers(&boards[bi]);
                    println!("Board {} won: {} sum of numbers left * {} number drawn = {}", bi, sum, n, sum * n);
                    winners.push(bi);
                }
            }
        }
    }
}

// Check if number appears on board, set it to zero if yes
fn draw_number(drawn: i32, board: &mut Vec<Vec<i32>>) -> bool {
    let mut ri = 0;
    while ri < board.len() {
        let mut ci = 0;
        while ci < board[ri].len() {
            if board[ri][ci] == drawn {
                board[ri][ci] = 0;
                return true;
            }
            ci += 1;
        }
        ri += 1;
    }
    false
}

// Check if a board has won, i.e., zeros in any row or column
fn board_wins(b: &Vec<Vec<i32>>) -> bool {

    // Check rows
    for r in b {        // each row
        let mut wins = true;
        for c in r {    //each cell
            if *c != 0 {
                wins = false;
                break;
            }
        }
        if wins {
            return true;
        }
    }

    // Check columns
    for ci in 0..b[0].len() {
        let mut wins = true;
        for r in b {    //  each row
            if r[ci] != 0 {
                wins = false
            }
        }
        if wins {
            return true;
        }
    }

    return false;
}

// Return the sum of the numbers on a board
fn sum_numbers(b: &Vec<Vec<i32>>) -> i32 {
    let mut sum = 0;
    for r in b {        // each row
        for c in r {    //each cell
            sum += *c;
        }
    }
    sum
}

// Read and parse input file
fn read_data(filename: &str) -> (Vec<i32>, Vec<Vec<Vec<i32>>>) {

    // Read input file
    let data = fs::read_to_string(filename).expect("Cannot read file");
    let lines: Vec<_> = data.lines().collect();

    // List of numbers drawn on first line
    let draws: Vec<_> = lines[0].split(",").map(|i| parse_int(i)).collect();

    // Read boards from remaining lines, lists of numbers, separated by newlines
    let mut boards: Vec<Vec<Vec<i32>>> = Vec::new();
    let mut board: Vec<Vec<i32>> = Vec::new();
    for i in 2..lines.len() {
        let row = lines[i];
        if row.len() == 0 {
            if board.len() > 0 {
                boards.push(board);
            }
            board = Vec::new();
        } else {
            let r: Vec<_> = row.split(" ").filter(|s| s.len() > 0).map(|i| parse_int(i)).collect();
            board.push(r);
        }
    }
    if board.len() > 0 {
        boards.push(board);
    }
    
    (draws, boards)
}

// Parse an integer string, return -1 if there was an error
fn parse_int(s: &str) -> i32 {
    match s.parse::<i32>() {
        Ok(n) => return n,
        Err(_) => return -999,
    }
}
