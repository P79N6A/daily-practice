def mutual_prime(a, b) -> bool:
    if a < b:
        a, b = b, a
    while b != 0:
        remainder = a % b 
        a, b = b, remainder
    if a == 1:
        return True
    return False


if __name__ == "__main__":
    print(mutual_prime(9851, 7963))