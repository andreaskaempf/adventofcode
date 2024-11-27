# Advent of Code 2020, Day 3
#
# Count up the number of '#' encountered when traversing
# an array with a given dx/dy slope for each step. The
# rows wrap if x coordinate is wider than the data.

# Read rows of data
f = 'sample.txt'
f = 'input.txt'
data = [l.strip() for l in open(f).readlines()]

# Count up how many trees (#) we encounter going from top
# right (0,0) to the last row, given slope; data wraps
# horizontally.
def count(dx, dy):
    r = c = trees = 0
    rowlen = len(data[0])
    while r < len(data):
        p = data[r][c % len(data[r])]
        if p == '#':
            trees += 1
        c += dx
        r += dy
    return trees

print('Part 1:', count(3,1), 'trees encountered')

ans2 = count(1,1) * count(3,1) * count(5,1)* count(7,1)* count(1,2)
print('Part 2:', ans2, 'trees encountered')
