from typing import List

class Directory:
    def __init__(self, parent=None) -> None:
       self.sub_directories = []
       self.size_of_files = 0
       self.parent = parent

    def AddFileSize(self, size):
        self.size_of_files += size

    def AddSubDirectory(self, directory):
        self.sub_directories.append(directory)

    def Size(self) -> int:
        size = self.size_of_files
        for d in self.sub_directories:
            size += d.Size()
        return size

def buildTree(input: List[str]) -> Directory:
    root = Directory()
    currentDirectory = root
    input.pop(0)
    for line in input:
        line = line.strip().split(" ")
        if line[0] == "$":
            if line[1] == "ls":
                continue
            if line[1] == "cd":
                if line[2] == "..":
                    currentDirectory = currentDirectory.parent
                else:
                    newDir = Directory(parent=currentDirectory)
                    currentDirectory.AddSubDirectory(newDir)
                    currentDirectory = newDir
        elif line[0] == "dir":
            continue
        else:
            currentDirectory.AddFileSize(int(line[0]))
    return root

def walk(d: Directory, walkFunc):
    walkFunc(d)
    for subDir in d.sub_directories:
        walk(subDir, walkFunc)

totalSize = 0
def partOne():
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/7/input.txt', 'r')
    lines = file.readlines()
    rootDir = buildTree(lines)

    def walkFunc (d):
        global totalSize
        size = d.Size()
        if size < 100000:
            totalSize += size
    
    walk(rootDir, walkFunc)

    print(totalSize)
    
deleteSize = 0
def partTwo():
    global deleteSize
    file = open('/Users/tylerhelmuth/Projects/advent-of-code-2022/7/input.txt', 'r')
    lines = file.readlines()
    rootDir = buildTree(lines)

    size = rootDir.Size()
    unusedSpace = 70000000 - size
    amountNeeded = 30000000 - unusedSpace

    deleteSize = size
    def walkFunc (d):
        global deleteSize
        size = d.Size()
        if size > amountNeeded and size < deleteSize:
            deleteSize = size
    
    walk(rootDir, walkFunc)

    print(deleteSize)


partOne()
partTwo()