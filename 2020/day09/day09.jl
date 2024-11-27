# Advent of Code 2020, Day 9
#
# Given a list of numbers, and a given sliding window width, find the first
# number that cannot be the sum of two digits in the preceding window. For Part
# 2, find any set of at least two contiguous numbers that add up to the number
# not found above, and return the sum of the minimum and maximum of this range.
#
# AK, 06/12/2021 (Part 1) and 01/01/2022 (Part 2)

# Test parameters
f = "sample.txt"
window = 5

# Real parameters
f = "input.txt"
window = 25

# Read data into a list of numbers
data = map(n -> parse(Int64, n), readlines(f))

# Function to see if the number provided is the sum of any two numbers in the slice
function found(n, nums)
    for i in 1:length(nums)
        for j in 1:length(nums)
            if i > j && nums[i] + nums[j] == n
                return true
            end
        end
    end
    return false
end

# Find the first number that cannot be the sum of two digits in the preceding window
function part1()
    notfound = 0
    for i in window+1:length(data)
        n = data[i]
        slice = data[i-window:i-1]
        #println("i = ", i, ", n = ", n, ", slice = ", slice)
        if ! found(n, slice)
            #println("Not found: ", data[i])
            notfound = data[i]
            break
        end
    end
    return notfound
end

nf = part1()
println("Part 1: ", nf, " was not found")

# Use this as a basis for part 2: find any set of at least two contiguous
# numbers that add up to the number not found above, and return the sum
# of the minimum and maximum of this range.
function part2(notfound)
    for i in 1:length(data)     # start of range
        tot = 0
        for j in i:length(data) # end of of range
            tot += data[j]
            if tot > notfound
                break
            elseif tot == notfound
                println("Found range ", data[i], " to ", data[j])
                return minimum(data[i:j]) + maximum(data[i:j])
            end
        end
    end
    return 0
end

ans = part2(nf)
println("Part 2: ", ans)
