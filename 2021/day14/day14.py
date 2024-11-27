# Advent of Code 2021, Day 14
#
# Apply character insertion rules to a sequence of characters,
# runs out of memory if you try to build up character strings,
# so had to build dictionary of pairs of characters.
#
# AK, 14/12/2021

import time
t0 = time.time()

# Input file name
f = 'sample.txt'
f = 'input.txt'

# Read data, pattern on line 1, rules thereafter
lines = [l.strip() for l in open(f)]
patt = None
rules = {}
for l in lines:
    if not patt:
        patt = l
    elif len(l) > 0:  # "AB -> C" to dictionary
        rules[l[:2]] = l[6]

# Parse starting pattern, and get frequency counts of letters,
# and number of transitions
trans = {}  # "AB" -> count
counts = {} # 'A' -> count
prevC = None
for c in patt:
    counts[c] = counts.get(c,0) + 1
    if prevC:
        t = prevC + c
        trans[t] = trans.get(t,0) + 1
    prevC = c

# Show starting data
print('Transitions:', trans)
print('Chars:', counts)
print('Rules:', rules)

# Do one iteration
def iter():

    global trans, counts

    # Look at each transition pair in pattern
    insertions = {}   # list of chars to insert inside each pair
    for pair in trans.keys():

        # Skip if no rule for this transition
        if not pair in rules:
            print('No rule for:', pair)
            continue

        # We add one letter for each time this pair appears in pattern
        c = rules[pair]
        counts[c] = counts.get(c,0) + trans[pair]

        # Get the char to insert between the pair, add to list of insertions
        # for this pair
        if not pair in insertions:
            insertions[pair] = []
        insertions[pair].append(c)

    # Now recalculate the transitions for future iterations from the
    # original pairs and the chars inserted
    trans2 = {}
    for pair in trans.keys():

        # Retain as-is if no insertions
        if not pair in insertions:
            trans2[pair] = trans[pair]
            continue

        # Transform the original pair into new transitions based on
        # inserted characters
        pins = insertions[pair]
        for i in range(len(pins)):
            if i == 0:
                p = pair[0] + pins[i]
            else:
                p = pins[i-1] + pins[i]
            trans2[p] = trans2.get(p,0) + trans[pair]

        # Last transition
        p = pins[-1] + pair[1]
        trans2[p] = trans2.get(p,0) + trans[pair]

    # This is the new transition  list
    trans = trans2

# Do iterations (10 for Part 1, 40 for Part 2)
for i in range(40):
    print('\nIteration', i+1)
    iter()
    print('Counts:', counts)

# Show final results
print('\nFinal character counts:', counts)
print('\nMax - min counts:', max(counts.values()) - min(counts.values()))
print(time.time() - t0, 'secs')
