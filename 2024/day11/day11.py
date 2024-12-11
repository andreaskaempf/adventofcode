# Advent of Code 2024, Day 11, Python solution because of map 
# problem in Go

from collections import defaultdict
from copy import copy

data = [125, 17] # sample
#data = [5910927, 0, 1, 47, 261223, 94788, 545, 7771]
iters = 25 # 25 for Part 1, 75 for Part 2

# Convert list of stones to a frequency dict
stones = defaultdict(int)
for n in data:
    stones[n] += 1
print(stones)

# Perform "blinks":
for i in range(iters):

    # Get list of keys, i.e., the stone values, since we can't update
    # dictionary while directly iterating over its keys
    nn = list(stones.keys())

    # Make changes on a deep copy, not the original list
    d2 = defaultdict(int)
    for k, v in stones.items():
        if v > 0:
            d2[k] = v

    # Transform each stone, updating counters
    for n in nn:

        # Skip if there are none of this stone
        if stones[n] == 0:
            continue

        nstones = stones[n]
        s = str(n)
        if n == 0:              # 0 becomes 1
            d2[0] -= nstones
            d2[1] += nstones
        elif len(s) % 2 == 0:   # even digits -> split in half
            d2[n] -= nstones
            cut = len(s)//2
            l = int(s[:cut])
            r = int(s[cut:])
            d2[l] += nstones
            d2[r] += nstones
        else:                   # otherwise multiply by 2024
            d2[n] -= nstones
            d2[n * 2024] += nstones

    stones = d2
    
    # Part 1: should be 193607 after 25 iterations
    # Part 2:  229557103025807 after 75 iterations
    print(i, sum(stones.values()))

