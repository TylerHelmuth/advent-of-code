



def partOne():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/3/input.txt', 'r')
    lines = file.readlines()
    total = 0
    for line in lines:
        line = line.strip()
        lineLength = len(line)
        middle = lineLength//2
        compartmentOne = line[0:middle]
        compartmentTwo = line[middle:lineLength]

        intersection = ''.join(set(compartmentOne).intersection(compartmentTwo))

        letters = intersection.split()

        for letter in letters:
            if letter.isupper():
                total += ord(letter) - 38
            else:
                total += ord(letter) - 96

    print(total)

def partTwo():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/3/input.txt', 'r')
    lines = file.readlines()
    total = 0
    for i in range(0, len(lines), 3):
        elf1 = lines[i].strip()
        elf2 = lines[i+1].strip()
        elf3 = lines[i+2].strip()

        intersection = ''.join(set(elf1).intersection(elf2).intersection(elf3))

        if intersection.isupper():
            total += ord(intersection) - 38
        else:
            total += ord(intersection) - 96
           

    print(total)

partOne()
partTwo()