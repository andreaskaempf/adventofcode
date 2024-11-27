// Day 6 of Advent of Code 2021
//
// Simulate evolution of population of "lanterfish", starting at 5 but
// reproducing every 7 days. Doing this with simulation of a list as described
// in the problem is way too slow, so converted it to using a dictionary to
// keep track of the number of fish by age, much faster.
//
// AK, 11 Jan 2022

use std::fs;

fn main() {

    // Read data file, a long list of comma-separated numbers, and convert to list of numbers
    let data = fs::read_to_string("input.txt").expect("File not found");
    let nums: Vec<i64> = data.split(",").map(|i| i.trim().parse().expect("Invalid number")).collect();
    //println!("Data: {:?}", nums);

     // Count up the number of "fish" with each counter value
     let mut fish = [0; 10]; // initialize a 10-element array of zeroes
	for n in nums {
		fish[n as usize] += 1;
	}
	println!("Counts by age: {:?} => {}", fish, fish.iter().sum::<i64>());

	// Start simulating up to 80 days (256 days in part 2). Each day, fish
	// of age 0 becomes 6 and each adds a new fish of age 8 to the population,
	// while each other fish decreases its age by 1.
	// For test data, should be 26 fish after day 18, 5934 after 80 days, and
	// 26984457539 after 256 days
	//let mut t = 1;
	for t in 1..257 {   // number of days plus one
  
		// Each fish with a counter of 0 will create a new fish with counter 8 (but
		// don't do it yet, as we don't want to decrement the age of the new fish)
		let new_fish_8 = fish[0];

		// All fish with counter 0 will get reset to 6 (but don't do it yet, so
		// we don't decrement the age of those fish in the next step)
		let new_fish_6 = fish[0];
		fish[0] = 0;

		// All fish of ages 1-8 get decremented, by moving them down the chain:
		//   fish0 += fish1
		//   fish1 = 0
		// then
		//   fish1 += fish2
		//   fish2 = 0
		// etc.
		for age in 1..10 {  // age 1 to 9
			fish[age-1] += fish[age];
			fish[age] = 0;
		}

		// Finally, add the new fish of ages 6 and 8
		fish[6] += new_fish_6;
		fish[8] += new_fish_8;

		// Show current counts
		if t == 18 || t == 80 || t == 256 {
			println!("After {} days: {} fish", t, fish.iter().sum::<i64>());
		}
	}
}
