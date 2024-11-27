# Advent of Code 2020, Day 1
#
# Read each line, and validate password against character-count 
# rules, e.g.,
#
# 1-3 a: abcde  (a must appear 1-3 times, so valid)
# 1-3 b: cdefg
# 2-9 c: ccccccccc
#
# AK, 17 Dec 2021

valid1 = valid2 = 0

#for l in open('sample.txt'):
for l in open('input.txt'):

    # Parse input line
    lims = l[:l.find(' ')]
    lims = [int(i) for i in lims.split('-')]
    letter = l[l.find(' ')+1:l.find(':')]
    pw = l[l.find(':')+1:].strip()

    # Part 1: Count the designated letter in the password,
    # check in range
    n = 0
    for c in pw:
        if c == letter:
            n += 1
    ok1 = n >= lims[0] and n <= lims[1]
    if ok1:
        valid1 += 1

    # Part 2: letter is in those two positions
    #ok2 = len(pw) >= lims[1] and pw[lims[0] - 1] == c and pw[lims[1] - 1] == c
    ok2 = 1 if pw[lims[0]-1] == letter else 0
    ok2 += 1 if pw[lims[1]-1] == letter else 0
    ok2 = ok2 == 1
    if ok2:
        valid2 += 1

    # Show this result
    print(lims, letter, pw, ok1, ok2)

print('Part 1:', valid1)
print('Part 2:', valid2)
