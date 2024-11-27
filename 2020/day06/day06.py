# Advent of Code 2020, Day 6
#
# Read a file containing a sequence of letters on each row, each letter being a
# customer's "yes" response to a survey question. Customers are in groups,
# separated by newlines.  Part 1: For each group, count the number of questions
# to which anyone answered "yes". What is the sum of those counts?  Part 2: For
# each group, count the number of questions to which everyone answered "yes".
# What is the sum of those counts? 
#
# AK, 01/01/2022

# Read file, split into sets of question anyone responded to,
# within each group (groups separated by newlines)
f = 'sample.txt'
f = 'input.txt'
responses = []
group = []
for l in [l.strip() for l in open(f)]:

    # Blank line starts a new group
    if len(l) == 0:
        if len(group) > 0:
            responses.append(group)
        group = []

    # Otherwise, answers for one person within the current group
    else:
        p = []
        for c in l:
            p.append(c)
        group.append(p)

if len(group) > 0:
    responses.append(group)


total1 = total2 = 0
for g in responses:

    # Count up number of answers for each question in this group
    ans = {}  # count of answers
    for p in g:  # Each person
        for q in p: # That person's answers
            ans[q] = ans.get(q,0) + 1

    # Part 1: sum of For each group, count the number of questions to which
    # anyone answered "yes". What is the sum of those counts? (11 in sample,
    # 6686 in answer)
    total1 += len(ans)

    # Part 2: For each group, count the number of questions to which everyone
    # answered "yes". What is the sum of those counts? Just count up the
    # questions for which the number of answers is the size of the group
    allAns = 0
    for q in ans.keys():
        if ans[q] == len(g):
            allAns += 1
    total2 += allAns

# Show results
print('Part 1:', total1)
print('Part 2:', total2)
