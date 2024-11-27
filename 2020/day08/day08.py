# Advent of Code 2020, Day 8
#
# Simulate execute of some simple bytecode, and detect when an instruction is
# executed the second time. For Part 2, find which nop or jmp needs to be
# changed to the opposite instruction, in order to fix the infinite loop.
#
# AK, 01/01/2022

# Read the program into a list of [instr, number] pairs
f = 'sample.txt'
f = 'input.txt'
program = []
for l in open(f):
    i, n = l.split()
    program.append([i, int(n)])

# Run the program returning final value of accumulator,
# and True if loops indefinitely (i.e., executes any
# instruction twice)
def run(pgm):

    # Set up
    ip = 0          # current instruction
    acc = 0         # the accumulator
    done = {}       # keep track of locations already executed
    loops = False   # set when attempted second execution of the same instruction is detected

    # Execute each instruction
    while True:

        # Terminated normally when attempt to execute beyond end of program
        if ip >= len(pgm):
            break

        # Stop if insruction has already been done
        if ip in done:
            loops = True
            break

        # Mark this instruction as done
        done[ip] = 1

        # Execute the instruction
        inst, n = pgm[ip]
        if inst == 'acc':
            acc += n
            ip += 1
        elif inst == 'jmp':
            ip += n
        else:
            ip += 1   # nop
    
    # Return final value in accumulator, as well as flag indicating whether the
    # program loops indefinitely
    return acc, loops


# Part 1: Run program until just before an instruction is executed the second
# time, show value in accumulator
acc, loops = run(program)
print('Part 1: acc =', acc)

# Part 2: change exactly one instruction, either nop -> jmp or vice versa,
# so that the program terminates normally (i.e., without looping forever)
print('Part 2:')
for i in range(len(program)):

    if program[i][0] in ['nop', 'jmp']:
        old = program[i][0]
        program[i][0] = 'jmp' if old == 'nop' else 'nop'
        acc, loops = run(program)
        if not loops:
            print('Found it!', acc)
        program[i][0] = old

