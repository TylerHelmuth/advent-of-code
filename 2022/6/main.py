def partOne():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/6/input.txt', 'r')
    line = file.readline()
    line.strip().split()

    answer = -1
    for i in range(4, len(line)+1):
        if len(set(line[i-4:i])) == len(line[i-4:i]):
            answer = i
            break

    print(answer)
    

def partTwo():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/6/input.txt', 'r')
    line = file.readline()
    line.strip().split()
    
    answer = -1
    for i in range(14, len(line)+1):
        if len(set(line[i-14:i])) == len(line[i-14:i]):
            answer = i
            break

    print(answer)

partOne()
partTwo()