# Advent of Code 2020, Day 5
#
# Apply a sequence of binary forward/back operators to find row
# and seat in an airplane, and compute a code for these. For Part 1,
# find the maximum seat ID, for Part 2 the missing one.
#
# AK, 01/01/2022

# Process one instruction sequence
def process(t):

    # Get row (chars 1-7): F means first half of remaining
    # array, B means second half
    row = 0
    size = 128
    for c in t[:7]:
        size /= 2
        if c == 'B':
            row += size

    # Get seat (chars 8-10)::w

    seat = 0
    size = 8
    for c in t[7:]:
        size /= 2
        if c == 'R':
            seat += size

    # Seat ID: row * 8 + column
    sid = row * 8 + seat

    # Return row, seat, and seat ID
    return (row, seat, sid)

# Process all input rows
f = 'sample.txt'
f = 'input.txt'
seats = [process(l.strip()) for l in open(f)]

# Part 1: maximum Seat ID
sids = [s[2] for s in seats]
print('Part 1:', max(sids))
        
# Part 2: find seat that is missing a boarding pass
alloc = {} # seatID => 1
for sid in sids:
    alloc[sid] = 1
for row in range(1, 128):  # leave out first & last rows
    for seat in range(8):
        sid = row * 8 + seat
        if sid-1 in alloc and sid+1 in alloc and not sid in alloc:
            print('Missing seat:', sid)

