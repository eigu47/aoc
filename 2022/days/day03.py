def day_03_1(input: str) -> int:
    ans = 0

    for line in input.splitlines():
        i1 = line[: len(line) // 2]
        i2 = line[len(line) // 2 :]
        rep = next((l for l in i1 if l in i2), None)
        # print(rep, ord(rep))
        dec = ord(rep)
        if dec >= ord("a"):
            dec -= ord("a") - 1
        else:
            dec -= ord("A") - 27
        ans += dec

    # ans = sum(
    #     (ord(rep) - ord("a") + 1) if rep.islower() else (ord(rep) - ord("A") + 27)
    #     for line in input.splitlines()
    #     for rep in [
    #         next(iter(set(line[: len(line) // 2]) & set(line[len(line) // 2 :])))
    #     ]
    # )

    return ans


def day_03_2(input: str) -> int:
    ans = 0

    input = input.splitlines()
    for i in range(0, len(input), 3):
        rep = next((iter(set(input[i]) & set(input[i + 1]) & set(input[i + 2]))), None)
        dec = ord(rep)
        if dec >= ord("a"):
            dec -= ord("a") - 1
        else:
            dec -= ord("A") - 27
        ans += dec

    return ans
