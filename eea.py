import sys


def eea(r0, r1):
    """
    Extended Euclidean Algorithm (EEA)
    Input: positive integer r0 and r1 with r0 > r1
    Output: gcd(r0, r1), as well as s and t such that gcd(r0, r1) = s*r0 + t*r1
    """
    assert r0 > 0 and r1 > 0, "{} > 0 and {} > 0".format(r0, r1)
    assert r0 > r1, "{} > {}".format(r0, r1)
    s0, s1 = 1, 0
    t0, t1 = 0, 1
    while True:
        r = r0 % r1
        q = (r0 - r) // r1
        s = s0 - q * s1
        t = t0 - q * t1
        r0, r1 = r1, r
        s0, s1 = s1, s
        t0, t1 = t1, t
        if r == 0:
            return r0, s0, t0


def multiplicative_inverse(a, b):
    """
    multiplicative inverse of a mod b
    """
    assert a < b, "{} < {}".format(a, b)
    gcd, _, inv = eea(b, a)
    assert gcd == 1, "gcd({}, {}) = {} != 1".format(b, a, gcd)
    while inv < 0:
        inv += b
    return inv


def main():
    if len(sys.argv) != 3 or not all(str.isdigit(n) for n in sys.argv[1:]):
        print("Usage: python3 filename.py a b")
        print("\twhere a, b are positive integers and a < b")
        exit(1)
    [a, b] = map(int, sys.argv[1:])
    print(multiplicative_inverse(a, b))


if __name__ == "__main__":
    main()

