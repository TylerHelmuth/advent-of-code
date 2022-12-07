file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/2/input.txt', 'r')
lines = file.readlines()

outcomeDict = {
    "A X": 3,
    "A Y": 4,
    "A Z": 8,
    "B X": 1,
    "B Y": 5,
    "B Z": 9,
    "C X": 2,
    "C Y": 6,
    "C Z": 7,
}

totalScore = 0
for line in lines:
    totalScore += outcomeDict[line.strip()]

print(totalScore)
    