package main

import (
	"flag"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	for a > b && b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
    flag.Usage = func() {
        fmt.Println("Usage: go run modtabs -n=X")
        flag.PrintDefaults()
    }
	var n int
	flag.IntVar(&n, "n", 0, "mod n")
	flag.Parse()

	nLen := len(strconv.Itoa(n))
	colFmt := fmt.Sprintf("%%%dd", nLen)
	cols := make([]string, 0, n)

	for i := 0; i < n; i++ {
		cols = append(cols, colFmt)
	}

	colFmt = strings.Join(cols, "  ")
	nums := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}

	fmt.Println(fmt.Sprintf(strings.Repeat(" ", nLen+4)+colFmt+"  (order/gcd)", nums...))
	rowFmt := fmt.Sprintf("%%%dd |  %s      %%s", nLen, colFmt)
	for r := 1; r < n; r++ {
		rowElem := make([]interface{}, 0, n+2)
		rowElem = append(rowElem, r)
		order := -1

		for c := 0; c < n; c++ {
			x := big.NewInt(int64(r)).Exp(big.NewInt(int64(r)), big.NewInt(int64(c)), big.NewInt(int64(n))).Int64()
			if order < 0 && c > 0 && x == 1 {
				order = c
			}
			rowElem = append(rowElem, x)
		}

		if order < 0 {
			rowElem = append(rowElem, fmt.Sprintf("%d (gcd)", gcd(n, r)))
		} else if order == n-1 {
			rowElem = append(rowElem, fmt.Sprintf("%d (p)", order))
        } else {
			rowElem = append(rowElem, fmt.Sprintf("%d", order))
		}

		fmt.Println(fmt.Sprintf(rowFmt, rowElem...))
	}
}
