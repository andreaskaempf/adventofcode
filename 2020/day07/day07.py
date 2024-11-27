# Advent of Code 2020, Day 7
#
# Recursively navigate a tree upwards (Part 1) and downwards (Part 2), to count
# how many bags contain a given type of bag, and how many bags a given type of
# bag ultimately contains.
#
# AK, 01/01/2022

# Read [outer, inner] rules into dictionary, duplicate where multiple
f = 'sample2.txt'
f = 'sample1.txt'
f = 'input.old'
f = 'input.txt'
bags = []  # list of [outer, inner] with duplicates
for l in open(f):

    # Parse left hand side
    l = l.strip().replace('.', '')
    outer = l[:l.find('bags')].strip()

    # Parse comma-separated list of counts and bags in right hand side (after
    # the word 'contain'), and create a list entry for each instance
    inner = l[l.find('contain')+8:].replace('bags', '').replace('bag', '').strip()
    if inner == 'no other':         # Don't bother if 'no other'
        continue
    for b in inner.split(','):      # Each count/type of bag contained
        b = b.strip()
        n = int(b[:b.find(' ')])    # The count
        b = b[b.find(' ')+1:]       # The bag type
        for i in range(n):          # Create a pair for each count
            bags.append([outer, b])

# Part 1: What bags can contain a "shiny gold" bag?
def contains1(bagtype):
    result = []
    for b in bags:
        if bagtype in b[1]:
            result.append(b[0])
            result += contains1(b[0])
    return result

ans = set(contains1("shiny gold"))
print('Part 1:', ans, len(ans))

# Part 2: how many bags must a "shiny gold" bag contain?
def contains2(bagtype):
    count = 0
    for b in bags:
        if b[0] == bagtype:
            count += 1
            count += contains2(b[1])
    return count


print('Part 2:', contains2("shiny gold"))
