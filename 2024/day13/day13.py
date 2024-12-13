# Advent of Code 2024, Day 13
#
# Optimize a set of "machines", by finding the number of presses for buttons A
# and B that move a "claw" to the right location to pick up a prize. For each
# machine, the button moves the claw a given x and y distance.  For Part 1, did
# this easily using integer optimization with Pulp. For Part 2, add a huge
# number to each prize location. This caused the Pulp optimizer to fail,
# presumably numeric overflow. Did Simplex in Go, and tried to find approximate
# integer answer around the optimal floating point result, but did not work.
# Finally, tried command line GLPK, and found it could handle the large
# integers. So the final solution (which works for Parts 1 and 2) writes a GLPK
# model, runs the solver, and parses the result.
#
# AK, 13 Dec 2024

import os

# Parse a pair of coordinates, either "Button A: X+94, Y+34" or "Prize: X=8400, Y=5400"
def parseCoords(l):
    l = l[l.find(':')+1:]
    a, b = [s.strip() for s in l.split(',')]
    if a[1] == '=':
        a = int(a[2:])
        b = int(b[2:])
    else:
        a = int(a[1:])
        b = int(b[1:])
    return a, b


# Read all machines from input file, return as list of params
def readMachines(fname):

    machs = []
    for l in open(fname):

        # If blank line, save machine
        l = l.strip()
        if len(l) == 0:
            machs.append((Ax, Ay, Bx, By, Px, Py))
            continue
        
        # Otherwise parse coordinates
        if l.startswith('Button A'):
            Ax, Ay = parseCoords(l)
        elif l.startswith('Button B'):
            Bx, By = parseCoords(l)
        elif l.startswith('Prize'):
            Px, Py = parseCoords(l)
        else:
            print('Uknown:', l)
    return machs


# Read all machines into list of tuples
#machs = readMachines('sample.txt')
machs = readMachines('input.txt')


# Solve each machine, by writing and running a GLPK model, and add up token
# costs for successful solutions
offset = 0
offset = 10000000000000 # Uncomment this line for Part 2
ans = 0
for m in machs:

    # Write a GLPK model, using the template
    Ax, Ay, Bx, By, Px, Py = m
    template = open('template.mpl').read()
    f = open('_tmp.mpl', 'w')
    f.write(template.format(Ax = Ax, Ay = Ay, Bx = Bx, By = By, Px = Px+offset, Py = Py+offset))
    f.close()
    
    # Run the model using the command-line solver
    os.system('glpsol --math tmp.mpl > _tmp.out')

    # Parse the output
    ok = False
    for l in open('_tmp.out'):
        if l.strip() == 'INTEGER OPTIMAL SOLUTION FOUND':
            ok = True
        elif l.startswith('A.val'): # e.g., A.val = 129284425857
            A = int(l[l.find('=')+2:])
        elif l.startswith('B.val'):
            B = int(l[l.find('=')+2:])

    # If solution found, add token cost to solution
    if ok:
        ans += A*3 + B

print('Answer =', ans)
