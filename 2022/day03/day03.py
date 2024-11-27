# Advent of Code 2022, Day 3 (Python solution)
#
# AK, 3/12/2022

filename = 'sample.txt'
filename = 'input.txt'

# Character value of a letter (according to problem definition rules)
def cval(c):
    if c >= 'a' and c <= 'z':
        return ord(c) - ord('a') + 1
    else:
        return ord(c) - ord('A') + 27

# Part 1: sum of values of common character in both halves of each string
tot = 0
lines = [l.strip() for l in open(filename)]
for l in lines:
    mid = int(len(l) / 2)
    h1 = l[:mid]
    h2 = l[mid:]
    common = list(set(h1).intersection(h2))
    tot += cval(common[0])
print('Part 1:', tot)

# Part 2: sum of common chars in each group of 3 lines
i = tot = 0
while i < len(lines):
    l1, l2, l3 = lines[i:i+3]
    c = set(l1).intersection(set(l2)).intersection(set(l3))
    tot += cval(list(c)[0])
    i += 3
print('Part 2:', tot)
