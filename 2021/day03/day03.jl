# Day 3 of Advent of Code
#
# AK, 3/12/2021

# Read text file, list of strings
data = readlines(open("input.txt"))
nrows = length(data)
ncols = length(data[1])

# Convert to array of arrays of integers, one per row
# To Matrix: m = transpose(hcat(vv...))
rows = map(r -> map(x -> parse(Int64, x), collect(r)), data)

# Convert vector of 1/0 binary to decimal number
function b2i(v)
    n = length(v)
    return sum(map(i -> v[i] * 2^(n - i), 1:n))
end

# Function to add up the ones in each column
function countOnes(rr)
    return map(c -> sum(map(r -> r[c], rr)), 1:ncols)
end

### PART 1 ###

# Add up each column in full data set to get most/least frequent value
sums = countOnes(rows)

# Gamma and Epsilon: 1/0 depending on most frequent (Epsilon opposite)
gamma = map(i -> i >= nrows / 2 ? 1 : 0, sums)
epsilon = map(i -> i >= nrows / 2 ? 0 : 1, sums)

# Convert both to decimal, and show result
g = b2i(gamma)
e = b2i(epsilon)
println("Gamma: ", gamma, " = ", g)
println("Epsilon: ", epsilon, " = ", e)
println("Product: ", g * e)

### PART 2 ###

# Go through numbers, start at first column, keep only numbers
# that meet bit criteria, until only one left:
# - oxygen: most common value (keep 1 if equally common)
# - co2: least common value (keep 0 if equally common)

println("\nPART 2")

# Filter list to numbers that have most frequent 1/0 in
# the current column
println("\nOxygen:")
nn = copy(rows)
for c in 1:ncols
    counts = countOnes(nn)
    freq1 = counts[c] / length(nn)
    bit = freq1 >= .5 ? 1 : 0
    global nn = filter(r -> r[c] == bit, nn)
    println("  Col $c: $(length(nn)) left")
    if length(nn) <= 1   # Only 1 left, so stop
        break
    end
end

ox = b2i(nn[1])
println("Oxygen: $nn = $ox")

# CO2: Filter list to numbers that have least frequent 1/0 in
# the current column
println("\nCO2:")
nn = copy(rows)
for c in 1:ncols
    counts = countOnes(nn)
    freq1 = counts[c] / length(nn)
    bit = freq1 >= .5 ? 0 : 1
    global nn = filter(r -> r[c] == bit, nn)
    println("  Col $c: $(length(nn)) left")
    if length(nn) <= 1   # Only 1 left, so stop
        break
    end
end

co2 = b2i(nn[1])
println("CO2: $nn = $co2")

# Show final result
println("\nProduct = $(ox * co2)")
