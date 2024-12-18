// Advent of Code 2024, Day 18
//
// Find shortest distance from one corner of a 71x71 grid to the oppposite corner,
// avoiding obstacles from a list of coordinates. For Part 1, just find the distance,
// using the first 1024 obstacles in the list. For Part 2, find the coordinates of the
// first obstacle that make it impossible to reach the location. Implemented Djikstra
// algorithm in Rust, works very well.
//
// AK, 18 Dec 2024

use std::fs;
use std::collections::HashMap;

// An x,y point
#[derive(Debug, Hash, Eq, Copy, Clone, PartialEq)]
struct Point {
    x: i32,
    y: i32
}

fn main() {

    // Set this to true for Part 2
    let part2 = false;

    // Read data file into a list of points
    //let fname = "sample12.txt";
    let fname = "input.txt";
    let data = fs::read_to_string(fname).expect("Missing file");
    let mut bytes = Vec::new();
    for l in  data.split("\n") {
        let pair: Vec<_> = l.split(",").map(atoi).collect();
        bytes.push(Point{x: pair[0], y: pair[1]});
        if bytes.len() == 1024 && !part2 { // stop at 1024 for Part 1
            break;
        }
    }

    // Part 1: get shortest path from top left to bottom right after 1024 bytes fallen,
    // note that second argument should be 7 for sample, 71 for input
    if !part2 {
        println!("Part 1 = {}", traverse(&bytes, 71)); // 338
        return;
    }

    // Part 2: find the first byte that causes blockage (answer was 20,44)
    let mut bytes1 = Vec::new(); // should really start with the first 1024 bytes
    while bytes1.len() < bytes.len() {
        let nb = bytes[bytes1.len()]; // next byte to add
        bytes1.push(nb);
        let dist = traverse(&bytes1, 71);
        println!("Trying {:?}, byte number {} => dist {}", nb, bytes1.len(), dist);
        if dist == 9999999 { // happens at iteration 2980
            break;
        }
    }
}

// Find shortest path from bottom-right to top-left, using Djikstra algorithm, avoiding any
// locations in the 'bytes' vector. Returns length of shortest path, or
// 999999 if no route found.
//
// I expected this was going to be about finding the shortest path with bytes falling in real time
// during the traversal, so decided to go backwards (started at the end point instead of
// beginning), and set up to determine if a byte had already fallen at a given point in the
// simulation. Not required, sigh ...
fn traverse(bytes: &Vec<Point>, grid: i32) -> i32 {

    // Set up priority queue with just starting point
    let start = Point{x: grid-1, y: grid-1}; // current position starts at bottom right
    let end = Point{x:0, y:0}; // finish at top left
    let mut q = Vec::new(); // (priority) queue of points to explore
    q.push(start);

    // Set up maps of visited points, and shortest distance found to each point
    let mut visited = HashMap::new();
    let mut dist = HashMap::new();
    dist.insert(start, 0);      // lowest cost to each position
    dist.insert(end, 9999999);  // initial cost to end is infinity

    // Start exploring, while there are still points to explore
    while q.len() > 0 {

        // Sort the queue in ascending order, so shortest distance is at the front,
        // and remove the first point (i.e., closest one)
        q.sort_by(|a, b| dist[a].cmp(&dist[b]));
        let p = q.remove(0);

        // Skip if already visited this
        if visited.contains_key(&p) { 
            continue;
        }

        // Note that reached end
        if p == end {
            println!("Reached end");
        }

        // Mark this point as visited so we don't return there
        visited.insert(p.clone(), true);

        // Cost of next step from here
        let next_dist = dist[&p]+1;

        // Look in each direction, consider going there if free
        let adjacent = vec![Point{x: p.x+1, y:p.y}, Point{x: p.x-1, y:p.y}, 
            Point{x: p.x, y:p.y+1}, Point{x: p.x, y:p.y-1}];
        for p1 in adjacent { 
           
            // Reject if out of bounds
            if p1.x < 0 || p1.x >= grid || p1.y < 0 || p1.y >= grid {
                continue;
            }

            // Reject if already visited
            if visited.contains_key(&p1) {
                continue;
            }

            // Reject if there is an obstacle here at this time in the simulation
            if byte_arrival_time(&p1, bytes) >= 0  {
                continue;
            }

            // Add to queue
            q.push(p1);

            // Update lowest distance to this next point
            if dist.contains_key(&p1) {
                if next_dist < dist[&p1] {
                    dist.insert(p1, next_dist);
                }
            } else {
                dist.insert(p1, next_dist);
            }
            
        }
    }

    // Return value is lowest distance to ending point, will be infinite if not reached
    dist[&end]
}

// Arrival time of a byte, just its index in the list, used here just to see if a given
// location has a byte in it (obstacle)
fn byte_arrival_time(p: &Point, bytes: &Vec<Point>) -> i32 {
    for i in 0..bytes.len() {
        if bytes[i] == *p {
            return i.try_into().unwrap();
        }
    }
    -1
}

// Parse integer
fn atoi(s: &str) -> i32 {
    s.parse::<i32>().expect("Bad number")
}
