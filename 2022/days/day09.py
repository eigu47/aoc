def day_09_1(input: str) -> int:
    ans = 0
    head = [0, 0]
    tail = [0, 0]
    visited = {",".join(map(str, head))}

    for line in input.splitlines():
        dir, steps = line.split()
        for _ in range(int(steps)):
            head = [
                head[0] + dirs[dir][0],
                head[1] + dirs[dir][1],
            ]
            dy = head[0] - tail[0]
            dx = head[1] - tail[1]
            if abs(dy) > 1 or abs(dx) > 1:
                tail = [
                    tail[0] + ((dy > 0) - (dy < 0)),
                    tail[1] + ((dx > 0) - (dx < 0)),
                ]

            # print(head, tail)
            visited.add(",".join(map(str, tail)))

    ans = len(visited)
    return ans


def day_09_2(input: str) -> int:
    ans = 0

    rope = [[0, 0]] * 10
    visited = {f"{rope[0]}"}

    for line in input.splitlines():
        dir, steps = line.split()
        for _ in range(int(steps)):
            for i in range(len(rope)):
                if i == 0:
                    rope[i] = [
                        rope[i][0] + dirs[dir][0],
                        rope[i][1] + dirs[dir][1],
                    ]
                else:
                    dy = rope[i - 1][0] - rope[i][0]
                    dx = rope[i - 1][1] - rope[i][1]
                    if abs(dy) > 1 or abs(dx) > 1:
                        rope[i] = [
                            rope[i][0] + ((dy > 0) - (dy < 0)),
                            rope[i][1] + ((dx > 0) - (dx < 0)),
                        ]
            visited.add(f"{rope[-1]}")

    print(visited)
    ans = len(visited)
    return ans


dirs = {
    "U": [-1, 0],
    "R": [0, 1],
    "D": [1, 0],
    "L": [0, -1],
}
