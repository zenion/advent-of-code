from dataclasses import dataclass
from pathlib import Path
from typing import List


@dataclass
class Columns:
    left: List[int]
    right: List[int]


def parse_columns(contents: str) -> Columns:
    lines = [line for line in contents.split("\n") if line.strip()]
    left, right = [], []

    for line in lines:
        left_num, right_num = map(int, line.split())
        left.append(left_num)
        right.append(right_num)

    return Columns(sorted(left), sorted(right))


def solve_part1(contents: str) -> int:
    cols = parse_columns(contents)
    return sum(
        abs(left_num - right_num) for left_num, right_num in zip(cols.left, cols.right)
    )


def solve_part2(contents: str) -> int:
    cols = parse_columns(contents)
    return sum(
        left_num * sum(1 for right_num in cols.right if right_num == left_num)
        for left_num in cols.left
    )


if __name__ == "__main__":
    contents = Path("../input.txt").read_text()
    print(f"Part 1: {solve_part1(contents)}")
    print(f"Part 2: {solve_part2(contents)}")


def test_solutions():
    contents = """3   4
4   3
2   5
1   3
3   9
3   3
"""
    assert solve_part1(contents) == 11
    assert solve_part2(contents) == 31
