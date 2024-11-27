# Day 3 of Advent of Code
#
# AK, 3/12/2021

# Read binary strings into a list
nums = [l.strip() for l in open('input.txt')]
nrows = len(nums)
ncols = len(nums[0])

# Function to count up 1s in each column
def countOnes(rows):
    ones = []
    for c in range(ncols):
        ones.append(0)
        for n in rows:
            if n[c] == '1':
                ones[-1] += 1
    return ones

### PART 1 ###

# Count up 1s in each column
ones = countOnes(nums)

# Get digits of gamma and epsilon most/least freq respectively)
gamma = ''
epsilon = ''
for c in range(ncols):
    if ones[c] >= nrows/2:
        gamma += '1'
        epsilon += '0'
    else:
        gamma += '0'
        epsilon += '1'

# Calculate product
print('*** Part 1\n')
g = int(gamma,2)
e = int(epsilon,2)
print(f'Gamma: {gamma} = {g}')
print(f'Epsilon: {epsilon} = {e}')
print('Product:', g * e)

### PART 2 ###

# Go through numbers, start at first column, keep only numbers
# that meet bit criteria, until only one left:
# - oxygen: most common value (keep 1 if equally common)
# - co2: least common value (keep 0 if equally common)

print('\n*** Part 2')

# Filter list to numbers that have most frequent 1/0 in
# the current column
print('\nOxygen:')
nn = nums.copy()
c = 0  # current column
while len(nn) > 1:
    ones = countOnes(nn)
    freq1 = ones[c] / len(nn)
    bit = '1' if freq1 >= .5 else '0'
    nn = list(filter(lambda s: s[c] == bit, nn))
    print(f'  Col {c}: {len(nn)} left')
    c += 1

ox = int(nn[0], 2)
print(f'Oxygen:', nn, '=', ox)

# Filter list to numbers that have least frequent 1/0 in
# the current column
print('\nCO2:')
nn = nums.copy()
c = 0  # current column
while len(nn) > 1:
    ones = countOnes(nn)
    freq1 = ones[c] / len(nn)
    bit = '0' if freq1 >= .5 else '1'
    nn = list(filter(lambda s: s[c] == bit, nn))
    print(f'  Col {c}: {len(nn)} left')
    c += 1

co2 = int(nn[0], 2)
print(f'CO2:', nn, '=', co2)

print('\nProduct =', ox * co2)
