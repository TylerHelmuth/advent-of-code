file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/2/input.txt', 'r')
lines = file.readlines()

pointDict = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

outcomeDict = {
    "A X": 3,
    "A Y": 6,
    "A Z": 0,
    "B X": 0,
    "B Y": 3,
    "B Z": 6,
    "C X": 6,
    "C Y": 0,
    "C Z": 3,
}

totalScore = 0
for line in lines:
    line = line.strip()
    letters = line.split(" ")
    totalScore += pointDict[letters[1]]
    totalScore += outcomeDict[line]

print(totalScore)
    