# Advent of Code 2022, Day 13
#
# Given pairs of nested lists of numbers, count up how many are
# in the right order according to an arcane comparison function
# (part 1), then combine all the pair elements into one big list,
# add a couple of marker elements, and sort the list according
# to the comparison function. For Part 2, report the product of
# the indices of the two marker elements.
#
# AK, 13 Dec 2022

# Read file, split into pairs
fname = 'input.txt'
pairs = [p.split('\n') for p in open(fname).read().split('\n\n')]
pairs = [p for p in pairs if len(p) == 2]

# Comparison function that embodies logic in problem
def compare(l, r):
    
    # If both values are integers, the lower integer should come first. 
    # If the left integer is lower than the right integer, the inputs are 
    # in the right order. If the left integer is higher than the right 
    # integer, the inputs are not in the right order. Otherwise, the inputs 
    # are the same integer; continue checking the next part of the input.
    #print(l, 'vs', r, type(l), type(r))
    if type(l) == int and type(r) == int:
        if l < r:
            return True
        if l > r:
            return False
        return None # Undecided

    # If both values are lists, compare the first value of each list, 
    # then the second value, and so on. 
    # If the left list runs out of items first, the inputs are in the right order. 
    # If the right list runs out of items first, the inputs are NOT in the right order. 
    # If the lists are the same length and no comparison makes a decision 
    # about the order, continue checking the next part of the input.
    elif type(l) == list and type(r) == list:
        i = 0
        lastComp = None
        while i < max(len(l), len(r)):
            if i >= len(l): 
                return True
            if i >= len(r): 
                return False
            lastComp = compare(l[i], r[i])
            if lastComp != None:
                return lastComp
            i += 1
        return lastComp
    
    # If exactly one value is an integer, convert the integer to a 
    # list which contains that integer as its only value, then retry 
    # the comparison. For example, if comparing [0,0,0] and 2, convert 
    # the right value to [2] (a list containing 2); the result is then 
    # found by instead comparing [0,0,0] and [2].
    else:
        if type(l) == int:
            l = [l]
        if type(r) == int:
            r = [r]
        return compare(l, r)

# For part 1, sum up the 1-based indices of all pairs that
# are in correct order
ans = 0
for i in range(len(pairs)):
    pair = pairs[i]
    l, r = [eval(x) for x in pair]
    if compare(l, r):
        ans += i+1
print('Part 1 (16, 6484):', ans)

# For part 2, put all packets in the right order, first appending two "divider
# packets"

# Convert pairs to a list
bigList = []
for p in pairs:
    bigList.append(eval(p[0]))
    bigList.append(eval(p[1]))
    
# Add the two divider packets
bigList.append([[2]])
bigList.append([[6]])

# Now sort the list, using custom sort key
from functools import cmp_to_key

def myCmp(a, b):
    return 1 if compare(a, b) else -1

bigList.sort(key=cmp_to_key(myCmp), reverse = True)

# Part 2 answer is the product of the one-based indices of the two markers
i1 = bigList.index([[2]])+1
i2 = bigList.index([[6]])+1
print("Part 2 (s/b 140):", i1 * i2)
