# Advent of Code 2020, Day 10
#
# Given a sequence of numbers, count up the number of 1- and 3- gaps
# within the list. Then, count up all the possible paths between the
# lowest and highest values, enforcing the constraint that the gap between
# any two number cannot be more than 3. Exhaustive recursive approach was
# too slow, but extremely fast if you cache interim results (7 ms in Python).
#
# AK, 04/01/2022

# Read rows and sort them (for Part 1)
f = 'sample2.txt'
f = 'sample1.txt'
f = 'input.txt'
d = [int(d) for d in open(f).readlines()]
d += [0, max(d) + 3]  # Add power outlet zero and my devices power

# Part 1: Choose the sequence of *all* adapters that starts from zero, and goes
# all the way up to top Joltage. Adapters can only connect to a source 1-3
# jolts lower than its rating. What is the number of 1-jolt differences
# multiplied by the number of 3-jolt differences?
#
# This can be done by simply sorting the list, then counting up the
# gaps between consecutive values
#
# Sample 1: 7 x 1 jolt, 5 x 3 jolts = 35
# Sample 2: 22 x 1 jolt, 10 x 3 jolts = 220
# Input data: 69 x 1, 33 x 3 = 2277

d = sorted(d)
diffs = {}
for i in range(1, len(d)):
    diff = d[i] - d[i-1]
    diffs[diff] = diffs.get(diff, 0) + 1

print('Part 1:', diffs, diffs[1] * diffs[3])

# Part 2: What is the total number of distinct ways you can arrange the
# adapters to connect the charging outlet to your device?
# Each adapter can connect to an adapter 1-3 more than itself
# Sample 1: 8
# Sample 2: 19208
# Input: 37024595836928 (!)
#
# Found it was too slow to recursively enumerate all paths from the start.
# Instead, cache (memoize) number of paths going forward from any point we
# encounter, and use this cached value instead of repeatedly traversing the
# same paths over and over again.

# For keeping track of the number of paths forward from any node
pathsfromhere = {}

# Count up the number of paths to the final point, from any  node
def paths2(a):

    # If already cached, use that value
    global pathsfromhere
    if a in pathsfromhere:
        return pathsfromhere[a]

    # Otherwise, find all the nodes we could go to from here,
    # and return 1 if there are none
    nodesfromhere = list(filter(lambda x: x > a and x <= a+3, d))
    if len(nodesfromhere) == 0:
        pathsfromhere[a] = 0
        return 1

    # Otherwise, recursively get the number of paths from this node
    n = 0
    for p in nodesfromhere:
        n += paths2(p)

    # Cache the value before returning it
    pathsfromhere[a] = n
    return n

# Compute the total number of paths from 0 to the end
total = paths2(0)
print('Part 2:', total)
