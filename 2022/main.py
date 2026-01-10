import sys

from days import day01, day02, day03, day04, day05, day06, day07, day08
from utils import get_day_part, get_input, get_sample, run_day

YEAR = 2022


def main():
    days = [
        day08.day_08_2,
        day08.day_08_1,
        day07.day_07_2,
        day07.day_07_1,
        day06.day_06_2,
        day06.day_06_1,
        day05.day_05_2,
        day05.day_05_1,
        day04.day_04_2,
        day04.day_04_1,
        day03.day_03_2,
        day03.day_03_1,
        day02.day_02_2,
        day02.day_02_1,
        day01.day_01_2,
        day01.day_01_1,
    ]

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

    target = sys.argv[3] if len(sys.argv) > 3 else None
    target_map = {"sample": get_sample(YEAR, day), "input": get_input(YEAR, day)}

    if target in target_map:
        run_day(selected_day, target_map[target], target.capitalize())
    else:
        for name, input in target_map.items():
            run_day(selected_day, input, name.capitalize())
            print()


if __name__ == "__main__":
    main()
