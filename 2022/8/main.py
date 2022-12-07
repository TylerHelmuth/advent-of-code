from typing import List

def partOne():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/8/input.txt', 'r')
    lines = file.readlines()

    forest = []
    for line in lines:
        line = line.strip()
        forest.append([n for n in line])


    numberVisible = 0
    for y in range(1, len(forest) - 1):
        for x in range(1, len(forest[y])-1):
            currentHeight = forest[y][x]
            if checkVisibility(forest, y, x, currentHeight):
                numberVisible += 1

    print(numberVisible + ((len(forest) * 4) - 4))
           

def checkVisibility(forest: List[int], row: int, col: int, currentHeight: int) -> bool:
    # check left
    visibleLeft = True
    for left in range(col-1, -1, -1):
        if forest[row][left] >= currentHeight:
            visibleLeft = False
            break
    if visibleLeft:
        return True

    # check right
    visibleRight = True
    for right in range(col+1, len(forest[row])):
        if forest[row][right] >= currentHeight:
            visibleRight = False
            break
    if visibleRight:
        return True

    # check up
    visibleUp = True
    for up in range(row-1, -1, -1):
        if forest[up][col] >= currentHeight:
            visibleUp = False
            break
    if visibleUp:
        return True

    # check down
    visibleDown = True
    for down in range(row+1, len(forest)):
        if forest[down][col] >= currentHeight:
            visibleDown = False
            break
    if visibleDown:
        return True
     
    return False

def partTwo():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/8/input.txt', 'r')
    lines = file.readlines()

    forest = []
    for line in lines:
        line = line.strip()
        forest.append([n for n in line])


    maxScore = 0
    for y in range(0, len(forest)):
        for x in range(0, len(forest[y])):
            currentHeight = forest[y][x]
            score = calculateScenicScore(forest, y, x, currentHeight)
            if score > maxScore:
                maxScore = score

    print(maxScore)

def calculateScenicScore(forest: List[int], row: int, col: int, currentHeight: int) -> int:
    # check left
    leftCount = 0
    for left in range(col-1, -1, -1):
        if forest[row][left] >= currentHeight:
            leftCount += 1
            break
        leftCount += 1


    # check right
    rightCount = 0
    for right in range(col+1, len(forest[row])):
        if forest[row][right] >= currentHeight:
            rightCount += 1
            break
        rightCount += 1

    # check up
    upCount = 0
    for up in range(row-1, -1, -1):
        if forest[up][col] >= currentHeight:
            upCount += 1
            break
        upCount += 1

    # check down
    downCount = 0
    for down in range(row+1, len(forest)):
        if forest[down][col] >= currentHeight:
            downCount += 1
            break
        downCount += 1
     
    return leftCount * rightCount * upCount * downCount


partOne()
partTwo()