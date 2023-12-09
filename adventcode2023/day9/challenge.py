import sys

LINES = list(
    map(
        lambda x: list(map(lambda y: int(y), x.split(" "))),
        open(sys.argv[1]).readlines(),
    )
)
# LINES = [[0, 3, 6, 9, 12, 15], [1, 3, 6, 10, 15, 21], [10, 13, 16, 21, 30, 45]]


def find_next_number(i: list[int], first: bool) -> int:
    if all(x == 0 for x in i):
        return 0

    new_i = []

    for idx in range(1, len(i)):
        new_i.append(i[idx] - i[idx - 1])

    return i[len(i) - 1 if first else 0] + (1 if first else -1) * find_next_number(
        new_i, first
    )


if __name__ == "__main__":
    for first in [True, False]:
        res = sum(map(lambda x: find_next_number(x, first), LINES))
        print(f"Result for {'first' if first else 'second'} part: {res}")
