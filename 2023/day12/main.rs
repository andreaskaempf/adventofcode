use std::fs;

fn main() {

    // Read input file, each line a sequence of chars and a trio of numbers
    let data = fs::read_to_string("sample.txt").expect("Read error");
    let mut part1 = 0;
    for l in data.lines() {

        // Get chars and numbers
        let space: usize = l.find(' ').unwrap();
        let chars = &l[..space];
        let nums: Vec<_> = l[(space+1)..].split(',').map(parse_int).collect();
    
        // Analyse this line
        let res = analyze(chars, &nums);
        println!("{:?} -- {:?} ==> {}", chars, nums, res);
        part1 += res;
        
    }
    println!("Part 1 (s/b 1,4,1,1,4,10 = 21) = {}", part1);
}

// From the current position, 
fn analyze(chars: &str, nums: &Vec<usize>) -> usize {

    //println!("{}", chars);

    // No more characters left, so done
    if chars.len() == 0 {
        return 0;
    }

    // If no numbers left, one result found
    if nums.len() == 0 {
        return 1;
    }

    // Number of chars we need to match, return 0 if not enough
    let n = nums[0];
    if chars.len() < n {
        return 0;
    }

    // Are the next n chars eligible, i.e., hash or question mark?
    let mut ok: usize = 1;
    for i in 0..n {
        if chars.chars().nth(i).unwrap() == '.' {
            ok = 0;
        }
    }

    // Even if it's okay, the sequence needs to be at the end or
    // followed by either a period or question mark
    if ok > 0 {
        if chars.len() >= n && chars.chars().nth(n).unwrap() == '#' {
            ok = 0;
        }
    }

    // Continue from next position in string
    let restc = &chars[n..];
    let restn: Vec<usize> = nums[1..].to_vec();
    return analyze(restc, &restn) + ok;
}

// Parse an integer into a usize
fn parse_int(s: &str) -> usize {
    s.parse::<usize>().unwrap()
}