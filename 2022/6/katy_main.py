def partSix():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/6/katy_input.txt', 'r')
    line = file.readline()

    line.strip().split()
 
    for i in range(0,len(line)-14):
        testlist = line[i:i+14]
 
        if len(set(testlist))== len(testlist):
            answer = i+14
            break
 
    print(answer)

partSix()