# Advent of Code 2020, Day 4
#
# Simple validation of semi-structured text "passports", first looking
# for missing fields, then checking field contents against some 
# simple rules.
#
# AK, 31/12/2021

# Read data into list of dictinaries
d = []
p = {}  # the current passport
f = 'sample.txt'
f = 'input.txt'
for l in open(f).readlines():
    l = l.strip()
    if len(l) == 0:
        if len(p) > 0:
            d.append(p)
        p = {}
    else:
        for f in l.split():
            k, v = f.split(':')
            p[k] = v
if len(p) > 0:
    d.append(p)

# List of required fields (cid optional)
req = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'] # 'cid',

# Part 1: Count valid passports, i.e., those that have all the required fields
print(f'Checking {len(d)} passports')

# For part 1, a passport is valid if it has the required fields
def check1(p):
    valid = True
    for f in req:
        if not f in p:
            return False
    return True

# Part2: more complex validation checks on field contents
def check2(p):

    # Must have the required fields as in Part 1
    if not check1(p):
        return False

    # byr (Birth Year) - four digits; at least 1920 and at most 2002.
    byr = atoi(p['byr'])
    if byr < 1920 or byr > 2002:
        return False

    # iyr (Issue Year) - four digits; at least 2010 and at most 2020.
    iyr = atoi(p['iyr'])
    if iyr < 2010 or iyr > 2020:
        return False

    # eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
    eyr = atoi(p['eyr'])
    if eyr < 2020 or eyr > 2030:
        return False

    # hgt (Height) - a number followed by either cm or in:
    # If cm, the number must be at least 150 and at most 193.
    # If in, the number must be at least 59 and at most 76.
    hgt = p['hgt']
    if not (hgt.endswith('cm') or hgt.endswith('in')):
        return False
    units = hgt[-2:]
    n = atoi(hgt[:-2])
    if units == 'cm' and (n < 150 or n > 193):
        return False
    if units == 'in' and (n < 59 or n > 76):
        return False

    # hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
    hcl = p['hcl']
    if hcl[0] != '#' or len(hcl) != 7:
        return False
    for c in hcl[1:]:
        if not ((c >= '0' and c <= '9') or (c >= 'a' and c <= 'f')):
            return False

    # ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
    ecl = p['ecl']
    if not ecl in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
        return False

    # pid (Passport ID) - a nine-digit number, including leading zeroes.
    pid = p['pid']
    if len(pid) != 9:
        return False
    for c in pid:
        if not (c >= '0' and c <= '9'):
            return False

    return True

# Text to int, -1 if invalid
def atoi(s):
    try:
        i = int(s)
    except:
        i = -1
    return i

# Part 1: count valid passports
valids = 0
for p in d:
    if check1(p):
        valids += 1

print(f'Part 1: {valids} valid passports')

# Part 2: count valid passports
valids = 0
for p in d:
    if check2(p):
        valids += 1

print(f'Part 2: {valids} valid passports')
