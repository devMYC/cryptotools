import sys

# P(x) = x^8 + x^4 + x^3 + x + 1
P_x = 283


def pos_offset(n):
    offset = 0
    while n >= 256:
        offset += 1
        n >>= 1
    return offset-1 if offset > 0 else offset


def main():
    args = sys.argv[1:]
    assert len(args) == 2, ("Usage: python3 filename.py a b\n" \
                            "\t\twhere a and b are polynomials " \
                            "in hexadecimal notation.")
    [a, b] = args
    try:
        a = int(a, 16)
        b = int(b, 16)
    except ValueError:
        print("Input contains invalid hex string.")
        exit(1)
    if a == 0 or b == 0:
        return print(hex(0))
    result = 0 if b & 1 == 0 else a
    while b != 1:
        a <<= 1
        b >>= 1
        if b & 1 == 0:
            continue
        result ^= a
    while result >= 256:
        result ^= P_x << pos_offset(result)
    print(hex(result))


if __name__ == "__main__":
    main()

