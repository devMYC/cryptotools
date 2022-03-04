package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func sqMul(x, H, n uint) uint {
	if x == 1 || H == 0 {
		return 1
	}
	t := uint(math.Ceil(math.Log2(float64(H+1))) - 1)
	r := x
	i := int(t - 1)
	for i >= 0 {
		r = (r * r) % n
		if (H & (1 << i)) > 0 {
			r = (r*x) % n
		}
		i--
	}
	return r
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: go run %s -n=N -e=E\n", os.Args[0])
		flag.PrintDefaults()
	}

	nptr := flag.Uint("n", 0, "n = p * q")
	eptr := flag.Uint("e", 0, "the public key and gcd(e, phi(n)) = 1")

	flag.Parse()

	n := *nptr
	e := *eptr

	if n == 0 || e == 0 {
		flag.Usage()
		os.Exit(1)
	}

	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for _, code := range letters {
        fmt.Printf("%s = %d = %d\n", string(code), code, sqMul(uint(code), e, n))
	}
}
