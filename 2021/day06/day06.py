# Day 6 of Advent of Code 2021
#
# Simulate evolution of population of "lanterfish", starting at 5 but
# reproducing every 7 days. Doing this with simulation of a list as described
# in the problem is way too slow, so converted it to using a dictionary to
# keep track of the number of fish by age, much faster.
#
# AK, 6 Dec 2021

# Read input data, convert to list of integers
#data = open('test_input.txt').readlines()
data = open('input.txt').readlines()
data = [int(x) for x in data[0].split(',')]
print('Initial:', data)

# Turn it into a dictionary, number of fish of each age
fish = {}
for age in range(9):        # Initialize every possible age to ease work
    fish[age] = 0
for f in data:              # Count up fish at each age
    fish[f] += 1
print('Initial dict:', fish, '=', sum(fish.values()))

# Start simulating up to 80 days (256 days in part 2). Each day, fish of age 0
# become 6 and each adds a new fish of age 8 to the population, while each
# other fish decreases its age by 1.
# For test data, should be 26 fish after day 18, 5934 after 80 days, and
# 26984457539 after 256 days
t = 1
while t <= 256:

    # Each fish with a counter of 0 will create a new fish with counter 8 (but
    # don't do it yet, as we don't want to decrement the age of the new fish)
    newFish8 = fish[0]

    # All fish with counter 0 will get reset to 6 (but don't do it yet, so
    # we don't decrement the age of those fish in the next step)
    newFish6 = fish[0]
    fish[0] = 0

    # All fish of ages 1-8 get decremented, by moving them down the chain:
    #   fish0 += fish1
    #   fish1 = 0
    # then
    #   fish1 += fish2
    #   fish2 = 0
    # etc.
    for age in range(1, 9): 
        fish[age-1] += fish[age]
        fish[age] = 0

    # Finally, add the new fish of ages 6 and 8
    fish[6] += newFish6
    fish[8] += newFish8

    # Show current count
    if t <= 18 or t == 80 or t == 256:
        nfish = sum(fish.values())
        print(f'After {t} days: {nfish}')

    # Next time period
    t += 1

