import sys
import time

from days import day01
from utils import get_day_part, get_input, get_sample

YEAR = 2022


def main():
    days = [day01.day_01_1]

    if len(sys.argv) > 2:
        day = int(sys.argv[1])
        part = int(sys.argv[2])
        selected_day = next(
            (fn for fn in days if get_day_part(fn) == (day, part)), None
        )
        if selected_day is None:
            raise RuntimeError("Selected day not found")

    else:
        selected_day = days[0]
        day, part = get_day_part(selected_day)

    input = get_input(YEAR, day)
    sample = get_sample(YEAR, day)

    print(f"Day {day}, part {part}:")
    if len(sample) > 1:
        start = time.perf_counter()
        ans = selected_day(sample)
        print(f"Sample time {(time.perf_counter() - start)*1000:.3f}ms:\n{ans}\n")

    start = time.perf_counter()
    ans = selected_day(input)
    print(f"Input time {(time.perf_counter() - start)*1000:.3f}ms:\n{ans}")


if __name__ == "__main__":
    main()
