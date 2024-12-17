// Advent of Code 2024, Day 17
//
// Simulate execution of an 8-bit machine language program with opcodes and arcane instruction
// rules such as dividing, xoring, outputting numbers, and jumping. For Part 1, show the output of
// running a short "program", a list of 16 numbers. For Part 2, find the "register A" starting
// value that causes the program to output the same 16 values as the input program. Part 1 was
// simple implementation of virtual CPU. For Part 2, inspected outputs for incremental input
// values, and used patterns found to converge on likely input ranges, and used brute force from
// there.
//
// AK, 17 Dec 2024

use std::fs;

fn main() {

    // Read and parse "program" and register A from input file
    let data = fs::read_to_string("input.txt").expect("Read error");
    let lines: Vec<_> = data.split("\n").collect();
    let reg_a = lines[0].split(" ").nth(2).unwrap().parse::<u64>().unwrap();
    let pgm: Vec<u64> = lines[4].split(" ").nth(1).unwrap()
        .split(",").map(|s| s.parse::<u64>().unwrap()).collect();

    // Test cases from problem
    if false {
        // Test cases
        run_program(&vec![2,6], 0, 0, 9); // sets B to 1
        run_program(&vec![5,0,5,1,5,4], 10, 0, 0); // outputs 0,1,2
        run_program(&vec![0,1,5,4,3,0], 2024, 0, 0); // outputs 4,2,5,6,7,7,7,7,3,1,0 and leaves 0 in A.
        run_program(&vec![1,7], 0, 29, 0); // sets B to 26
        run_program(&vec![4,0], 0, 2024, 43690); // sets B to 44354*/

        // Sample input, should output 4,6,3,5,6,3,5,2,1,0.
        run_program(&vec![0,1,5,4,3,0], 729, 0, 0);
    }

    // Part 1: input, answer s/b 2,0,7,3,0,3,1,3,7
    run_program(&pgm, reg_a, 0, 0);

    // Part 2: What is the lowest positive initial value for register A that causes 
    // the program to output a copy of itself?
    // Example: let pgm: Vec<u64> = vec![0,3,5,4,3,0];  // converges at 117440
    // Did Part 2 by inspecting output lists, and converging towards likely
    // initial register A value that would reproduce program, as follows:
    // - Ran program with inputs from 0 to about a million
    // - Noticed that length of output starts at 1, but increases, i.e.,
    //   length becomes 2 at 8, 3 at 64, 4 at 512, 5 at 4096, 6 at 32768, 
    //   7 at 262144, and 8 at 2097152. Exploring this in a spreadsheet, 
    //   found that length = 8^(A-1).
    // - So for our target length of 16, start at A = 35184372088832.
    // - At any length, the last number(s) of the output appear to cycle, e.g.,
    //   start at 7, with zero (the last value of our input program) the last
    //   ending value.
    // - So started where length is 17 digits, and counted backwards in large
    //   increments, to find roughly where block of 16 with ending value of 0 starts.
    // - Counted forward from this in less large increments, to find where the output
    //   starts ending in 3, 0.
    // - Repeated same approach to find where the last few digits of the output
    //   match the input program, e.g.,
    //   let mut a = 35184372088832;                // output reaches 16 numbers
    //   let mut a = 281474976710656;               // output reaches 17 digits
    //   let mut a = 246290936710656 - 10000000;    // outputs that end in 3,0 start about here
    //   let mut a = 247390116710656 -1000000;      // 5,3,0 start about here
    //   let mut a = 247802433110656 - 100000;      // 5,5,3,0 start about here
    // - Then counted forward by 1 from here, and found answer
    let mut a = 247836792850656 - 10000;       // 6,5,5,3,0 starts about here
    loop {
        if run_program(&pgm, a, 0, 0) {
            println!("Found Part 2 at {}", a);
            break;
        }
        a += 1;  // vary this to explore large jumps
    }
}

// Run a program, given initial register values, and show final registers
// and output values. For part 2, return true if output equals input.
fn run_program(program: &Vec<u64>, reg_a: u64, reg_b: u64, reg_c: u64) -> bool {

    // Input data
    let mut a: u64 = reg_a;
    let mut b: u64 = reg_b;
    let mut c: u64 = reg_c;

    // println!("Running {:?} with A={}, B={}, C={}", program, a, b, c);
    let mut output = Vec::<u64>::new();
    let mut ip: usize = 0; // instruction pointer
    let two: u64 = 2;  // for pow(2, x)
    while ip < program.len() {

        // Get the operator and its literal and transformed "combo" operands
        // Combo: 0 - 3 = literal, 4 = A, 5 = B, 6 = C
        let op = program[ip];
        let  literal: u64 = program[ip+1];
        let mut combo = literal;
        if literal == 4 {
            combo = a;
        } else if literal == 5 {
            combo = b;
        } else if literal == 6 {
            combo = c;
        }

        // Op 0: "adv"  A / pow(2, operand), truncates and puts in A
        if op == 0 {
            a = a / two.pow(combo.try_into().unwrap());
            ip += 2;
        }

        // 1: "bxl" bitwise XOR of B and literal operand => B
        else if op == 1 {
            b = b ^ literal; // ^ is bitwise xor
            ip += 2;
        }

        // 2: "bst" calculates the value of its combo operand modulo 8 (thereby 
        // keeping only its lowest 3 bits), then writes that value to B
        else if op == 2 {
            b = combo % 8;
            ip += 2;
        } 
        
        // 3: "jnz" if A is nonzero, jumps by setting IP to the value of its literal,
        // otherwise IP skips +2 as normal
        else if op == 3 {
            if a != 0 {
                let n: usize = literal.try_into().unwrap();
                ip = n;
            } else {
                ip += 2;
            }
        }
        
        // 4: "bxc" calculates the bitwise XOR of B and C, then stores in B.
        // (For legacy reasons, this instruction reads an operand but ignores it.)
        else if op == 4 {
            b = b ^ c;
            ip += 2;
        }
        
        // 5: "out" outputs value of combo operand modulo 8
        else if op == 5 {
            let out: u64 = combo % 8;
            output.push(out);
            ip += 2;
        }
        
        // 6: "bdv" exactly like adv instruction except result stored in B
        else if op == 6 {
            b = a / two.pow(combo.try_into().unwrap());
            ip += 2;
        }

        // 7: "cdv" exactly like adv instruction except result stored in C 
        else if op == 7 {
            c = a / two.pow(combo.try_into().unwrap());
            ip += 2;
        }
    }

    // Show final answers, for Part 1, comment-out for Part 2 
    //println!("  final registers:  A = {}, B = {}, C = {}", a, b, c);
    //println!("  output: {:?}", output);

    // Uncomment this for finding patterns in Part 2
    //println!("{} => {:?} {}", reg_a, output, output.len());

    // For Part 2, return true if output same as program
    return output == *program;
}
