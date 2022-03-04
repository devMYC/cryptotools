package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type void struct{}

var empty void

func main() {
    flag.Usage = func () {
        fmt.Fprintln(os.Stdout, "Usage: go run polytype.go 0,1,2,...,m")
        fmt.Fprintln(os.Stdout, "Examples: x^4 + x + 1 => 0,1,4")
        fmt.Fprintln(os.Stdout, "          x^4 + x^3 + x^2 + x + 1 => 0,1,2,3,4")
        fmt.Fprintln(os.Stdout, "          x^5 + x^3 + x^2 + 1 => 0,2,3,5")
        flag.PrintDefaults()
    }

    flag.Parse()

	polystr := strings.Split(os.Args[1], ",")
	poly := make([]int, len(polystr))

	for i, numstr := range polystr {
		n, err := strconv.Atoi(numstr)
		if err != nil {
			log.Fatalf("Unable to parse number %s", numstr)
			os.Exit(1)
		}
		poly[i] = n
	}

	m := poly[len(poly)-1]
	maxLen := int(math.Pow(2, float64(m)) - 1)
	p := poly[:len(poly)-1]
    isPrimitive := false
    uniqueCycleLen := map[int]void{}

	for iv := 1; iv <= maxLen; iv++ {
		clk := 0
		states := make(map[int]void)
		state := iv
		states[state] = empty
		for {
			fmt.Printf("%2d  %0*b\n", clk, m, state)
			clk++
			nextBit := 0
			for _, i := range p {
                if (1 << i) & state > 0 {
                    nextBit ^= 1
                } else {
                    nextBit ^= 0
                }
			}
			state >>= 1
			state |= nextBit << (m-1)
			if _, ok := states[state]; ok {
				fmt.Printf("%2d  %0*b\n", clk, m, state)
                if len(states) == maxLen {
                    isPrimitive = true
                }
                uniqueCycleLen[len(states)] = empty
                break
			}
			states[state] = empty
		}
        fmt.Println()
	}

    lens := make([]int, 0, len(uniqueCycleLen))
    for k, _ := range uniqueCycleLen {
        lens = append(lens, k)
    }

    fmt.Println("> max cycle length (if polynomial is primitive):", maxLen)
    fmt.Println("> cycle lengths:", lens)

    if isPrimitive {
        fmt.Println("> primitive")
    } else if len(uniqueCycleLen) == 1 {
        fmt.Println("> irreducible (but not primitive)")
    } else {
        fmt.Println("> reducible")
    }
}
