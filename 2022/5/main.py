def partOne():
    stacks = {
        "1": ["S", "C", "V", "N"],
        "2": ["Z", "M", "J", "H", "N", "S"],
        "3": ["M", "C", "T", "G", "J", "N", "D"],
        "4": ["T", "D", "F", "J", "W", "R", "M"],
        "5": ["P", "F", "H"],
        "6": ["C", "T", "Z", "H", "J"],
        "7": ["D", "P", "R", "Q", "F", "S", "L", "Z"],
        "8": ["C", "S", "L", "H", "D", "F", "P", "W"],
        "9": ["D", "S", "M", "P", "F", "N", "G", "Z"],
    }

    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/5/input.txt', 'r')
    lines = file.readlines()

    for line in lines:
        if line[0:4] != "move":
            continue
        
        numberStr = line.strip().replace("move", "").replace("from", "").replace("to", "")
        numbers = numberStr.strip().split()
        
        for _ in range (0, int(numbers[0])):
            moving = stacks[numbers[1]].pop()
            stacks[numbers[2]].append(moving)

    result = ""
    for i in range (1, 10):
        result += stacks[str(i)][-1]
    print(result)

["Z", "G", "N"]
    

def partTwo():
    stacks = {
        "1": ["S", "C", "V", "N"],
        "2": ["Z", "M", "J", "H", "N", "S"],
        "3": ["M", "C", "T", "G", "J", "N", "D"],
        "4": ["T", "D", "F", "J", "W", "R", "M"],
        "5": ["P", "F", "H"],
        "6": ["C", "T", "Z", "H", "J"],
        "7": ["D", "P", "R", "Q", "F", "S", "L", "Z"],
        "8": ["C", "S", "L", "H", "D", "F", "P", "W"],
        "9": ["D", "S", "M", "P", "F", "N", "G", "Z"],
    }

    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/5/input.txt', 'r')
    lines = file.readlines()

    for line in lines:
        if line[0:4] != "move":
            continue
        
        numberStr = line.strip().replace("move", "").replace("from", "").replace("to", "")
        numbers = numberStr.strip().split()
        
        itemsBeingMoved = []
        for _ in range (0, int(numbers[0])):
            item = stacks[numbers[1]].pop()
            itemsBeingMoved.append(item)

        itemsBeingMoved.reverse()

        for i in range (len(itemsBeingMoved)):
            stacks[numbers[2]].append(itemsBeingMoved[i])

    result = ""
    for i in range (1, 10):
        result += stacks[str(i)][-1]
    print(result)

partOne()
partTwo()