# Advent of Code 2021, Day 17
#
# Simple projectile simulation, with search for
# velocities that achieve highest position, and total
# number of possible velocity values that reach target
# area; just used a simple grid search.
#
# AK, 17/12/2021

# Target area (x range and y range)
T = ((20,30), (-10, -5))     # Sample data
T = ((192, 251), (-89, -59)) # Real data

# Determine whether coordinates are within target area
def withinTarget(x, y):
    return x >= T[0][0] and x <= T[0][1] and y >= T[1][0] and y <= T[1][1]

# Returns the highest position achieved
def simulate(xv, yv):

    # Start at 0,0
    #print(f'Simulating {xv}, {yv}')
    x = y = 0
    highest = -99999        # Highest y reached, for Part 1
    hitTarget = False       # Whether or not hit target area

    # Simulate steps
    iter = 0
    while iter < 300:       # Stop after lots of iterations

        # Position changes according to velocity
        iter += 1
        x += xv
        y += yv

        # Record highest position
        if iter == 0 or y > highest:
            highest = y

        # Adjust velocity
        if xv > 0:
            xv -= 1
        elif xv < 0:
            xv += 1
        yv -= 1

        #print(f'  Position = {x},{y}, velocity = {xv},{yv}')

        # Stop if reached target area
        if withinTarget(x, y):
            #print('  Within target!')
            hitTarget = True
            break

    # If hit target, return highest position reached, otherwise
    # return None
    #print('  Highest y =', highest)
    return highest if hitTarget else None

# Simple pseudo-optimization, just a grid search, finds
# the highest position reached (for Part 1) and also counts
# up the number of solutions that hit target (for Part 2)
def optimize():
    best = count = 0
    for xv in range(0, 300):    # Set bounds range of problem
        for yv in range(-100, 100):
            ypos = simulate(xv, yv)
            if ypos != None:    # Only count if hit target
                if ypos > best:
                    best = ypos
                count += 1

    print('Best found =', best)
    print('Solutions found =', count)

# Test examples
#for xv, yv in [(7, 2), (6, 3), (9, 0), (17, -4), (6, 9)]:
#    print(xv, yv, simulate(xv, yv))

# Run the solution
optimize()

