package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run gcd.go X Y")
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Expecting two integers")
		os.Exit(1)
	}

	y, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Expecting two integers")
		os.Exit(1)
	}

	fmt.Printf("gcd(%d, %d) = %d\n", x, y, gcd(x, y))
}
