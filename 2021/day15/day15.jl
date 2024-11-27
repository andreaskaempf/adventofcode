# Advent of Code 2021, Day 15
#
# Objective is to start at the top right, and end up at 
# the bottom right, adding up any cells you enter, so as to
# minimize total cost.
#
# First attempt depth-first search with maximum look-ahead,
# but did not succeed on full data set. New solution uses 
# Dijkstra's algorithm, adapted from:
# https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
# Note that the numbered steps in the comments are copied
# from this Wikipedia article.
#
# AK, 15/12/2021

# Name of input file
f = "sample.txt"
f = "input.txt"

# Read data into a matrix
d = split(strip(read(f, String)), "\n")
l = map(r -> [x - '0' for x in r], d)
m = Matrix(transpose(reshape(vcat(l...), (:, length(l)))))
#display(m)

# Use this value to represent extremely high cost
PENALTY = 99999999

# Find the lowest total cost path from 1,1 to m,n of a matrix
function solve(M::Matrix{Int64})

    # 1. Mark all nodes as unvisited
    nr, nc = size(M)
    visited = zeros(nr, nc)

    # 2. Assign to every node a tentative distance value: set it to zero for our
    # initial node and to infinity for all other nodes. The tentative distance of a
    # node v is the length of the shortest path discovered so far between the node
    # v and the starting node. Since initially no path is known to any other vertex
    # than the source itself (which is a path of length zero), all other tentative
    # distances are initially set to infinity. Set the initial node as current.[15]
    distances = fill(PENALTY, nr, nc)
    distances[1,1] = 0

    # Start algoritm iterations, will stop when reach bottom right
    iter = 0
    r, c = 1, 1     # Initial node: top left
    while ! (r == nr && c == nc)

        # 3. For the current node, consider all of its unvisited neighbors and
        # calculate their tentative distances through the current node. Compare the
        # newly calculated tentative distance to the current assigned value and assign
        # the smaller one. For example, if the current node A is marked with a distance
        # of 6, and the edge connecting it with a neighbor B has length 2, then the
        # distance to B through A will be 6 + 2 = 8. If B was previously marked with a
        # distance greater than 8 then change it to 8. Otherwise, the current value
        # will be kept.
        for p in [[r, c-1], [r, c+1], [r-1,c], [r+1,c]] # left, right, up, down
            y, x = p
            if y >= 1 && y <= nr && x >= 1 && x <= nc && visited[y, x] == 0
                c1 = M[y, x] + distances[r, c]
                if c1 < distances[y, x]
                    distances[y, x] = c1
                end
            end
        end

        # 4. When we are done considering all of the unvisited neighbors of the current
        # node, mark the current node as visited and remove it from the unvisited set.
        # A visited node will never be checked again.
        visited[r, c] = 1

        # 5. If the destination node has been marked visited (when planning a route
        # between two specific nodes) or if the smallest tentative distance among the
        # nodes in the unvisited set is infinity (when planning a complete traversal;
        # occurs when there is no connection between the initial node and remaining
        # unvisited nodes), then stop. The algorithm has finished.
        #
        # 6. Otherwise, select the unvisited node that is marked with the smallest
        # tentative distance, set it as the new current node, and go back to step 3.
        #
        # When planning a route, it is actually not necessary to wait until the
        # destination node is "visited" as above: the algorithm can stop once the
        # destination node has the smallest tentative distance among all "unvisited"
        # nodes (and thus could be selected as the next "current").
        #
        # In this case, get the coordinates of the minimum value in the matrix, first
        # changing visited cells to high-cost penalty so they are not considered.
        nextCosts = distances .+ visited * PENALTY
        r, c = argmin(nextCosts).I   # coords of min cost (next iter starts there)

        # Report progress
        iter += 1
        if iter % 1000 == 0
            print("\rIteration $iter: now at $r,$c   ")
        end

    end

    # Return solution, cost set in the distances matrix for bottom right point
    return distances[nr, nc]

end

# For part 2, make a 5x5 copy of the matrix, each row/col being one higher
# than the previous, with 9s wrapped back to 1
function enlarge(x, n)

    # To grow 5x, add 4 copies either direction
    n -= 1

    # Grow horizontally
    m1 = copy(x)
    for i in 1:n
        m1 = hcat(m1, x .+ i)
    end

    # Grow vertically
    m2 = copy(m1)
    for i in 1:n
        m2 = vcat(m2, m1 .+ i)
    end

    # Adjust cells that are bigger than 9
    nr, nc = size(m2)
    for r in 1:nr
        for c in 1:nc
            while m2[r,c] > 9
                m2[r,c] -= 9
            end
        end
    end

    # Return result
    return m2
end

# Part 1 solution
s1 = solve(m)
println("\nPart 1: ", s1)

# Part 2, same but on enlarged matrix
mbig = enlarge(m, 5)
s2 = solve(mbig)
println("\nPart 2: ", s2)





