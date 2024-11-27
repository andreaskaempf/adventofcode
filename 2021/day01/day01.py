# Python version of Day 1 exercise, count number of
# increases in a text file

# Read numbers into a list
nn = [int(x) for x in open('input.txt')]

# Part 1: count up number of increase from previous value
increases = 0
for i in range(1, len(nn)):
    if nn[i] > nn[i - 1]:
        increases += 1
print(f'Part 1: {increases} increases found')

# Part 2: count up sliding window of 3
prev = increases = 0
for i in range(3, len(nn)+1):
    l = nn[i-3:i]
    if prev > 0 and sum(l) > prev:
        increases += 1
    prev = sum(l)
print(f'Part 2: {increases} increases found')

