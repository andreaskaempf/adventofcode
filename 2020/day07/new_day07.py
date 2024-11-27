# Advent of Code 2020, Day 7
#
# Recursively navigate a tree upwards (Part 1) and downwards (Part 2), to count
# how many bags contain a given type of bag, and how many bags a given type of
# bag ultimately contains.
#
# AK, 01/01/2022
# Revised 27/07/2022 because didn't work with new data set

# Read [outer, inner] rules into dictionary, duplicate where multiple
f = 'sample2.txt'
f = 'sample1.txt'
f = 'input.old'
f = 'input.txt'

# Dictionary of colours encountered, with list of colors within
bags = {}  # { "color" : ["color1", "color2", ...] }
for l in open(f):

    # Parse left hand side, colour name of the outer bag
    l = l.strip().replace('.', '')
    outer = l[:l.find('bags')].strip()
    if not outer in bags:
        bags[outer] = {}

    # Parse comma-separated list of counts and bags in right hand side (after
    # the word 'contain'), and create a list of inner count/colours
    inner = l[l.find('contain')+8:].replace('bags', '').replace('bag', '').strip()
    if inner == 'no other':         # Don't bother if 'no other'
        continue

    inner = [s.strip() for s in inner.split(',')]
    #inner = [i[i.find(' ')+1:].strip() for i in inner]  # drop leading number
    for i in inner:      # Each count/type of bag contained
        n = int(i[:i.find(' ')])      # the count
        c = i[i.find(' ')+1:].strip() # the colour
        bags[outer][c] = bags[outer].get(c,0) + n

# Does a bag of given colour contain a bag of the other colour,
# directly or indirectly?
def contains1(outercolor, innercolor):
    if innercolor in bags[outercolor].keys():
        return True
    for c in bags[outercolor].keys():
        if contains1(c, innercolor):
            return True
    return False

# Part 1: What bags can contain a "shiny gold" bag?
print('Part 1: bag colours that contain "shiny gold"')
n = 0
for c in bags.keys():  # every outer bag color
    if contains1(c, 'shiny gold'):
        n += 1
print(n, 'colors')

# Part 2: how many bags must a "shiny gold" bag contain?
def contains2(bagtype):
    bt = bags[bagtype]
    count = 0
    for c in bt.keys():
        count += bt[c] + bt[c] * contains2(c)
    return count

print('Part 2:', contains2("shiny gold"))
