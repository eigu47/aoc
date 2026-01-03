def day_01_1(input: str) -> int:
    ans = 0

    # for food in input.split("\n\n"):
    #     elf = 0
    #     for calories in food.splitlines():
    #         elf += int(calories)
    #     ans = max(ans, elf)
    ans = max(
        sum(int(cal) for cal in food.splitlines()) for food in input.split("\n\n")
    )

    return ans


def day_01_2(input: str) -> int:
    ans = 0

    # elves = []
    # for food in input.split("\n\n"):
    #     elf = 0
    #     for calories in food.splitlines():
    #         elf += int(calories)
    #     elves.append(elf)
    # elves.sort(reverse=True)
    # ans = sum(elves[:3])

    elves = [sum(int(cal) for cal in food.splitlines()) for food in input.split("\n\n")]
    ans = sum(sorted(elves, reverse=True)[:3])

    return ans
