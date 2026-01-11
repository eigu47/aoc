def day_10_1(input: str) -> int:
    ans = 0

    signal = {}
    cycle = 0
    x = 1
    for line in input.splitlines():
        op, *add = line.split()
        for _ in range(1 if op == "noop" else 2):
            cycle += 1
            if cycle == 20 or (cycle - 20) % 40 == 0:
                signal[cycle] = x
        x += int(*add)

    ans = sum([cycle * x for cycle, x in signal.items()])
    return ans


def day_10_2(input: str) -> int:
    ans = 0

    grid = [["."] * 40 for _ in range(6)]
    cycle = 0
    x_re = 1
    for line in input.splitlines():
        op, *add = line.split()
        for _ in range(1 if op == "noop" else 2):
            y = cycle // 40
            x = cycle % 40
            if abs(x_re - x) < 2:
                grid[y][x] = "#"
            cycle += 1
        x_re += int(*add)

    for line in grid:
        print("".join(line))

    return ans


# ZRARLFZU
# ####.###...##..###..#....####.####.#..#.
# ...#.#..#.#..#.#..#.#....#.......#.#..#.
# ..#..#..#.#..#.#..#.#....###....#..#..#.
# .#...###..####.###..#....#.....#...#..#.
# #....#.#..#..#.#.#..#....#....#....#..#.
# ####.#..#.#..#.#..#.####.#....####..##..
