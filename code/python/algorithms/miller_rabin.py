from sys import maxsize
from random import randint


def rand():
    return randint(0, maxsize)


def f(x, y, n):
    if y == 0:
        return 1
    if y == 1:
        return x
    w = f(x, int(y/2), n)
    w = (w * w) % n
    if y % 2 == 1:
        w = (w * x) % n
    return w


def prime(a) -> bool:
    if a == 2:
        return True
    for i in range(1, 5+1):
        b = rand() % (a - 2) + 2
        if f(b, a-1, a) != 1:
            return False
    return True


if __name__ == "__main__":
    print(prime(13))
