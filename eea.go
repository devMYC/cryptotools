package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: go run eea.go -r0=X -r1=Y  (X >= Y)\n")
        flag.PrintDefaults()
	}

	r0ptr := flag.Int("r0", 0, "1st parameter of gcd")
	r1ptr := flag.Int("r1", 0, "2nd parameter of gcd")
	flag.Parse()

	r0, r1 := *r0ptr, *r1ptr
	s0, s1 := 1, 0
	t0, t1 := 0, 1

	for {
		r := r0 % r1
		q := (r0 - r) / r1
		s := s0 - q*s1
		t := t0 - q*t1

        r0, r1 = r1, r
		s0, s1 = s1, s
		t0, t1 = t1, t

		if r == 0 {
			break
		}
	}

    fmt.Printf("gcd(%d, %d) = %d\ns = %d\nt = %d\n", *r0ptr, *r1ptr, r0, s0, t0)
}

