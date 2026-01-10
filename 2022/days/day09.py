def day_09_1(input: str) -> int:
    ans = 0
    head = [0, 0]
    tail = [0, 0]
    prev = [0, 0]
    visited = {",".join(map(str, head))}
    path = []

    for line in input.splitlines():
        dir, steps = line.split()
        # print(dir, steps)
        for _ in range(int(steps)):
            head = [
                head[0] + dirs[dir][0],
                head[1] + dirs[dir][1],
            ]
            for d in dirs.values():
                if tail[0] == head[0] + d[0] and tail[1] == head[1] + d[1]:
                    x = prev[0] + dirs[dir][0]
                    y = prev[1] + dirs[dir][1]
                    prev = [(x > 0) - (x < 0), (y > 0) - (y < 0)]
                    break
            else:
                tail[0], tail[1] = (
                    tail[0] + prev[0],
                    tail[1] + prev[1],
                )
                visited.add(",".join(map(str, tail)))
                # path.append(",".join(map(str, tail)))
                prev = dirs[dir]
            # print(head, tail)

    # print(path)

    ans = len(visited)
    return ans


def day_09_2(input: str) -> int:
    ans = 0

    return ans


dirs = {
    "I": [0, 0],
    "U": [-1, 0],
    "R": [0, 1],
    "D": [1, 0],
    "L": [0, -1],
    "UR": [-1, 1],
    "DR": [1, 1],
    "DL": [1, -1],
    "UL": [-1, -1],
}
