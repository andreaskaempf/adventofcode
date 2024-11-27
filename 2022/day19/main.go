// Advent of Code 2022, Day 19
//
// Basically a set of optimizations, to find the maximum number of "geodes"
// that can be produced over 24 time periods from a set of four types of
// "robots".
//
// Each robot can produce one mineral of its own kind each time period. There
// are 30 "blueprints" (cost schedules), each of which lists the number of each
// type of ingredient required to build a robot. So it's a production plan
// optimization. Part 1 asks you to optimize all 30 "blueprints", Part 2 only
// the first 3 blueprints, but for 32 periods instead of 24.
//
// First attempt was brute force (removed, but still in earlier versions in the
// repository), got the answer for Part 1 after many hours of execution time.
// Revisited and wrote as a recursive depth-first search, that stops when a
// branch fails to achieve the same number of "geodes" as another branch, at
// the same time in the simulation.  Runs concurrently, takes about .5 seconds
// for both parts.
//
// AK, 19-23 Dec 2022

package main

import (
	"fmt"
	"strings"
)

// A blueprint, set of recipes for creating the different types of robots
type Blueprint struct {
	number  int
	recipes []Recipe
	bestAt  map[int]int // keep track of best found at every time step
}

// Recipe for making a type of robot, i.e., list of ingredients
type Recipe struct {
	robotType   string       // the type of robot
	ingredients []Ingredient // what is required
}

// One ingredient required for making a type of robot
type Ingredient struct {
	material string
	requires int
}

// Number of minutes (24 for part 1, 32 for part 2)
var minutes int = 24

func main() {

	// Read blueprints
	fname := "sample.txt"
	fname = "input.txt"
	blueprints := readBlueprints(fname)
	fmt.Println("Part 1: optimizing", len(blueprints), "blueprints")

	// For Part 1, optimize all 30 blueprints in parallel, sum up the
	// maximum number of geodes possible multiplied by the blueprint
	// number
	ch := make(chan int)
	for i := 0; i < len(blueprints); i++ {
		go optimize(&blueprints[i], ch)
	}

	// Retrieve results and add them up for Part 1
	part1 := 0
	for i := 0; i < len(blueprints); i++ {
		x := <-ch
		part1 += x
	}
	fmt.Println("Part 1:", part1)

	// Part 2
	fmt.Println("\nPart 2 starting")
	ch2 := make(chan int)
	part2 := 1                    // initialize multiplier
	minutes = 32                  // 32 mins instead of 24!
	nbp := mn(3, len(blueprints)) // only do up to 3 blueprints
	for i := 0; i < nbp; i++ {    // start in background
		go optimize(&blueprints[i], ch2)
	}
	for i := 0; i < nbp; i++ { // await results
		x := <-ch2
		part2 *= x
	}
	fmt.Println("Part 2:", part2)
}

// Find the maximum number of geodes you can produce in 24 minutes, given
// the cost structure in the blueprint:
//   - Start with one ore robot
//   - Each robot produces one unit of its type per minute
//   - You can build one new robot per time step, if you have the necessary
//     materials on-hand (defined in the blueprint)
//   - But the robot is only ready at the end of the time step
func optimize(bp *Blueprint, ch chan int) {

	// Initialize lists of materials we have, and robots we have
	fmt.Println("Starting blueprint", bp.number)
	robots := map[string]int{"ore": 1}
	materials := map[string]int{}

	// Prepare to keep track of best result at each time, stored in the
	// blueprint rather than in global variable so we can run in parallel
	bp.bestAt = map[int]int{}

	// Start the optimization at time 1
	geodes := optimize1(bp, robots, materials, 1)
	fmt.Printf("Blueprint %d: %d geodes\n", bp.number, geodes)

	// Send result back to channel
	if minutes == 24 { // for part 1, multiply geods by blueprint number
		geodes *= (*bp).number
	}
	ch <- geodes
}

// One recursive step of the optimization, returns the number of geodes made
// to end of simulation. Note that function call changes maps, so need to work
// on copies for each recursive call.
func optimize1(bp *Blueprint, robots, materials map[string]int, time int) int {

	// Update materials supply with the robots we already have
	materials0 := copyMap(materials) // work on a copy
	for rtype, n := range robots {   // each robot produces one of its kind
		materials0[rtype] += n // each time step
	}

	// If out of time, return final number of geodes
	if time >= minutes {
		return materials0["geode"]
	}

	// Stop here if number of geodes is not the best achieved so far by
	// this time
	if materials0["geode"] < bp.bestAt[time] {
		return 0
	} else if materials0["geode"] > bp.bestAt[time] {
		bp.bestAt[time] = materials0["geode"]
	}

	// Check all the recipes we *could* build, given materials we now have.
	cb := []Recipe{}
	grec := -1
	orec := -1
	for i, recipe := range bp.recipes {
		if canBuild(recipe, materials) { // not materials0!
			if recipe.robotType == "geode" {
				grec = i
			} else if recipe.robotType == "obsidian" {
				orec = i
			}
			cb = append(cb, recipe)
		}
	}

	// If we have enough to build a geode, try only that. Otherwise, if we
	// have enough to build an obsidian, try only that. Otherwise, try the
	// other robot types we have enough for.
	if grec >= 0 {
		cb = []Recipe{bp.recipes[grec]}
	} else if orec >= 0 {
		cb = []Recipe{bp.recipes[orec]}
	}

	// If we can't build anything yet, try simulating the next time step without
	// building anything (necessary early on in the simulation)
	best := 0 // best result achieve this iteration
	if len(cb) == 0 {
		best = optimize1(bp, robots, materials0, time+1)
	}

	// For each robot type we could build, simulate building it and continuing
	for _, recipe := range cb {

		// Use up required materials, and add the new robot
		materials1 := copyMap(materials0) // don't alter original
		for _, i := range recipe.ingredients {
			materials1[i.material] -= i.requires
		}

		// Add the new robot
		robots1 := copyMap(robots) // don't alter original
		robots1[recipe.robotType] += 1

		// Try next time step with this configuration
		geodes := optimize1(bp, robots1, materials1, time+1)
		if geodes > best {
			best = geodes
		}
	}

	// Return the best number of geodes achieved to the end of the simulation
	// from any branch tried here
	return best
}

// Can we build at least one of this recipe with materials on-hand?
func canBuild(recipe Recipe, materials map[string]int) bool {
	for _, i := range recipe.ingredients {
		if materials[i.material] < i.requires { // redundant
			return false
		}
	}
	return true
}

// Make a deep copy of a string->int map
func copyMap(m map[string]int) map[string]int {
	m2 := map[string]int{}
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

// Minimum of two numbers
func mn(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Read blueprints from file and return list of them
// E.g., Blueprint 1 (can create 9 geodes in 24 mins):
// - Each ore robot costs 4 ore.
// - Each clay robot costs 2 ore.
// - Each obsidian robot costs 3 ore and 14 clay.
// - Each geode robot costs 2 ore and 7 obsidian.
//
//	bp1 := []Recipe{
//	  Recipe{"ore", []Ingredient{
//	    Ingredient{"ore", 4}}},
//	  Recipe{"clay", []Ingredient{
//	    Ingredient{"ore", 2}}},
//	  Recipe{"obs", []Ingredient{
//	    Ingredient{"ore", 3}, Ingredient{"clay", 14}}},
//	  Recipe{"geode", []Ingredient{
//	    Ingredient{"ore", 2}, Ingredient{"obs", 7}}},
//	}
func readBlueprints(fname string) []Blueprint {

	// Parse each blueprint, one per line
	blueprints := []Blueprint{}
	for _, l := range readLines(fname) {

		parts := strings.Split(l, ":")
		bp := Blueprint{number: len(blueprints) + 1}
		costs := strings.Split(parts[1], ".")
		for i := 0; i < len(costs)-1; i++ {
			words := strings.Split(strings.TrimSpace(costs[i]), " ")
			rec := Recipe{robotType: words[1]}
			ing := Ingredient{words[5], atoi(words[4])}
			rec.ingredients = append(rec.ingredients, ing)
			if len(words) > 6 {
				ing = Ingredient{words[8], atoi(words[7])}
				rec.ingredients = append(rec.ingredients, ing)
			}
			bp.recipes = append(bp.recipes, rec)
		}
		blueprints = append(blueprints, bp)
	}
	return blueprints
}
