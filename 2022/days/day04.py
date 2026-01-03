import re
from itertools import permutations


def day_04_1(input: str) -> int:
    ans = 0

    # for line in input.splitlines():
    #     pair = []
    #     for pairStr in line.split(","):
    #         num = pairStr.split("-")
    #         p = set()
    #         for n in range(int(num[0]), int(num[1]) + 1):
    #             p.add(n)
    #         pair.append(p)
    #     if pair[0].issubset(pair[1]) or pair[1].issubset(pair[0]):
    #         ans += 1

    for line in input.splitlines():
        a_str, b_str = line.split(",")

        a_start, a_end = map(int, a_str.split("-"))
        a_set = set(range(a_start, a_end + 1))
        b_start, b_end = map(int, b_str.split("-"))
        b_set = set(range(b_start, b_end + 1))

        # print(a_set, b_set)
        if a_set.issubset(b_set) or b_set.issubset(a_set):
            ans += 1

    return ans


def day_04_2(input: str) -> int:
    ans = 0

    for line in input.splitlines():
        a_str, b_str = line.split(",")

        a_start, a_end = map(int, a_str.split("-"))
        a_set = set(range(a_start, a_end + 1))
        b_start, b_end = map(int, b_str.split("-"))
        b_set = set(range(b_start, b_end + 1))

        # print(a_set, b_set)
        rep = next((iter(a_set & b_set)), None)
        if rep is not None:
            ans += 1

    return ans
