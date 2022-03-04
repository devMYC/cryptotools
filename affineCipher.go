package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const M = 26

func main() {
    flag.Usage = func () {
        fmt.Fprintln(os.Stdout, "Usage: go run affineCipher.go -a=X -b=Y [-e=true, -d=true]")
        flag.PrintDefaults()
    }
	a := flag.Int("a", 1, "an integer `a` within [0..25] which has the restriction gcd(a, 26) = 1")
	b := flag.Int("b", 0, "an integer in the range of [0..25]")
	encrypt := flag.Bool("e", false, "encrypt")
	decrypt := flag.Bool("d", false, "decrypt")

	flag.Parse()

	aInverse := 0
	for i := 1; i < M; i++ {
		if (*a*i)%M == 1 {
			aInverse = i
			break
		}
	}

	if aInverse == 0 {
		log.Fatal("`a` does not have an inverse")
		os.Exit(1)
	}

	input := flag.Arg(0)
	result := make([]rune, len(input))

	for i, c := range []rune(input) {
		if *encrypt {
			result[i] = rune(*a*int(c-'a') + *b%26 + 'a')

		} else if *decrypt {
			result[i] = rune(aInverse*(int(c-'a')-*b+M)%26 + 'a')
		}
	}

	fmt.Println(strings.ToUpper(string(result)))
}

