import sys
import math

def sq_mul(x, H, n):
    assert H >= 0, "exponent can't be negative"
    if x == 1 or H == 0:
        return 1
    t = math.ceil(math.log2(H + 1)) - 1
    r = x
    i = t - 1
    while i >= 0:
        r = r**2 % n
        if (H & (1 << i)) > 0:
            r = r*x % n
        i -= 1
    return r

def main():
    args = sys.argv[1:]
    if len(args) != 3:
        print("Usage: python3 {} x H n" \
              "\n\twhere x - base element" \
              "\n\t      H - exponent" \
              "\n\t      n - modulus"
              .format(sys.argv[0]))
        exit(1)
    try:
        [x, H, n] = [int(n) for n in args]
        print(sq_mul(x%n if x >= n else x, H, n))
    except ValueError:
        print("Invalid input")
        exit(1)

if __name__ == "__main__":
    main()

