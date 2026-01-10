import re
from collections import defaultdict


def day_07_1(input: str) -> int:
    ans = 0

    root = {"_size": 0, "_curr": "root"}

    def updateSize(curr: dict, size: int):
        curr["_size"] += size
        if ".." in curr:
            updateSize(curr[".."], size)

    curr = root
    for line in input.splitlines():
        match = re.match(r"^\$ (\S+)(?: (.+))?", line)
        if match is not None and match.group(1) == "cd":
            dir = match.group(2)
            if dir not in curr:
                curr[dir] = {"..": curr, "_size": 0, "_curr": dir}
            curr = curr[dir]
            continue
        match = re.match(r"^(\d+) (.+)$", line)
        if match is not None:
            curr[match.group(2)] = int(match.group(1))
            updateSize(curr, int(match.group(1)))

    def printDirs(dir: dict, depth: int):
        for key, value in dir.items():
            if key == ".." or key.startswith("_"):
                continue
            if isinstance(value, dict):
                print(f"{" "*depth}/{key} (dir) {dir[key]["_size"]}")
                printDirs(value, depth + 1)
            else:
                print(f"{" "*depth}{key} (file) {value}")

    printDirs(root, 0)

    def searchDirs(dir: dict, max: int) -> int:
        ans = 0
        for key, value in dir.items():
            if key == "_size" and value <= max:
                ans += value
            elif key != ".." and isinstance(value, dict):
                ans += searchDirs(value, max)
        return ans

    ans = searchDirs(root, 100_000)

    return ans


def day_07_2(input: str) -> int:
    ans = 0

    dirs = defaultdict(int)
    curr = []
    for line in input.splitlines():
        match = re.match(r"^\$ (\S+)(?: (.+))?", line)
        if match is not None and match.group(1) == "cd":
            dir = match.group(2)
            if dir == "..":
                curr.pop()
            else:
                curr.append(dir)
            continue
        match = re.match(r"^(\d+) (.+)$", line)
        if match is not None:
            path = ""
            for part in curr:
                path = f"{path}/{part}"
                dirs[path] += int(match.group(1))

    # print(dirs)

    need = dirs["/root"] - (70_000_000 - 30_000_000)
    ans = dirs["/root"]

    for size in dirs.values():
        if size >= need:
            ans = min(ans, size)

    return ans
