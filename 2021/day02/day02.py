# Advent of Code, Day 2

# State variables for problems 1 & 2
horiz1 = depth1 = 0
horiz2 = depth2 = aim2 = 0

# Process each line of text
for l in open('input.txt'):

    # Get instruction and value
    instruction, n = l.split()
    n = int(n)

    # Process instructions for both problems
    if instruction == 'forward':
        horiz1 += n
        horiz2 += n
        depth2 += aim2 * n
    elif instruction == 'down':
        depth1 += n
        aim2 += n
    elif instruction == 'up':
        depth1 -= n
        aim2 -= n
    else:
        print('Bad instruction:', instruction)

# Show results for both problems
print(f'Problem 1: Ending pos = {horiz1}, depth = {depth1}, product = {horiz1 * depth1}')
print(f'Problem 2: Ending pos = {horiz2}, depth = {depth2}, aim = {aim2}, product = {horiz2 * depth2}')

