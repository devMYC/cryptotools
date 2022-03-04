package main

import (
	"flag"
	"fmt"
	"os"
)

const M = 26

func main() {
    flag.Usage = func () {
        fmt.Fprintln(os.Stdout, "Usage: go run shiftCipher.go [params...]")
        flag.PrintDefaults()
    }

    kPtr := flag.Int("k", 0, "encryption/decryption key")
    cipherPtr := flag.String("cipher", "", "Ciphertext")
    plainPtr := flag.String("plain", "", "Plaintext")

    flag.Parse()

    if (len(*cipherPtr) > 0) {
        cipher := *cipherPtr
        plain := make([]rune, len(cipher))

        for k := 0; k < M; k++ {
            for i, r := range []rune(cipher) {
                if r == '\n' {
                    plain[i] = r
                    continue
                }

                j := int(r) - 'a' - k
                if j < 0 {
                    plain[i] = rune(j + M + 'a')
                } else {
                    plain[i] = rune(j + 'a')
                }
            }

            fmt.Printf("key = %d\nPlaintext = %s\n\n", k, string(plain))
        }
    } else {
        k := *kPtr
        plain := *plainPtr
        cipher := make([]rune, 0, len(plain))
        for _, r := range []rune(plain) {
            cipher = append(cipher, rune((int(r)-'a'+k) % 26 + 'a'))
        }

        fmt.Println("Ciphertext =", string(cipher))
    }

}

