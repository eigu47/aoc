import re
from collections import defaultdict
from operator import concat


def day_05_1(input: str) -> int:
    ans = 0

    cranes_split, step_split = input.split("\n\n")

    lines = cranes_split.splitlines()[:-1]
    cranes = {i: [] for i, _ in enumerate(lines[0][1::4])}
    for _, line in enumerate(lines):
        for j, char in enumerate(line[1::4]):
            if char != " ":
                cranes[j].insert(0, char)

    for line in step_split.splitlines():
        move, from_, to = map(int, re.findall(r"\d+", line))
        for _ in range(move):
            # print(f"move {move}, from {cranes[from_ -1]}, to {cranes[to -1]}")
            cranes[to - 1].append(cranes[from_ - 1].pop())

    # print(cranes)
    ans = "".join(crane[-1] for crane in cranes.values())

    return ans


def day_05_2(input: str) -> int:
    ans = 0

    cranes_split, step_split = input.split("\n\n")

    lines = cranes_split.splitlines()[:-1]
    cranes = {i: [] for i, _ in enumerate(lines[0][1::4])}
    for _, line in enumerate(lines):
        for j, char in enumerate(line[1::4]):
            if char != " ":
                cranes[j].insert(0, char)

    for line in step_split.splitlines():
        move, from_, to = map(int, re.findall(r"\d+", line))
        # print(f"move {move}, from {cranes[from_ -1]}, to {cranes[to -1]}")
        cranes[to - 1].extend(cranes[from_ - 1][-move:])
        cranes[from_ - 1] = cranes[from_ - 1][:-move]

    # print(cranes)
    ans = "".join(crane[-1] for crane in cranes.values())

    return ans
