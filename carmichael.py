import math
import random
import sys

DEFAULT_SECURITY_PARAM = 20

def sq_mul(x, H, n):
    assert H >= 0
    if x == 1 or H == 0: return 1
    r, t = x, math.ceil(math.log2(H + 1))-1
    for i in range(t-1, -1, -1):
        r = r**2 % n
        if (H & (1 << i)) > 0:
            r = r*x % n
    return r

def is_prime(n):
    sqrt = int(math.sqrt(n))
    for i in range(2, sqrt+1):
        if n % i == 0: return False
    return True

def gcd(a, C):
    if a == 0 or C == 0:
        return C if a == 0 else a
    while C != 0:
        a, C = C, a%C
    return a

def is_prime_fermat_modified(n, s):
    flag = False
    a_set = range(2, n-1)
    for _ in range(0, s):
        a = a_set[random.randint(0, len(a_set)-1)]
        if gcd(a, n) == 1:
            flag = True
            if sq_mul(a, n-1, n) != 1:
                return False
    return flag

def main():
    usage = "Usage:\n\tpython3 {} <limit> [<s>]" \
            "\n\nThe output will be the '3rd from last' " \
            "Carmichael number less than the specified <limit>." \
            "\nThe security parameter <s> is optional. " \
            "(default to {})".format(sys.argv[0], DEFAULT_SECURITY_PARAM)
    args = sys.argv[1:]
    if len(args) != 1 and len(args) != 2:
        print(usage)
        exit(1)
    limit, *s = args
    try:
        limit = int(limit)
        s = DEFAULT_SECURITY_PARAM if len(s) == 0 else int(s[0])
    except ValueError:
        print("Invalid input")
        exit(1)
    assert limit >= 2
    found, result = 0, None
    for i in range(limit-1, 1, -1):
    # for i in range(2, limit):
        if not is_prime(i) and is_prime_fermat_modified(i, s):
            found, result = found+1, i
            if found == 3: break
            # if found == ?: break
    print(result if found == 3 else None)
    # print(result if found == ? else None)

if __name__ == "__main__":
    main()
