def partOne():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/4/input.txt', 'r')
    lines = file.readlines()

    total = 0
    for line in lines:
        line = line.strip()
        assignments = line.split(",")

        elf1 = []        
        r = assignments[0].split("-")
        for i in range(int(r[0]), int(r[1])+1):
            elf1.append(i)

        elf2 = []        
        r = assignments[1].split("-")
        for i in range(int(r[0]), int(r[1])+1):
            elf2.append(i)

        if set(elf1).issubset(set(elf2)) or set(elf2).issubset(set(elf1)):
            total += 1
    
    print(total)

def partTwo():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/4/input.txt', 'r')
    lines = file.readlines()

    total = 0
    for line in lines:
        line = line.strip()
        assignments = line.split(",")

        elf1 = []        
        r = assignments[0].split("-")
        for i in range(int(r[0]), int(r[1])+1):
            elf1.append(i)

        elf2 = []        
        r = assignments[1].split("-")
        for i in range(int(r[0]), int(r[1])+1):
            elf2.append(i)

        if not set(elf1).isdisjoint(set(elf2)):
            total += 1
    
    print(total)

partOne()
partTwo()