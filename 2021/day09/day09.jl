# Advent of Code, Day 9
#
# Find lowest points in a 2-dimensional space, and the areas around these low
# points, bounded by edges or 9s.
#
# AK, 9 Dec 2021

# Read input, convert to list of vectors, one per row
f = "test_input.txt"
f = "input.txt"
d = split(strip(read(f, String)), "\n")
m = map(r -> [x - '0' for x in r], d)
#display(m)
#println("\n")

# Initialize lists of basins: every basin is centred around a low point, and
# consists of list of points around that point (a point is [row, col] tuple).
basins = Vector{Vector{Vector{Int64}}}()

# Part 1: Find the low points - the locations that are lower than any of its adjacent
# locations -- and add together their "risk", i.e., the value in the cell plus 1.
# Also initialize a list of basins around each low point, and set the value of the low
# point to -1 to set up for part 2.
function part1()

    totalRisk = 0
    for ri in 1:length(m)
        r = m[ri]
        for ci in 1:length(r)

            # If this is this a low point, add up total risk,
            # and set cell to -1 and initialize basin to set up for part 2
            x = r[ci]
            lowX = (ci == 1 || x < r[ci-1]) && (ci == length(r) || x < r[ci+1])
            lowY = (ri == 1 || x < m[ri-1][ci]) && (ri == length(m) || x < m[ri+1][ci])
            if lowX && lowY
                totalRisk += x + 1
                m[ri][ci] = -1   # mark the low point, first in this basin
                #println("Found low point at $ri,$ci")
                push!(basins, [[ri,ci]])  # create start of a new basin
            end
        end
    end

    println("Part 1: total risk = $totalRisk")

end

# Part 2: find the size of the three largest "basins" around the low points, i.e.,
# the areas around low points bounded by edges or 9s, and multiply these three
# sizes together. Since a basin can include points adjacent to a cell found by
# search left/right/up/down from the original low point, need to iterate until
# no more points are found. Then, just add up the lengths of each basin and find
# the highest values.
function part2()

    # Find all basins
    println("Part 2: finding adjacent points")
    iter = 1
    while findBasins()
        println("  iteration ", iter)
        iter += 1
    end
    
     # Sort the basin lengths, and multiply the top 3 together
     #println("Ending basins: ")
     #for b in basins
     #   println(b)
     #end
     blens = [length(b) for b in basins]
     blens = reverse(sort(blens))
     println("Sorted basin lengths: ", blens)
     result = blens[1] * blens[2] * blens[3] 
     println("Part 2 result = $result")
end

# For part 2, find basins iteratively, i.e., areas around -1 bounded by
# 9s or edges. Looks for every -1, and expands in all four directions,
# assigning each point to the basin of the starting point. 
# Returns true if there were changes, so we know when to stop iterating.
function findBasins()

    changed = false  # if something changed, will need to iterate further
    for ri in 1:length(m)
        r = m[ri]
        for ci in 1:length(r)

            # Skip if this point is not within a basin
            if m[ri][ci] != -1
                continue
            end

            # Find the basin to which this point belongs
            p = [ri, ci]
            basin = nothing
            for b in basins
                if in(p, b)
                    basin = b
                    break
                end
            end

            # If no basin was found, something is wrong,
            # so stop here (does not happen with data provided)
            if basin == nothing
                println("No basin found for", p)
                return false  # to stop iterations
            end

            # Check points to the left/right/up/down of this point,
            # and mark/save any points until we reach a boundary,
            # i.e., edge of matrix or a 9

            # Check to right of this point
            i = ci + 1
            while i <= length(r) && r[i] != 9 && r[i] != -1
                r[i] = -1  # mark point as -1
                push!(basin, [ri, i])
                i += 1
                changed = true
            end

            # Check to left of this point
            i = ci - 1
            while i >= 1 && r[i] != 9 && r[i] != -1
                r[i] = -1
                push!(basin, [ri, i])
                i -= 1
                changed = true
            end
    
            # Check above this point
            i = ri - 1
            while i >= 1 && m[i][ci] != 9 && m[i][ci] != -1
                m[i][ci] = -1
                push!(basin, [i, ci])
                i -= 1
                changed = true
            end

            # Check below this point
            i = ri + 1
            while i <= length(m) && m[i][ci] != 9 && m[i][ci] != -1
                m[i][ci] = -1
                push!(basin, [i,  ci])
                i += 1
                changed = true
            end
        end
    end

    # Return status indicates whether something changed, so will need
    # to iterate further until there are no more changes.
    return changed

end

# Execute both parts (need to do part 1 in order to do part 2)
part1()
part2()

