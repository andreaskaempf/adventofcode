# Day 8 of Advent of Code
#
# Decipher scrambled digits on a digital display
#
# AK, 8 Dec 2021

# You need to install this
using Combinatorics

# Read input text, each line consisting of a list of inputs, a bar, then a list
# of outputs (first edit input to remove any double spaces)
#raw_data = read("test_input.txt", String)
#raw_data = read("test_input2.txt", String)
raw_data = read("input.txt", String)
raw_data = filter(s -> length(s) > 0, split(raw_data, "\n"))
splits = map(s -> split(s, " | "), raw_data)
inputs  = filter(s -> length(s) > 0, map(r -> map(strip, split(r[1], " ")), splits))
outputs = filter(s -> length(s) > 0, map(r -> map(strip, split(r[2], " ")), splits))
#println("Inputs: ", inputs)
#println("Outputs: ", outputs)

# Easy digits:
# 1 -> 2 segments (c and f)
# 4 -> 4 segments (b, c, d and f)
# 7 -> 3 segments (a, c and f)
# 8 -> 7 segments (a-g)
#
# Harder:
# 0 -> 6 segments (abcefg
# 2 -> 5 segments (a, c, d, e and g)
# 3 -> 5 segments (a, c d f g)
# 5 -> 5 segments (a b d f g)
# 6 -> 6 segments (abdefg)
# 9 -> 6 segments (abcdfg)
#
# So 5 segs -> 2, 3, or 5
#    6 segs -> 0, 6, or 9
#
# 10 inputs = signal patterns
# 4 outputs = four digits output

### PART 1 ###

# Part 1: count up how many times digits 1, 4, 7, 8 appear in output
# i.e., how many outputs have length 2, 3, 4 or 7 
function part1()
    n = 0
    for o in outputs
        for d in o
            if length(d) == 2 || length(d) == 3 || length(d) == 4 || length(d) == 7
                n += 1
            end
        end
    end
    println("Part 1: digits 1, 4, 7, 8 appear $n times")
end

println("Part 1:")
part1()

### PART 2 ###

function part2()

    # Analyse each input row
    tot = 0
    for i in 1:length(inputs)
        ii = map(string, inputs[i])
        oo = map(string, outputs[i])
        println("\nAnalysing ", ii, " => ", oo)
        tot += analyse(ii, oo)
    end
    println("\nTotal = $tot, $errors errors occurred")

end

# Decoding of patterns to digits
digits = Dict{String,Int64}("abcefg" => 0, "cf" => 1, "acdeg" => 2,
    "acdfg" => 3, "bcdf" => 4, "abdfg" => 5, "abdefg" => 6,
    "acf" => 7, "abcdefg" => 8, "abcdfg" => 9)

println("Digits: ", digits)

# Global error counter
errors = 0

# Analyse inputs/outputs from one row: figure out which input
# pins control which segments, then use this knowledge to decipher
# the 4-digit output value
function analyse(ii, oo)

    # Come up with every possible mapping of i->o for i and o both in a-g,
    perms = permutations(['a', 'b', 'c', 'd', 'e', 'f', 'g'], 7)
    println(length(perms), " permutations to test")

    # Test each permutation to find one that matches
    good_mapping = Dict{Char,Char}()
    nmatches = 0  # number of permutations that "work"
    for perm in perms

        # Turn the permutation into a dictionary
        #println("Permutation = ", perm)
        mapping = Dict{Char,Char}()
        for ci in 1:length(perm)  # each character
            mapping['a' + ci - 1] = perm[ci]
        end
        #println("Mapping = ", mapping)

        # Go through each input and check if it makes sense
        # against this mapping
        perm_ok = true      # Start by assuming it's okay
        for inp in ii       # Look at each input code on this line

            # Unscramble this input using the mapping we are trying
            mapped = join(map(c -> mapping[c], inp))

            # Test if it's a valid combination given the length, i.e.,
            # the given number of letters match the expected LED segments
            ok = true
            if length(mapped) == 2
                if ! matches(mapped, "cf")
                    ok = false
                end
            elseif length(mapped) == 3
                if ! matches(mapped, "acf")
                    ok = false
                end
            elseif length(mapped) == 4
                if ! matches(mapped, "bcdf")
                    ok = false
                end
            elseif length(mapped) == 7
                if ! matches(mapped, "abcdefg")
                    ok = false
                end
            elseif length(mapped) == 5
                if ! (matches(mapped, "acdeg") || matches(mapped, "acdfg") || matches(mapped, "abdfg"))
                    ok = false
                end
            elseif length(mapped) == 6
                if ! (matches(mapped, "abdefg") || matches(mapped, "abcdfg") || matches(mapped, "abcefg"))
                    ok = false
                end
            end

            # Set flag to indicate that this mapping is no good, if this input didn't match
            if ! ok
                perm_ok = false
            end
        end

        # If the permutation worked for all inputs, that's good, but update
        # counter to detect so we can check if there are duplicates (i.e., if
        # more than one permutation works for the input -- did not happen in
        # either test or problem data)
        if perm_ok
            nmatches += 1
            good_mapping = mapping
        end
    end

    # There should be exactly one match: if there are no matches, 
    # or more than one, return a score of zero
    if nmatches != 1
        println("ERROR: $nmatches permutations match!")
        global errors += 1
        return 0
    end

    # Otherwise, proceed to use the matching permutation to decode the outputs
    println("Decoding using unique mapping: ", good_mapping)
    result = ""  # Will append decoded digits to this
    for o in oo

        # Decode the input signal
        decoded = join(sort([good_mapping[c] for c in o]))

        # Find the digit that matches this pattern
        digit = get(digits, decoded, -1)

        # Show decoded digit, or error if not found (does not happen)
        if digit == -1
            println(o, " => ", decoded, " = INVALID!")
            global errors += 1
            digit = 0
        else
            println(o, " => ", decoded, " = ", digit)
        end

        # Append digit to the result
        result = result * string(digit)  # append digit
    end

    # Convert result from string to a number and return it
    result = parse(Int64, result)
    println("Result = ", result)
    return(result)

end

# See if two strings match, even if chars are in different order
# To do: could just compare sorted strings
function matches(s1::String, s2::String)::Bool
    return Set([c for c in s1]) == Set([c for c in s2])
end

println("\nPart 2:")
part2()

