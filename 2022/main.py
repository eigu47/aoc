import sys

from days import day01, day02
from utils import get_day_part, get_input, get_sample, run_day

YEAR = 2022


def main():
    days = [day01.day_01_1, day01.day_01_2, day02.day_02_1, day02.day_02_2]

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
    run_day(selected_day, sample, "Sample")
    print("")
    run_day(selected_day, input, "Input")


if __name__ == "__main__":
    main()
