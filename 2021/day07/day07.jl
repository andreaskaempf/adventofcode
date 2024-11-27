# Advent of Code, day 7
#
# Find target position that minimizes fuel use for a population
# of crabs that need to move there, under different assumptions
# about fuel use per each crab's distance moved (optimization
# problem, but used brute force).
#
# AK, 7 Dec 2021

# Read input data, convert to list of numbers
#raw_data = readlines(open("test_input.txt"))
raw_data = readlines(open("input.txt"))
data = map(x -> parse(Int64, x), split(raw_data[1], ","))
maxPos = maximum(data)
#println("Data:", data)
println("Max position =", maximum(data))

# Find the "horizontal" position that minimizes fuel use.  For part 1 (function
# argument true), use one unit of fuel expended per step. For part 2 (false),
# fuel use increases by 1 for each step taken.
function crabs(part1::Bool)

    bestPos = 0
    bestFuel = 0

    # Try each possible position
    for pos in 1:maxPos

        # Add up fuel use
        fuel = 0
        for crab in data
            dist = abs(crab - pos)
            if part1
                fuel += dist
            else
                fuel += sumnums(dist)
            end
        end
        #println("Position $pos: fuel = $fuel")

        # Remember if better than last (i.e., less fuel use)
        if bestPos == 0 || fuel < bestFuel
            bestPos = pos
            bestFuel = fuel
        end
    end

    # Show the best result found
    println("Best position $bestPos: fuel = $bestFuel")

end

# Sum up integers up to n
function sumnums(n)
    nn = 0
    for i in 1:n
        nn += i
    end
    return nn
end

println("--- PART 1 ---")
crabs(true)

println("--- PART 2 ---")
crabs(false)

