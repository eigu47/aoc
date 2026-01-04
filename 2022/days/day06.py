def day_06_1(input: str) -> int:
    ans = 0

    # hashmap current seen values
    seen = {}
    for i, char in enumerate(input):
        if char not in seen:
            seen[char] = i
            if len(seen) >= 4:
                print("found", seen, char, i)
                return i + 1
        else:
            seen = {key: value for key, value in seen.items() if value > seen[char]}
            seen[char] = i

    return ans


def day_06_2(input: str) -> int:
    ans = 0

    # sliding window implementation
    seen = {}
    left = 0
    for right, char in enumerate(input):
        if char in seen and seen[char] >= left:
            if seen[char] == left:
                pass
            left = seen[char] + 1

        seen[char] = right

        if right - left + 1 >= 14:
            return right + 1

    return ans
