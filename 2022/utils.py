import os
import re
import time
from pathlib import Path
from typing import Callable

import requests
from dotenv import load_dotenv

INPUT_DIR = Path("input")
INPUT_DIR.mkdir(parents=True, exist_ok=True)


def get_input(year: int, day: int) -> str:
    file_path = INPUT_DIR / f"{year}_{day:02d}.txt"

    if file_path.exists():
        return file_path.read_text("utf-8")

    print(f"Fetching data: year {year}, day {day}")

    load_dotenv()
    session_id = os.getenv("SESSION_ID")
    if session_id is None:
        raise RuntimeError("SESSION_ID was not found")

    response = requests.get(
        url=f"https://adventofcode.com/{year}/day/{day}/input",
        cookies={"session": session_id},
    )
    response.raise_for_status()

    file_path.write_text(response.text, encoding="utf-8")
    return response.text


def get_sample(year: int, day: int) -> str:
    file_path = INPUT_DIR / f"{year}_{day:02d}_sample.txt"

    file_path.touch(exist_ok=True)
    return file_path.read_text(encoding="utf-8")


def get_day_part(fn: Callable) -> tuple[int, int] | None:
    match = re.findall(r"\d+", fn.__name__)
    if len(match) != 2:
        return None
    return int(match[0]), int(match[1])


def run_day(fn: Callable, input: str, type: str):
    if len(input) < 1:
        return
    start = time.perf_counter()
    ans = fn(input)
    print(f"{type} time {(time.perf_counter() - start)*1000:.3f}ms:\n{ans}")
