# Advent of Code 2020, Day 1
#
# Read a list of  numbers, and find combinations of 2 or 3 that
# add up to 2020, and multiply them together.
#
# AK, 17 Dec 2021

# Read numbers from file
nums = [int(n) for n in open('input.txt').readlines()]

# Find two numbers that add up to 2020
print('Part 1:')
for i in nums:
    for j in nums:
        if i + j == 2020:
            print(i, j, i * j)
            break

# Same for 3 numbers
print('Part 2:')
for i in nums:
    for j in nums:
        for k in nums:
            if i + j +k == 2020:
                print(i, j, k, i * j * k)
                break

