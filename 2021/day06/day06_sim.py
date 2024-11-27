# First attempt at Day 6 lanternfish simulation, numerically correct and fine
# for 18 days but too slow for 256 days.

# Read input data, convert to list of integers
data = open('input.txt').readlines()
data = open('test_input.txt').readlines()
data = [int(x) for x in data[0].split(',')]
print('Initial:', data)

# Start simulating up to 80 days.
# Each day, a 0 becomes a 6 and adds a new 8 to the end of the list, while each
# other number decreases by 1 if it was present at the start of the day.
t = 1
while t <= 80:
    newFish = []
    for i in range(len(data)):
        if data[i] == 0:
            data[i] = 6
            newFish.append(8)
        else:
            data[i] -= 1
    data += newFish

    if t <= 18:
        print(f'After {t} days: {data} ({len(data)})')
    elif t == 80:
        print(f'After {t} days: {len(data)}')
    t += 1
