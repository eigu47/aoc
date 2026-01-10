def day_08_1(input: str) -> int:
    ans = 0

    grid = [[int(cell) for cell in line] for line in input.splitlines()]
    ans = len(grid) * 2 + len(grid[0]) * 2 - 4

    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[i]) - 1):
            for dir in dirs:
                di = i + dir[0]
                dj = j + dir[1]
                while 0 <= di < len(grid) and 0 <= dj < len(grid[0]):
                    if grid[i][j] <= grid[di][dj]:
                        break
                    di += dir[0]
                    dj += dir[1]
                else:
                    ans += 1
                    break

    return ans


def day_08_2(input: str) -> int:
    ans = 0

    grid = [[int(cell) for cell in line] for line in input.splitlines()]
    scores = {}

    for i in range(1, len(grid) - 1):
        for j in range(1, len(grid[i]) - 1):
            score = 1
            t = []
            for dir in dirs:
                trees = 0
                di = i + dir[0]
                dj = j + dir[1]
                while 0 <= di < len(grid) and 0 <= dj < len(grid[0]):
                    if grid[i][j] <= grid[di][dj]:
                        trees += 1
                        break
                    trees += 1
                    di += dir[0]
                    dj += dir[1]
                score *= trees
                t.append(trees)
            ans = max(ans, score)
            scores[f"{i},{j}:{grid[i][j]}"] = t

    print(scores)
    return ans


dirs = [
    [-1, 0],  # UP
    # [-1, 1],  # UP RGT
    [0, 1],  # RGT
    # [1, 1],   # DWN RGT
    [1, 0],  # DWN
    # [1, -1],  # DWN LFT
    [0, -1],  # LFT
    # [-1, -1], # UP LFT
]
