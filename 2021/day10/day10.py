# Advent of Code 2021, Day 10
#
# Identify and resolve different types of mis-matched brackets in a file of
# strings.
#
# AK, 10 Dec 2021

# Name of input file
f = 'test_input.txt'
f = 'input.txt'

### Part 1: score and discard corrupt lines ###

# Closing -> Opening bracket of each type
opener = {')': '(', ']':'[', '}':'{', '>':'<'}

# "Score" for each illegal bracket found
scores =  {')': 3, ']':57, '}':1197, '>':25137}

# Check if a line is corrupt, return score if it is
def part1check(l):
    state = ''  # Starting state
    for c in l:
        if c in opener.values():
            state += c
        elif len(state) == 0 or state[-1] != opener.get(c,'?'):
            return scores[c]
        else:
            state = state[:-1]
    return 0

# Part 1: go through, score any corrupt lines and discard them
part1_score = 0
ok = []  # non-corrupt lines, for part 2
for l in [l.strip() for l in open(f)]:
    lscore = part1check(l)
    part1_score += lscore
    if lscore == 0:
        ok.append(l)

print('Part 1:', part1_score)

### Part 2: complete non-complete lines ###

# Create inverse of closer dictionary
closer = {}
for c in opener.keys():
    closer[opener[c]] = c

# Points allocated for every completing closer
points = {')': 1, ']':2, '}':3,'>':4}

print(len(ok), 'non-corrupt lines left')
scores = []
for l in ok:

    # Go through line and accumulate state for each bracket type
    state = []
    for c in l: 
        if c in opener.values():
            state.append(c)
        else:
            state.pop()

    # To complete the line, we need to add a closing bracket
    # for each dangling opening bracket in the state, in reverse
    # order
    complete = [closer[c] for c in state]
    complete.reverse()

    # Create score for this line's completion
    score = 0
    for c in complete:
        score = score * 5 + points[c]
    scores.append(score)

# Result is the middle point of the sorted scores
scores = sorted(scores)
result = scores[int(len(scores) / 2)]
print('Part 2:', result)

