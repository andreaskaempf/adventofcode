# Advent of Code, Day 11
#
# Simulate energy growth of a field of "octopusses", with each
# exploding when they reach a level, infecting those around them.
#
# AK, 11 Dec 2021

# Parameters
f = "test_input.txt"
f = "input.txt"
iterations = 1000  # For part 2, need to go beyond 100

# Read data into list of rows, convert to matrix
d = split(strip(read(f, String)), "\n")
l = map(r -> [x - '0' for x in r], d)
m = Matrix(transpose(reshape(vcat(l...), (:, length(l)))))
println("Starting matrix:")
display(m)
nrows, ncols = size(m)

# Run the simulation
function simulate(m)

    # Start the simulation
    totFlashes = 0
    for t in 1:iterations

        # At the beginning of each step, add 1 to each cell
        m = m .+ 1

        # Initialize a matrix of flags indicating whether a cell
        # has already flashed this iteration
        notFlashed = ones(Int64, nrows, ncols)

        # Any cells with a value > 9 flash, incrementing the cells
        # around them, continue until no more flashes
        flashesFound = true
        while flashesFound
            flashesFound = false
            for r in 1:nrows
                for c in 1:ncols
                    if m[r,c] > 9 && notFlashed[r,c] == 1
                        flash(m, r, c)
                        notFlashed[r,c] = 0
                        flashesFound = true
                        totFlashes += 1
                    end
                end
            end
        end

        # Finally, set all the cells that have flashed to zero
        m = m .* notFlashed

        # Show totals and matrix this step
        println("\n\nAfter step $t, Total flashes = $totFlashes:")
        display(m)

        # For part 2, stop when all cells have flashed simulataneously
        if sum(notFlashed) == 0
            println("\nAll cells flashed at t = $t")
            break
        end
    end
end

# Flash one cell in the matrix: this increments all diagonal cells by one
function flash(m, r, c)
    for ri in r-1:r+1
        for ci in c-1:c+1
            if ri >= 1 && ri <= nrows && ci >= 1 && ci <= ncols
                m[ri,ci] += 1
            end
        end
    end
end        

# Run the simulation
simulate(m)
