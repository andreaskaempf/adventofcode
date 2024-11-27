# Advent of Code, Day 16
#
# Parse hierarchy of binary codes into numeric literals and operations.
# For part 2, recursively evaluate the resulting expression using a
# LISP-like approach.
#
# AK, 16/12/2021

# For converting hex digit to binary
h2b = {'0': '0000', '1': '0001', '2': '0010', '3': '0011',
    '4': '0100', '5': '0101', '6': '0110', '7': '0111', '8': '1000',
    '9': '1001', 'A': '1010', 'B': '1011', 'C': '1100', 'D': '1101',
    'E': '1110', 'F': '1111'}

# Convert hex string to binary digits
def hex2binary(h):
    return ''.join([h2b[c] for c in h])

# Convert binary string to integer
def b2i(b):
    return 0 if len(b) == 0 else int(b, 2)

# Determine if string is just zeroes
def padding(s):
    nonzero = [0 if c == '0' else 1 for c in s]
    return sum(nonzero) == 0

# Determine if a number
def isnumber(n):
    return type(n) == type(1) or type(n) == type(1.0)

# Determine if a list
def islist(l):
    return type(l) == type([])

# Global variable for tracking total versions in Part 1
totalVersions = 0

# Op codes for Part 2
opcodes = { 0: '+', 1: '*', 2: 'min', 3: 'max', 5: '>', 6: '<', 7: '=' }

# Determine if an operation code
def isop(o):
    return o in opcodes.values()

# Top-level function for parsing packets, takes a string
# and current position, and returns data and new position
def parsePackets(s, i, doAll = True):

    # Go through string from current position
    data = []
    didOne = False
    while i < len(s):

        # If not doing all and already did one, stop here
        if didOne and not doAll:
            break

        # If the string is short and all zeros, safe to stop here
        if padding(s[i:]):
            break

        # Parse header
        header, i = parseHeader(s, i)
        if not header:
            print('Parsing header failed:', s[i:])
            break  #exit()

        # Packet type 4: literal packet
        htype = header['type']
        if htype == 4:
            num, i = parseLiteral(s, i)
            data.append(num)
            didOne = True
            continue

        # Otherwise parse an operator packet:
        # - headers
        # - length type ID (1 bit): 
        #   0 -> next 15 bits lenght of sub-packets
        #   1 -> next 11 bits are *number* of sub-packets
        # - the 15 or 11-bit length
        # - the sub-packets

        # For part 2, add the operation for the header type
        data.append(opcodes[htype])

        # Depending on the length ID bit, get length of sub-packets, or the
        # number of sub-packets, then get the packets.
        lid = s[i] # Starting with length ID bit
        i += 1
        if lid == '0': # 0 -> next 15 bits *length* of sub-packets
            lsubs = b2i(s[i:i+15])
            i += 15
            stopAt = i + lsubs
            while i < stopAt:
                p, i = parsePackets(s, i, False)  # just one at a time
                data.append(p)
        else:       # 1 -> next 11 bits are *number* of sub-packets
            nsubs = b2i(s[i:i+11])
            i += 11
            for _ in range(nsubs):
                p, i = parsePackets(s, i, False)  # just one at a time
                data.append(p)

        # So we can optionally stop after one top-level packet
        didOne = True

    # Unravel nested lists of 1 (e.g., [2])
    if islist(data) and len(data) == 1:
        data = data[0]

    # Return data found and current position
    return data, i


# Parse header: version and type (each 3 bits). Pass it the string
# and current index, returns dictionary of values and new index
def parseHeader(s, i):

    global totalVersions

    # Stop if not enough chars left to parse version and type
    charsLeft = len(s) - i
    if charsLeft < 6:
        print(charsLeft, 'not enough chars for header')
        return None, i

    # First 3 bits is the version number
    version = b2i(s[i:i+3])
    i += 3

    # For part 1, add to total of versions
    totalVersions += version

    # Next 3 bits are the packet type
    ptype = b2i(s[i:i+3])
    i += 3

    # Return version and type as dictionary
    return {'version': version, 'type': ptype}, i


# Parse a numeric literal, without header
def parseLiteral(s, i):

    # Each chunk of 5 bits starting with 1 is 1 bit indicator,
    # followed by 4 binary digits
    i0 = i
    num = ''
    while s[i] == '1':
        num += s[i+1:i+5]
        i += 5

    # Last chunk (starts with 0)
    num += s[i+1:i+5]
    i += 5

    # Return number and new position
    n = b2i(num)
    return n, i


# Run tests for Part 1

print('Sample 1: literal 2021')
sample1 = '110100101111111000101000'
print(parsePackets(sample1, 0))

print('\nSample 2: operator packet containing literals 10 and 20')
sample2 = '00111000000000000110111101000101001010010001001000000000'
print(parsePackets(sample2, 0))

print('\nSample 3: operator packet with number 3 sub-packets (literals, 1, 2, 3)')
sample3 = '11101110000000001101010000001100100000100011000001100000'
print(parsePackets(sample3, 0))

print('\nSample 4: literal inside 3 nested operator packets')
sample4 = hex2binary('8A004A801A8002F478')
totalVersions = 0
print(parsePackets(sample4, 0))
print(f'Total of versions = {totalVersions} (should be 16)')

print('\nSample 5: operator packet containing two operator packets, each containing two literals')
sample5 = hex2binary('620080001611562C8802118E34')
totalVersions = 0
print(parsePackets(sample5, 0))
print(f'Total of versions = {totalVersions} (should be 12)')

print('\nSample 6: operator packet containing two operator packets, each containing two literals')
sample6 = hex2binary('C0015000016115A2E0802F182340')
totalVersions = 0
print(parsePackets(sample6, 0))
print(f'Total of versions = {totalVersions} (should be 23)')

print('\nSample 7: 5 literals inside 3 nested operator packets')
sample7 = hex2binary('A0016C880162017C3686B18A3D4780')
totalVersions = 0
print(parsePackets(sample7, 0))
print(f'Total of versions = {totalVersions} (should be 31)')

# Try part 1 on input data
if False:
    print('\nPart 1, input data')
    hinput = open('input.txt').read().strip()
    binput = hex2binary(hinput)
    totalVersions = 0
    print(parsePackets(binput, 0))
    print(f'Total of versions = {totalVersions} (should be 934)')

# For Part 2, recursively evaluate an expression with LISP-like structure
def evaluate(ast):

    # A number evaluates to itself
    if isnumber(ast):
        return ast

    # Otherwise apply opcodes, either to following list or to remaining elements
    # opcodes = { 0: '+', 1: '*', 2: 'min', 3: 'max', 5: '>', 6: '<', 7: '=' }
    if islist(ast) and len(ast) > 0 and isop(ast[0]):
        args = ast[1:]
        if ast[0] == '+':
            return sum([evaluate(x) for x in args])
        elif ast[0] == '*':
            ans = 1.0
            for x in [evaluate(x) for x in args]:
                ans *= x
            return ans
        elif ast[0] == 'min':
            return min([evaluate(x) for x in args])
        elif ast[0] == 'max':
            return max([evaluate(x) for x in args])
        elif ast[0] == '>':
            return 1 if evaluate(args[0]) > evaluate(args[1]) else 0
        elif ast[0] == '<':
            return 1 if evaluate(args[0]) < evaluate(args[1]) else 0
        elif ast[0] == '=':
            return 1 if evaluate(args[0]) == evaluate(args[1]) else 0

    # If we get to here, could not handle this
    print('Cannot handle:', ast)
    return 0

# Test cases for Part 2
tests = [('C200B40A82', 'finds the sum of 1 and 2, resulting in the value 3'),
    ('04005AC33890', 'finds the product of 6 and 9, resulting in the value 54'),
    ('880086C3E88112', 'finds the minimum of 7, 8, and 9, resulting in the value 7.'),
    ('CE00C43D881120', ' finds the maximum of 7, 8, and 9, resulting in the value 9.'),
    ('D8005AC2A8F0', ' produces 1, because 5 is less than 15.'),
    ('F600BC2D8F', ' produces 0, because 5 is not greater than 15.'),
    ('9C005AC2F8F0', ' produces 0, because 5 is not equal to 15.'),
    ('9C0141080250320F1802104A08', ' produces 1, because 1 + 3 = 2 * 2.')]

for t in tests:
    print('\n', t[0], t[1])
    ast, _ = parsePackets(hex2binary(t[0]), 0)
    print(ast)
    print("Answer =", evaluate(ast))

# Try part 2 on input data
if True:
    print('\nPart 2, input data')
    hinput = open('input.txt').read().strip()
    binput = hex2binary(hinput)
    ast, _ = parsePackets(binput, 0)
    print("\nAST =", ast)
    print("\nAnswer =", evaluate(ast))
