# Advent of Code, Day 12
#
# Find the number of paths from start to end in a graph, allowing repeat
# visits for "large caves" (capitalized), but either no (part 1) or at
# most 2 (part 2) visits for "large caves" (lower case).
#
# AK, 12 Dec 2021

part1 = false       # Set to true for Part 2
paths = []          # List of solutions found
f = "input.txt"     # Input file name

# Try all possible trajectories from the current location
function traverse(path::Vector{String}, node::String)

    # Add this node to a copy of the list of visited nodes
    path = copy(path)
    push!(path, node) 

    # If this path ends with "end", add it to solutions
    if node == "end"
        push!(paths, join(path, ','))
        return
    end

    # Otherwise, recursively try each potential next node
    for dest in dests[node]
        if legal(path, dest)
            traverse(path, dest)
        end
    end

end

# Determine if it is legal to visit from the current path to the given location
function legal(path::Vector{String}, node::String) 

    # For part 1, this is just uppercase, or not already visited
    if part1
        return !(isSmall(node) && node in path)
    end

    # Can't go back to start
    if node == "start" 
        return false
    end

    # Can't allow if any duplicate small caves
    if dupes(filter(isSmall, path))
        return !(isSmall(node) && node in path)
    end

    # Otherwise, okay
    return true 

end

# Check whether there are any duplicates in a list
function dupes(l::Vector{String})::Bool
    ctr = Dict{String,Int}()
    for s in l
        if haskey(ctr, s)
            return true
        end
        ctr[s] = 1
    end
    return false
end

# Check whether a name is all lower case
function isSmall(s::String)::Bool
    return s == lowercase(s)
end

# Read input file, accumulate dictionary of possible next destination 
# for any node, in either direction
dests = Dict{String,Vector{String}}()
lines = filter(s -> length(s) > 0, split(read(f, String), "\n"))
for l in lines
    s, d = split(l, '-')
    if ! haskey(dests, s)
        dests[s] = []
    end
    if ! haskey(dests, d)
        dests[d] = []
    end
    push!(dests[s], d)
    push!(dests[d], s)
end

# Start traversing from the start node
emptyPath = Vector{String}()
traverse(emptyPath, "start")

# Show number of paths found
println("Number of paths: ", length(paths))

