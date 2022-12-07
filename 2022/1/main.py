file1 = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/1/input.txt', 'r')
lines = file1.readlines()
firstCalories = 0
secondCalories = 0
thirdCalories = 0
runningSum = 0
for line in lines:
    if line.strip().isnumeric():
        runningSum += int(line)
    else:
        if runningSum > firstCalories:
            thirdCalories = secondCalories
            secondCalories = firstCalories
            firstCalories = runningSum
        
        elif runningSum > secondCalories:
            thirdCalories = secondCalories
            secondCalories = runningSum
        
        elif runningSum > thirdCalories:
            thirdCalories = runningSum

        runningSum = 0

print(firstCalories + secondCalories + thirdCalories)