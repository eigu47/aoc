def day_02_1(input: str) -> int:
    ans = 0

    outcome = {
        # ROCK
        "A": {
            # ROCK
            "X": 3 + 1,
            # PAPER
            "Y": 6 + 2,
            # SCISSOR
            "Z": 0 + 3,
        },
        # PAPER
        "B": {
            # ROCK
            "X": 0 + 1,
            # PAPER
            "Y": 3 + 2,
            # SCISSOR
            "Z": 6 + 3,
        },
        # SCISSOR
        "C": {
            # ROCK
            "X": 6 + 1,
            # PAPER
            "Y": 0 + 2,
            # SCISSOR
            "Z": 3 + 3,
        },
    }

    # for line in input.splitlines():
    #     game = line.split(" ")
    #     ans += outcome[game[0]]

    ans = sum(
        outcome[opp][you] for line in input.splitlines() for opp, you in (line.split(),)
    )

    return ans


def day_02_2(input: str) -> int:
    ans = 0

    outcome = {
        # ROCK
        "A": {
            # LOSE / SCISSOR
            "X": 0 + 3,
            # DRAW / ROCK
            "Y": 3 + 1,
            # WIN / PAPER
            "Z": 6 + 2,
        },
        # PAPER
        "B": {
            # LOSE / ROCK
            "X": 0 + 1,
            # DRAW / PAPER
            "Y": 3 + 2,
            # WIN / SCISSOR
            "Z": 6 + 3,
        },
        # SCISSOR
        "C": {
            # LOSE / PAPER
            "X": 0 + 2,
            # DRAW / SCISSOR
            "Y": 3 + 3,
            # WIN / ROCK
            "Z": 6 + 1,
        },
    }

    ans = sum(
        (outcome[opp][you])
        for line in input.splitlines()
        for opp, you in (line.split(),)
    )

    return ans
