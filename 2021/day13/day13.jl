# Day 13 of Advent of Code 2021
#
# Repeatedly "fold" a matrix of 1/0 values along horizontal
# or vertical lines, and (part 1) count up the number of points
# remaining after one fold, and (part 2) visualize the matrix
# after all folds to yield a dot matrix alphabetic code.
#
# AK, 13/12/2021

infile = "test_input.txt"
infile = "input.txt"

# Read data into a sparse matrix, i.e., a dict of (x,y) pairs
points = Dict()
folds = []
for l in readlines(infile)
    if length(l) == 0
        continue
    end
    if startswith(l, "fold")
        push!(folds, l)
    else
        p = map(n -> parse(Int64, n), split(l, ","))
        points[p] = 1
    end
end

# Process each fold (for part 1, only the first)
part1 = 0
for f in folds

    # Turn "fold along x=3" into "x" and 3
    f = split(f," ")[3]
    dir, n = split(f, "=")
    n = parse(Int64, n)
    
    # Process each point below/right of fold
    pp = keys(points)
    if dir == "y"
        for p in pp
            if p[2] > n
                newP = [p[1], 2 * n - p[2]]
                points[newP] = 1
                delete!(points, p) # Delete old point
            end
        end
    else   # fold along X
        for p in pp
            if p[1] > n
                newP = [2 * n - p[1], p[2]]
                points[newP] = 1
                delete!(points, p) # Delete old point
            end
        end
    end

    # For part 1, report number of points after first iteration
    if part1 == 0
        global part1 = sum(values(points))
    end

end

# Show part 1, total number of points remaining after first fold
println("Part 1: ", part1)

# Output for part 2: visualize matrix to get 8 capital letters
output = []  # Each row a list of characters
for p in keys(points)
    x, y = map(n -> n + 1, p)
    while length(output) < y
        push!(output, [])
    end
    while length(output[y]) < x
        push!(output[y], " ")
    end
    output[y][x] = "X"
end

# Show the matrix
println("\nPart 2:\n")
for l in output
    println(join(l))
end
